package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type wcOpts struct {
	pLinesFlag bool
	pBytesFlag bool
	pWordsFlag bool
	pCharsFlag bool
}

type catOpts struct {
	printFlag     bool
	linePrintFlag bool
}

type inoutputOpts struct {
	helpFlag bool
	args     []string
}

type options struct {
	wc       *wcOpts
	cat      *catOpts
	inoutput *inoutputOpts
}

func parseArgs(args []string) (*options, error) {
	opts := &options{wc: &wcOpts{}, cat: &catOpts{}, inoutput: &inoutputOpts{}}

	flags := flag.NewFlagSet("teacat", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.wc.pLinesFlag, "lines", "l", false, "prints of line count")
	flags.BoolVarP(&opts.wc.pBytesFlag, "bytes", "b", false, "prints of byte count")
	flags.BoolVarP(&opts.wc.pWordsFlag, "words", "w", false, "prints of word count")
	flags.BoolVarP(&opts.wc.pCharsFlag, "chars", "c", false, "prints of char count")
	flags.BoolVarP(&opts.cat.printFlag, "print", "p", false, "prints all contents")
	flags.BoolVarP(&opts.cat.linePrintFlag, "lineprint", "n", false, "prints all contents and line numbers")
	flags.BoolVarP(&opts.inoutput.helpFlag, "help", "h", false, "prints this message")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	opts.inoutput.args = flags.Args()[1:]
	fmt.Println(opts.inoutput.args)

	return opts, nil
}
