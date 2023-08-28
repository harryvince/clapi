package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendRequest(request Request) (*http.Response, error) {
    request.Method = strings.ToUpper(request.Method)
    
    Log(fmt.Sprintf("Sending %s request on the following url => %s", request.Method, request.Url))
	data := []byte(request.Body)

	if len(request.Parameters) > 0 {
        params := url.Values{}
		for key, value := range request.Parameters {
            Log(fmt.Sprintf("Setting Parameter => Key: %s | Value: %s", key, value))
            params.Add(key, value)
		}
        request.Url = request.Url + "?" + params.Encode()
        Log(fmt.Sprintf("Url after appending params => %s", request.Url))
	}

	apiRequest, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(data))

	for key, value := range request.Headers {
        Log(fmt.Sprintf("Setting Header => Key: %s | Value: %s", key, value))
		apiRequest.Header.Set(key, value)
	}

    apiRequest.Header.Set("User-Agent", "clapi/" + GetVersion())

	client := &http.Client{}
	response, err := client.Do(apiRequest)

	if err != nil {
		return nil, err
	}

    Log(fmt.Sprintf("Successfully sent %s request on the following url => %s", request.Method, request.Url))
	return response, nil
}
