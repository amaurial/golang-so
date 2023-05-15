#!/bin/bash

go build -o lib3.so -buildmode=plugin lib3.go
cp lib3.so ../libs/
