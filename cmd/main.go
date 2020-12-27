package main

import (
	"fmt"
	"os"

	surt "simplesurt-proto"
)

const (
	SUCCESS_EXITCODE = iota
	FATAL_EXITCODE
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	formatted, err := surt.Format(os.Args[1])
	if err != nil {
		fatal("failed to format: %s", err)
	}
	fmt.Println(formatted)
}

// usage prints the applications usage/help and exits
func usage() {
	fmt.Printf(
		"usage: %s <uri>"+
			" -- format URI in SURT format\n",
		os.Args[0],
	)
	os.Exit(SUCCESS_EXITCODE)
}

// fatal log the formatted message as an error and exits on FATAL_EXITCODE
func fatal(format string, a ...interface{}) {
	fmt.Println(fmt.Errorf(format, a...))
	os.Exit(FATAL_EXITCODE)
}
