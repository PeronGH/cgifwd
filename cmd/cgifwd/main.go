package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/PeronGH/cgifwd/internal/errorprinter"
	"github.com/PeronGH/cgifwd/internal/forwarder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cgifwd <target base url>")
		os.Exit(1)
	}

	target := os.Args[1]
	targetURL, err := url.Parse(target)
	if err != nil {
		errorprinter.PrintErrorInCGI(err)
		os.Exit(1)
	}

	err = forwarder.StartForwarderFor(targetURL)
	if err != nil {
		errorprinter.PrintErrorInCGI(err)
		os.Exit(1)
	}
}
