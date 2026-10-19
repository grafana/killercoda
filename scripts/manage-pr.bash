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

# GitHub signs commits created through this API with the workflow token.
# Stage first so new files and deletions are included in the no-change check.
git add .
if git diff --cached --quiet; then
  exit 0
fi

readonly HEAD_OID="$(git rev-parse HEAD)"
if [[ -z "$(git ls-remote --heads origin "refs/heads/${BRANCH}")" ]]; then
  gh api repos/grafana/killercoda/git/refs \
    -f "ref=refs/heads/${BRANCH}" -f "sha=${HEAD_OID}" >/dev/null
fi

python3 "$(dirname "${BASH_SOURCE[0]}")/commit-payload.py" "${BRANCH}" "${SUBJECT}" |
  gh api graphql --input -

# Updating an existing PR only needs the commit; do not try to create it again.
if [[ -z "$(gh pr list --head "${BRANCH}" --base staging --state open --json number --jq '.[].number')" ]]; then
  gh pr create --title "${SUBJECT}" --body '' --base staging --head "${BRANCH}"
fi
