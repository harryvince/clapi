package internal

import (
	"bytes"
	"net/http"
	"strings"
)

func SendRequest(request Request) (*http.Response, error) {
	data := []byte(request.Body)

	if len(request.Parameters) > 0 {
        request.Url = request.Url + "?"
		for key, value := range request.Parameters {
            request.Url = request.Url + key + "=" + value + "&"
		}
        request.Url = request.Url[:len(request.Url)-1]
	}

	apiRequest, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewBuffer(data))

	for key, value := range request.Headers {
		apiRequest.Header.Set(key, value)
	}

    apiRequest.Header.Set("User-Agent", "clapi/" + GetVersion())

	client := &http.Client{}
	response, err := client.Do(apiRequest)

	if err != nil {
		return nil, err
	}

	return response, nil
}
