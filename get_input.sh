#! /bin/bash

set -e

source .env

if [ $# -eq 0 ]
then
    echo "Need day to be specified"
fi

mkdir -p "day$1"
curl -sH "Cookie: session=$SESSION_TOKEN" https://adventofcode.com/2023/day/$1/input > ./day$1/input.txt
echo "Input is now in day$1/input.txt ! :D"