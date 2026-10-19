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

if [ -n "${RUNNER_DEBUG+x}" ]; then
  set -x
fi

readonly BRANCH="${BRANCH:-update-generated-tutorials}"
readonly SUBJECT="${SUBJECT:-Update generated tutorials}"

# Updating an existing PR only needs the commit; do not try to create it again.
if [[ -z "$(gh pr list --head "${BRANCH}" --base staging --state open --json number --jq '.[].number')" ]]; then
  gh pr create --title "${SUBJECT}" --body '' --base staging --head "${BRANCH}"
fi
