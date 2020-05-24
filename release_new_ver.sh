#!/bin/sh

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 [new version]" >&2
  exit 1
fi

git-chglog --next-tag $1 -o CHANGELOG.md

git add -u 
git commit -am "release $1"

git tag $1