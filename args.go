package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type outputOpts struct {
	helpFlag bool
}

type options struct {
	output  *outputOpts
}

func validate(opts *options) (*options, error) {
	if opts.output.helpFlag {
		return opts, nil
	}
	return opts, nil
}

func parseArgs(args []string) (*options, error) {
	opts := &options{output: &outputOpts{}}
	flags := flag.NewFlagSet("teacat", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.output.helpFlag, "help", "h", false, "prints this message")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	return validate(opts)
}