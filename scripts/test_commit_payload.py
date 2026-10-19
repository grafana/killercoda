import base64
import importlib.util
import os
from pathlib import Path
import subprocess
import tempfile
import unittest

spec = importlib.util.spec_from_file_location("commit_payload", Path(__file__).with_name("commit-payload.py"))
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)


class CommitPayloadTest(unittest.TestCase):
    def test_staged_add_modify_delete_and_rename(self):
        with tempfile.TemporaryDirectory() as directory:
            previous = os.getcwd()
            os.chdir(directory)
            self.addCleanup(os.chdir, previous)
            def git(*args):
                return subprocess.check_output(["git", *args], stderr=subprocess.DEVNULL)
            git("init")
            git("config", "user.name", "Test")
            git("config", "user.email", "test@example.com")
            for name in ("modified", "deleted", "old name"):
                Path(name).write_text("original")
            git("add", ".")
            git("-c", "commit.gpgsign=false", "commit", "-m", "initial")
            head = git("rev-parse", "HEAD").decode().strip()
            Path("modified").write_text("staged")
            Path("deleted").unlink()
            Path("old name").rename("new name")
            Path('new "file"').write_bytes(b"\x00\xff")
            git("add", ".")
            Path("modified").write_text("unstaged")
            data = module.payload("updates", 'Update "tutorials"')["variables"]["input"]
            self.assertEqual(data["expectedHeadOid"], head)
            self.assertEqual(data["branch"]["branchName"], "updates")
            additions = {item["path"]: base64.b64decode(item["contents"]) for item in data["fileChanges"]["additions"]}
            self.assertEqual(additions, {"modified": b"staged", "new name": b"original", 'new "file"': b"\x00\xff"})
            self.assertEqual({item["path"] for item in data["fileChanges"]["deletions"]}, {"deleted", "old name"})
            os.chdir(previous)


if __name__ == "__main__":
    unittest.main()
