package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	VERSION = "0.0.1"
)

func GetVersion() string {
	return VERSION
}

func GetRequestData(request *http.Response) (int, string) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer request.Body.Close()

    return request.StatusCode, strings.TrimSpace(string(body))
}
