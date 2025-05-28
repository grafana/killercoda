#!/usr/bin/env bash

set -euf -o pipefail

function usage {
  cat <<EOF
Check out the branch for regenerating tutorials.

Usage:
  $0

Examples:
  $0
EOF
}

if [[ $# -ne 0 ]]; then
  usage
  exit 1
fi

if [ -n "${RUNNER_DEBUG+x}" ]; then
  set -x
fi

readonly BRANCH="${BRANCH:-update-generated-tutorials}"

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
  git config --local user.email bot@grafana.com
  git config --local user.name grafanabot
  git remote set-url origin "https://x-access-token:${GH_TOKEN}@github.com/grafana/killercoda"

  git push origin --delete "${BRANCH}"
fi

git checkout -b "${BRANCH}"
