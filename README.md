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
TeaCat [FILEs...|DIRs...]
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
