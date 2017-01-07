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


GOFMT="gofmt -s -w";
GOLINT="golint";
GOVET="go vet";
find_files | xargs $GOFMT | true;
find_files | xargs $GOLINT | true;
find_files | xargs $GOVET | true;