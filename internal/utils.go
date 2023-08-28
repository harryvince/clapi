package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	VERSION = "0.0.1"
)

func GetVersion() string {
	return VERSION
}

type Response struct {
	Body       string
	StatusCode int
	Headers    http.Header
}

func GetRequestData(request *http.Response) (Response, error) {
	Log("Getting Request Data...")
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return Response{}, err
	}
	defer request.Body.Close()

	response := Response{Body: strings.TrimSpace(string(body)), StatusCode: request.StatusCode, Headers: request.Header}
	Log(fmt.Sprintf("Response parsed => %v", response))

	return response, nil
}

func Log(message string) {
	logLevel := os.Getenv("LOG")

	// Check if the log level is set to DEBUG
	if strings.ToUpper(logLevel) == "DEBUG" {
		log.Printf("DEBUG %s", message)
	}
}
