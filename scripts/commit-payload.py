"""Serialize the staged tutorial changes for GitHub's signed commit API."""

import base64
import json
import subprocess
import sys


def git(*args):
    return subprocess.check_output(["git", *args])


def payload(branch, subject):
    additions = []
    deletions = []
    # Disable rename detection: a rename is a deletion plus an addition.
    changes = git("diff", "--cached", "--name-status", "--no-renames", "-z").split(b"\0")
    for status, path in zip(changes[0::2], changes[1::2]):
        path = path.decode("utf-8")
        if status == b"D":
            deletions.append({"path": path})
        else:
            additions.append({
                "path": path,
                "contents": base64.b64encode(git("show", ":" + path)).decode("ascii"),
            })
    return {
        "query": """mutation($input: CreateCommitOnBranchInput!) {
          createCommitOnBranch(input: $input) { commit { url } }
        }""",
        "variables": {"input": {
            "branch": {"repositoryNameWithOwner": "grafana/killercoda", "branchName": branch},
            "expectedHeadOid": git("rev-parse", "HEAD").decode().strip(),
            "message": {"headline": subject},
            "fileChanges": {"additions": additions, "deletions": deletions},
        }},
    }


if __name__ == "__main__":
    json.dump(payload(*sys.argv[1:]), sys.stdout)
