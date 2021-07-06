package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

// FlagOptions is file options
type FlagOptions struct {
	printLines bool
	printBytes bool
	printWords bool
	printCharacters bool
	printHelp  bool
	printCat   bool
	printNum   bool
}

// Counter is counter of elements
type Counter struct {
	lines int
	words int
	bytes int
	characters int
}

func parseFlagOptions() FlagOptions {
	var opts FlagOptions
	flag.BoolVar(&opts.printLines, "l", false, "print lines")
	flag.BoolVar(&opts.printBytes, "b", false, "print bytes")
	flag.BoolVar(&opts.printCharacters, "c", false, "print characters")
	flag.BoolVar(&opts.printWords, "w", false, "print words")
	flag.BoolVar(&opts.printHelp, "h", false, "print help")
	flag.BoolVar(&opts.printCat, "p", false, "print all words")
	flag.BoolVar(&opts.printNum, "lp", false, "print all words and line number")
	flag.Parse()

	if !opts.printLines && !opts.printBytes && !opts.printWords && !opts.printCharacters{
		opts.printLines = true
		opts.printBytes = true
		opts.printWords = true
		opts.printCharacters = true
	}

	return opts
}

// Count is count of elements
func (c *Counter) Count(r io.Reader) (bool, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return false, err
	}
	c.lines = bytes.Count(b, []byte{'\n'})
	c.bytes = len(b)
	c.words = len(bytes.Fields(b))
	return true, nil
}

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <FILEs...|DIRs...>
	-b, --byte                  Prints the number of bytes in each input file.
	-c, --character             Prints the number of characters in each input file.
	-l, --line                  Prints the number of lines in each input file.
	-w, --word                  Prints the number of words in each input file.
	-p, --print                 Prints the all words in each input file.
	-lp,--lineprint             prints the all words and line number in each input file.
	-h, --help                  Prints this message.
ARGUMENTS
	FILEs...                    Specifies targets. TeaCat accepts zip/tar/tar.gz/ files.
	DIRs...                     Files in the given directory are as the input files.`, programName)
}

// Show is output of elements
func (c Counter) Show(opts FlagOptions, filename string) {
	fmt.Printf("File name is %s\n", filename)
	if opts.printLines {
		fmt.Printf("%7d lines", c.lines)
	}
	if opts.printWords {
		fmt.Printf("%7d words", c.words)
	}
	if opts.printBytes {
		fmt.Printf("%7d bytes", c.bytes)
	}
	if opts.printCharacters {
		fmt.Printf("%7d chars", c.characters)
	}
	fmt.Println()
}

// Add is add of elements
func (c *Counter) Add(src Counter) {
	c.lines += src.lines
	c.bytes += src.bytes
	c.words += src.words
}

func filecat(opts FlagOptions) int{
	filenames := flag.Args()
	i := 1

	for _, filename := range filenames {
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return 1
		} else {
			scanner := bufio.NewScanner(fp)
			for ; scanner.Scan(); i++ {
				if opts.printNum {
					//オプションがある場合
					fmt.Printf("%v: ", i)
				}
				fmt.Println(scanner.Text())
			}
		}
	}
	return 0
}

func goMain(args []string) int {
	opts := parseFlagOptions()
	filenames := flag.Args()

	var totalCount Counter

	if opts.printHelp {
		fmt.Println(helpMessage(filepath.Base(args[1])))
		return 0
	}

	if len(filenames) == 0 {
		var c Counter
		_, err := c.Count(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		c.Show(opts, "")
	}

	if opts.printCat || opts.printNum {
		status := filecat(opts)
		return status
	}

	for _, filename := range filenames {
		var c Counter
		if opts.printCharacters {
			data, _ := ioutil.ReadFile(filename)
			c.characters = utf8.RuneCountInString(string(data))
		}
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = c.Count(fp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		totalCount.Add(c)
		c.Show(opts, filename)
		fp.Close()
	}

	if len(filenames) > 1 {
		totalCount.Show(opts, "total")
	}

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}