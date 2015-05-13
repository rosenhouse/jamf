package application

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type App struct {
	TargetBaseURL string
	LogWriter     io.Writer
	OutputDir     string
}

func (a *App) Run() int {
	logger := log.New(a.LogWriter, "", log.LstdFlags)
	client := &http.Client{}

	request, err := http.NewRequest("GET", a.TargetBaseURL, nil)
	if err != nil {
		panic(err)
	}
	requestBytes := &bytes.Buffer{}
	err = request.Write(requestBytes)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get(a.TargetBaseURL)
	if err != nil {
		logger.Printf("%s\n", err)
		return 1
	}

	logger.Printf("Response Code: %s\n", resp.Status)

	return 0
}
