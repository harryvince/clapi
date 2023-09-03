package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/harryvince/clapi/internal"
	"golang.org/x/exp/slices"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clapi <path-to-file> <--headers?>")
		return
	}

	filePath := os.Args[1]
	data, err := internal.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error getting file: ", err)
		return
	}

	content, err := internal.ParseContent(data)
	if err != nil {
		fmt.Println("Error when trying to parse content: ", err)
		os.Exit(1)
	}

	fmt.Println("--------------------------")

	var wg sync.WaitGroup
	for _, entry := range content.Requests {
		wg.Add(1)

		go func(entry internal.Request) {
			defer wg.Done()

			request, err := internal.SendRequest(entry)
			if err != nil {
				fmt.Println("Error when trying to send the request: ", err)
				os.Exit(1)
			}
			response, err := internal.GetRequestData(request)
			if err != nil {
				fmt.Println("Error when getting request data: ", err)
				os.Exit(1)
			}

			fmt.Printf("Request: %s\n", entry.Name)
			fmt.Printf("%s Response for => %s\n", strings.ToUpper(entry.Method), entry.Url)
			fmt.Printf("Status Code: %v\n", response.StatusCode)
			fmt.Printf("Body:\n%s\n", strings.ReplaceAll(response.Body, "\n", ""))
			if slices.Contains(os.Args, "--headers") {
				fmt.Println("Headers:")
				for key, value := range response.Headers {
					fmt.Printf("%s: %s\n", key, value)
				}
			}

			fmt.Println("--------------------------")

		}(entry)
	}

	wg.Wait()

}
