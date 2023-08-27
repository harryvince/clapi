package main

import (
	"fmt"
	"io"

	"github.com/harryvince/clapi/internal"
)

func main() {
	// Read file from location
	filePath := "resources/test.yaml"
	data, err := internal.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Parse the content of the file
	content, err := internal.ParseContent(data)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("dump:\n%v\n\n", content)

	// test for my own use
	fmt.Println("Yaml Fields:")
	fmt.Print("------------------------\n")
	for i, entry := range content.Requests {
		// Data printing for request
		fmt.Print("-------Start of Yaml Printing-------\n")
		fmt.Printf("--- Request %d\n", i+1)
		fmt.Printf("--- Name: %s\n", entry.Name)
		fmt.Printf("--- Url: %s\n", entry.Url)
		fmt.Printf("--- Method: %s\n", entry.Method)
		fmt.Println("--- Parameters:")
		for key, value := range entry.Parameters {
			fmt.Printf("------ %s: %s\n", key, value)
		}
		fmt.Println("--- Headers:")
		for key, value := range entry.Headers {
			fmt.Printf("------ %s: %s\n", key, value)
		}
		fmt.Println("--- Body:")
		fmt.Printf("---- Type: %s\n", entry.Body.Type)
		fmt.Printf("---- Content: %s\n", entry.Body.Content)
		fmt.Print("-------End of Yaml Printing-------\n")
		fmt.Print("-------Start of Response Printing-------\n")
		request, err := internal.SendRequest(entry)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("---- Status Code: %s\n", request.Status)
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("---- Body: %s\n", string(body))
		defer request.Body.Close()
		fmt.Print("-------End of Response Printing-------\n")
		fmt.Print("------------------------\n")
	}

}
