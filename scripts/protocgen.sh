#!/usr/bin/env bash

set -eo pipefail

cd proto
proto_dirs=$(find ./desmos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
    buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..
# move proto files to the right places
cp -r github.com/desmos-labs/desmos/v5/* ./
rm -rf github.com

./scripts/protocgen-pulsar.sh