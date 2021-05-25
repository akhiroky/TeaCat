# TeaCat
Implementation of additional functions to wc

## Description
This software adds functionality to the wc command.
The wc command displays the number of lines, words, and bytes in a specified text file.
However, TeaCat will also display the contents of the file.

## Usage
```
TeaCat [FILEs...|DIRs...]
    -b, --byte                  Prints the number of bytes in each input file.
    -c, --character             Prints the number of characters in each input file.
    -l, --line                  Prints the number of lines in each input file.
    -w, --word                  Prints the number of words in each input file.
    -h, --help                  Prints this message. 
ARGUMENTS
    FILEs...                    Specifies counting targets. wildcat accepts zip/tar/tar.gz/ files.
    DIRs...                     Files in the given directory are as the input files.
