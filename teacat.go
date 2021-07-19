package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func perform(opts *options, filenames[] string) int {
	if opts.wc.pBytesFlag || opts.wc.pLinesFlag || opts.wc.pCharsFlag || opts.wc.pWordsFlag {
		return wordCount(opts, filenames)
	}

	if opts.cat.printFlag || opts.cat.linePrintFlag {
		return cat(opts, filenames)
	}

	return 1
}

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <FILEs...|DIRs...>
	-b, --byte                  Prints the number of bytes in each input file.
	-c, --character             Prints the number of characters in each input file.
	-l, --line                  Prints the number of lines in each input file.
	-w, --word                  Prints the number of words in each input file.
	-p, --print                 Prints the all words in each input file.
	-n, --lineprint             prints the all words and line number in each input file.
	-h, --help                  Prints this message.
ARGUMENTS
	FILEs...                    Specifies targets. TeaCat accepts zip/tar/tar.gz/ files.
	DIRs...                     Files in the given directory are as the input files.`, programName)
}

func goMain(args[] string) int {
	opts, err := parseArgs(args)

	if err != nil {
		fmt.Printf("%s: %s\n", filepath.Base(args[0]), err.Error())
		fmt.Println(helpMessage(args[0]))
		return 1
	}

	if opts.inoutput.helpFlag || len(opts.inoutput.args) == 0{
		fmt.Println(helpMessage(args[0]))
		return 0
	}

	return perform(opts, opts.inoutput.args)
}

func main() {
	status:= goMain(os.Args)
	os.Exit(status)
}