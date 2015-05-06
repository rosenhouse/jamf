package main

import (
	"flag"
	"os"

	"github.com/rosenhouse/jamf/application"
)

var (
	TargetBaseURL string
)

func main() {

	flag.StringVar(&TargetBaseURL, "t", "http://localhost:8888", "Target server's base URL")

	flag.Parse()

	app := application.App{
		TargetBaseURL: TargetBaseURL,
		LogWriter:     os.Stderr,
	}

	app.Run()
}
