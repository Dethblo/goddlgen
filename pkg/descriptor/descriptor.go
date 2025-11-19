package descriptor

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type JsonInput struct {
	FolderName string `yaml:"folderName"`
}

type SqlOutputType struct {
	FolderName string `yaml:"folderName"`
}
type InputType struct {
	JsonInput JsonInput `yaml:"json"`
}

type OutputType struct {
	SqlOutputType JsonInput `yaml:"sql"`
}

type Descriptor struct {
	Input  InputType  `yaml:"input"`
	Output OutputType `yaml:"output"`
}

func (d *Descriptor) ReadFromYml(filePathAndName string) error {
	var data []byte
	data, err := os.ReadFile(filePathAndName)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to read file: %v", filePathAndName))
	}

	// Unmarshal yaml text to struct
	return yaml.Unmarshal(data, d)
}
