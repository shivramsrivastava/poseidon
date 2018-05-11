#!/bin/bash
cat license.txt | cat - $1 > temp && mv temp $1 
