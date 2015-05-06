package application

import (
	"io"
	"log"
	"net/http"
)

type App struct {
	TargetBaseURL string
	LogWriter     io.Writer
}

func (a *App) Run() int {
	logger := log.New(a.LogWriter, "", log.LstdFlags)
	client := &http.Client{}

	resp, err := client.Get(a.TargetBaseURL)
	if err != nil {
		logger.Printf("%s\n", err)
		return 1
	}

	logger.Printf("Response Code: %s\n", resp.Status)

	return 0
}
