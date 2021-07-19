package main

import(
	"fmt"
	"io"
	"bytes"
	"os"
	"io/ioutil"
	"unicode/utf8"
)

type counter struct {
	lines int
	words int
	bytes int
	chars int
}

func (c *counter) countShow(r io.Reader, opts *options) (bool, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return false, err
	}

	if opts.wc.pBytesFlag {
		c.bytes = len(b)
		fmt.Printf("%7d bytes", c.bytes)
	}

	if opts.wc.pLinesFlag {
		c.lines = bytes.Count(b, []byte{'\n'})
		fmt.Printf("%7d lines", c.lines)
	}

	if opts.wc.pWordsFlag {
		c.words = len(bytes.Fields(b))
		fmt.Printf("%7d words", c.words)
	}

	return true, nil
}

func wordCount(opts *options, filenames[] string) int {

	for _, filename := range filenames {
		var c counter

		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = c.countShow(fp, opts)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if opts.wc.pCharsFlag {
			data, _ := ioutil.ReadFile(filename)
			c.chars = utf8.RuneCountInString(string(data))
			fmt.Printf("%7d chars", c.chars)
		}

		fp.Close()
	}

	return 0
}