#! /bin/bash

set -e

if [ $# -eq 0 ]; then
    echo "Need day as first argument!"
    exit
fi

if [[ ! $1 =~ ^[0-9]+$ ]] || [ ! $1 -ge 1 ] || [ ! "$1" -le 25 ]; then
    echo "Number must an integer between 1 and 25!"
    exit
fi

mkdir -p "day$1"

source .env
curl -sH "Cookie: session=$SESSION_TOKEN" https://adventofcode.com/2023/day/$1/input > ./day$1/input.txt

echo "Input is now in day$1/input.txt ! :D"