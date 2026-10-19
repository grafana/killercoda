"""Exercise PR management without contacting GitHub."""
import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

SCRIPTS = Path(__file__).resolve().parent


class ManagePRTest(unittest.TestCase):
    def run_scenario(self, *, changed=True, existing=False, fail=False):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            repo = root / "repo"
            repo.mkdir()
            def git(*args):
                return subprocess.check_output(["git", *args], cwd=repo, stderr=subprocess.DEVNULL)
            git("init")
            git("config", "user.name", "Test")
            git("config", "user.email", "test@example.com")
            git("remote", "add", "origin", str(root / "remote.git"))
            subprocess.run(["git", "init", "--bare", str(root / "remote.git")], check=True, capture_output=True)
            shutil.copytree(SCRIPTS, repo / "scripts", ignore=shutil.ignore_patterns("__pycache__"))
            (repo / "tutorial.md").write_text("original")
            git("add", ".")
            git("-c", "commit.gpgsign=false", "commit", "-m", "initial")
            git("checkout", "-b", "update-generated-tutorials")
            if existing:
                git("push", "origin", "HEAD")
            if changed:
                (repo / "new tutorial.md").write_text("new tutorial")
            bin_dir = root / "bin"
            bin_dir.mkdir()
            gh = bin_dir / "gh"
            gh.write_text('''#!/usr/bin/env python3
import json, os, pathlib, sys
args = sys.argv[1:]
with open(os.environ["CALLS"], "a") as log:
    log.write(json.dumps(args) + "\\n")
if args[:2] == ["api", "graphql"]:
    pathlib.Path(os.environ["PAYLOAD"]).write_text(sys.stdin.read())
    sys.exit(int(os.environ["FAIL"]))
if args[:2] == ["pr", "list"] and os.environ["EXISTING"] == "1":
    print("42")
''')
            gh.chmod(0o755)
            env = dict(os.environ, PATH=str(bin_dir) + os.pathsep + os.environ["PATH"], CALLS=str(root / "calls"), PAYLOAD=str(root / "payload"), FAIL=str(int(fail)), EXISTING=str(int(existing)))
            result = subprocess.run(["bash", "scripts/manage-pr.bash"], cwd=repo, env=env, capture_output=True, text=True)
            self.assertEqual(result.returncode, int(fail), result.stderr)
            calls = [json.loads(line) for line in (root / "calls").read_text().splitlines()] if (root / "calls").exists() else []
            if not changed:
                self.assertEqual(calls, [])
                return
            data = json.loads((root / "payload").read_text())["variables"]["input"]
            self.assertEqual(data["fileChanges"]["additions"][0]["path"], "new tutorial.md")
            self.assertEqual(any(c[:2] == ["pr", "create"] for c in calls), not existing and not fail)
            self.assertEqual(any(c[:2] == ["api", "repos/grafana/killercoda/git/refs"] for c in calls), not existing)

    def test_no_changes(self):
        self.run_scenario(changed=False)

    def test_new_branch_and_new_files(self):
        self.run_scenario()

    def test_update_existing_pr(self):
        self.run_scenario(existing=True)

    def test_commit_failure_stops_pr_creation(self):
        self.run_scenario(fail=True)


if __name__ == "__main__":
    unittest.main()
