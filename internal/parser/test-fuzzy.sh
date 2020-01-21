#!/bin/bash

cd $(dirname "$0")

echo "Install go-fuzz"
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

echo "Build fuzzy executable"
go-fuzz-build -o testdata/parser-fuzz.zip .
echo "Start fuzzing"
go-fuzz -bin testdata/parser-fuzz.zip -workdir testdata .