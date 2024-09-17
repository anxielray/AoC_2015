#!/bin/bash

mkdir "day_$1"
cd "day_$1"
mkdir puzzle1 puzzle2
touch puzzle1/main.go puzzle2/main.go
cd puzzle1
echo 'package main' >> main.go
cd ../puzzle2
echo 'package main' >> main.go
