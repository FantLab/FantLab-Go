#!/bin/bash

rm -rf app/pb
mkdir -p app/pb
protoc --go_out=. proto/*.proto
