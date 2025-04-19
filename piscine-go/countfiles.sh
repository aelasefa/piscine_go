#! /bin/bash

COUNT=$(find . -type f -o -type d | wc -l)
echo $COUNT
