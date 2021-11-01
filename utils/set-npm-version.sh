#!/bin/sh
version=$1
cat package.json | jq ".version = \"$version\"" > package.json

