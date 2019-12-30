#!/bin/bash

rm -rf sources/server/internal/pb
mkdir -p sources/pb
protoc --go_out=. proto/*.proto
