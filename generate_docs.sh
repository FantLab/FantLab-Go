#!/bin/bash

cd sources || exit
go run . -gendocs > ../docs/generated.md
cd ..