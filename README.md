[![build](https://github.com/akhiroky/TeaCat/actions/workflows/build.yml/badge.svg)](https://github.com/akhiroky/TeaCat/actions/workflows/build.yml)
[![License](http://img.shields.io/badge/license-CC0-green.svg)](https://github.com/akhiroky/TeaCat/blob/main/LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/akhiroky/TeaCat/badge.svg?branch=main)](https://coveralls.io/github/akhiroky/TeaCat?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/akhiroky/TeaCat)](https://goreportcard.com/report/github.com/akhiroky/TeaCat)
[![codebeat badge](https://codebeat.co/badges/501f4324-3296-4c48-ba93-08b2a8459067)](https://codebeat.co/projects/github-com-akhiroky-teacat-main)

[![Docker](https://img.shields.io/badge/Docker-hirokiiii%2Fteacat%3A1.0.0-green?logo=docker)](https://hub.docker.com/repository/docker/hirokiiii/teacat)
[![DOI](https://zenodo.org/badge/369699051.svg)](https://zenodo.org/badge/latestdoi/369699051)
# TeaCat
Implementation of additional functions to wc

![TeaCat](cartoon-little-cat-with-tea-cup.jpg)

## Description
This software adds functionality to the wc command.
The wc command displays the number of lines, words, and bytes in a specified text file.
However, TeaCat will also display the contents of the file. 
It is a combination of the cat command and the wc command.

## Usage
```
TeaCat [Options]... [FILEs...|DIRs...]
    -b, --byte                  Prints the number of bytes in each input file.
    -c, --character             Prints the number of characters in each input file.
    -l, --line                  Prints the number of lines in each input file.
    -w, --word                  Prints the number of words in each input file.
    -p, --print                 Prints the all words in each input file.
    -lp,--lineprint             prints the all words and line number in each input file.
    -h, --help                  Prints this message. 
ARGUMENTS
    FILEs...                    Specifies targets. TeaCat accepts zip/tar/tar.gz/ files.
    DIRs...                     Files in the given directory are as the input files.

