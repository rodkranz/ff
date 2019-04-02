#!/usr/bin/env bash

rm ./tmp/ff

make build-dev

./tmp/ff

ls -la ./tmp/

go tool trace ./tmp/t.out
