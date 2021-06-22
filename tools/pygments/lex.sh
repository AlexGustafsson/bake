#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

python3 -m pygments -x -l "$DIR/lexer.py:BakeLexer" "$@"
