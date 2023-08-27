package internal

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type FileStructure struct {
	Requests []Request `yaml:"requests"`
}

type Request struct {
	Name       string            `yaml:"name"`
	Url        string            `yaml:"url"`
	Method     string            `yaml:"method"`
	Parameters map[string]string `yaml:"parameters"`
	Headers    map[string]string `yaml:"headers"`
	Body       string            `yaml:"body"`
}

func ParseContent(content string) (FileStructure, error) {
	data := FileStructure{}
	err := yaml.Unmarshal([]byte(content), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
