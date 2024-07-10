#!/usr/bin/env bash

set -euf -o pipefail

function usage {
  cat <<EOF
Open a PR for updating generated tutorials if there are changes.

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

readonly BRANCH="${BRANCH:-update-generated-tutorials}"
readonly SUBJECT="${SUBJECT:-Update generated tutorials}"

# commit adds and commits the updated files as the Grafanabot GitHub user.
function commit {
  git add .
  git config --local user.email bot@grafana.com
  git config --local user.name grafanabot
  git commit --message "${SUBJECT}"
}

if ! git diff --exit-code; then
  commit
  git push origin "refs/heads/${BRANCH}"
  gh pr create --title "${SUBJECT}" --body '' --base staging --head "${BRANCH}"
fi
