package internal

import (
	"bytes"
	"net/http"
	"strings"
)

func SendRequest(request Request) (*http.Response, error) {
	data := []byte(request.Body.Content)

	apiRequest, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewBuffer(data))

	for key, value := range request.Headers {
		apiRequest.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(apiRequest)

	if err != nil {
		return nil, err
	}

	return response, nil
}
