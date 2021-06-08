package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func helpMessage(originalProgramName string) string {
    programName := filepath.Base(originalProgramName)
    return fmt.Sprintf(`%s [FILEs...|DIRs...] [OPTIONS]
    -b, --byte                  Prints the number of bytes in each input file.
    -c, --character             Prints the number of characters in each input file.
    -l, --line                  Prints the number of lines in each input file.
    -w, --word                  Prints the number of words in each input file.
    -p, --print                 Prints the all words in each input file.
    -lp,--lineprint             prints the all words and line number in each input file.
    -h, --help                  Prints this message. 
ARGUMENTS
    FILEs...                    Specifies targets. TeaCat accepts zip/tar/tar.gz/ files.
    DIRs...                     Files in the given directory are as the input files.`, programName, programName)
}

func goMain(args []string) int {
    fmt.Println(helpMessage(filepath.Base(args[0])))
    return 0
}

func main() {
    status := goMain(os.Args)
    os.Exit(status)
}