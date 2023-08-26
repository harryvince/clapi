package main

import (
	"fmt"
	"log"

	"github.com/harryvince/clapi/internal"
	"gopkg.in/yaml.v3"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type FileStructure struct {
	Requests []Request `yaml:"requests"`
}

type Request struct {
	Name       string            `yaml:"name"`
	Url        string            `yaml:"url"`
	Type       string            `yaml:"type"`
	Parameters map[string]string `yaml:"parameters"`
	Headers    map[string]string `yaml:"headers"`
	Body       Body              `yaml:"body"`
}

type Body struct {
	Type    string `yaml:"type"`
	Content string `yaml:"content"`
}

func main() {
	content := FileStructure{}

	// Read file from location
	filePath := "resources/test.yaml"
	data, err := internal.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = yaml.Unmarshal([]byte(data), &content)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("dump:\n%v\n\n", content)

	// test for my own use
	fmt.Println("Yaml Fields:")
	fmt.Print("------------------------\n")
	for i, entry := range content.Requests {
		fmt.Printf("--- Request %d\n", i+1)
		fmt.Printf("--- Name: %s\n", entry.Name)
		fmt.Printf("--- Url: %s\n", entry.Url)
		fmt.Printf("--- Type: %s\n", entry.Type)
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
		fmt.Print("------------------------\n")
	}

}
