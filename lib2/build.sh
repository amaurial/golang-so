#!/bin/bash

go build -o lib2.so -buildmode=plugin lib2.go
cp lib2.so ../libs/
