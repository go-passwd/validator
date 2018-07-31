#!/bin/bash
set -ev
cd passwd
if [ "${COVERAGE}" = "true" ]; then
  go test -v -covermode=count -coverprofile=coverage.out ./...
  goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
else
  go test -v ./...
fi
