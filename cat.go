package main

import (
	"bufio"
	"os"
	"fmt"
)

func cat(opts *options, filenames[] string) int{
	i := 1

	for _, filename := range filenames {
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return 1
		} else {
			scanner := bufio.NewScanner(fp)
			for ; scanner.Scan(); i++ {
				if opts.cat.linePrintFlag {
					//オプションがある場合
					fmt.Printf("%v: ", i)
				}
				fmt.Println(scanner.Text())
			}
		}
	}
	return 0
}