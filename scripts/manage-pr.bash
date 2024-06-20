#!/usr/bin/env bash

set -euf -o pipefail

readonly BRANCH="${BRANCH:-update-generated-tutorials}"
readonly SUBJECT="${SUBJECT:-Update generated tutorials}"

# commit adds and commits the updated files as the Grafanabot GitHub user.
function commit {
  git add .
  git config --local user.email bot@grafana.com
  git config --local user.name grafanabot
  git commit --message "${SUBJECT}"
}

PR_STATUS=$(gh pr view "${BRANCH}" --json state --jq .state || true)
readonly PR_STATUS

if [[ "${PR_STATUS}" == OPEN ]]; then
  gh pr checkout "${BRANCH}"

  if ! git diff --exit-code; then
    commit
    git push
  fi

  exit 0
fi

# Remove the remote branch if it exists because there is no PR associated with it.
if git fetch origin "${BRANCH}"; then
  git push origin --delete "${BRANCH}"
fi

git checkout -b "${BRANCH}"

if ! git diff --exit-code; then
  commit
  git push origin "refs/heads/${BRANCH}"
  gh pr create --title "${SUBJECT}" --body '' --base main --head "${BRANCH}"
fi
