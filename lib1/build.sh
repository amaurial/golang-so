#!/bin/bash

go build -o lib1.so -buildmode=plugin lib1.go
cp lib1.so ../libs/