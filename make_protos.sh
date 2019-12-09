#!/bin/bash

rm -rf sources/server/internal/pb
mkdir -p sources/server/internal/pb
protoc --go_out=. proto/*.proto
