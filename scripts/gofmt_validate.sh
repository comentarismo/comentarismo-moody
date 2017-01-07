#!/usr/bin/env bash

find_files() {
  find . -not \( \
      \( \
        -wholename './static' \
        -o -wholename './git-hooks' \
        -o -wholename './scripts' \
        -o -wholename './templates' \
        -o -wholename './testdata' \
        -o -wholename '*/vendor/*' \
      \) -prune \
    \) -name '*.go'
}

GOFMT="gofmt -s"
bad_files=$(find_files | xargs $GOFMT -l)
if [[ -n "${bad_files}" ]]; then
  echo "!!! '$GOFMT' needs to be run on the following files: "
  echo "${bad_files}"
  exit 1
fi