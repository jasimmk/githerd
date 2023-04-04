package yamlwrapper

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/careem/githerd/pkg/filewrapper"
	"gopkg.in/yaml.v2"
)

// Write Yaml file
func WriteYamlFile(filename string, data interface{}) error {

	err := filewrapper.CreateDirIfNotExists(filename)
	if err != nil {
		return err
	}
	// Create the file.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Marshal the data.
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// Write the YAML data to the file.
	_, err = file.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

// ReadYAMLFile reads a YAML file and unmarshals its data into the specified struct.
func ReadYAMLFile(filename string, data interface{}) error {
	// Read the YAML file.
	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file %s\n%w", filename, err)
	}

	// Unmarshal the YAML data into the specified struct.
	err = yaml.Unmarshal(yamlData, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML data into struct\n %w", err)
	}

	return nil
}

// PrintYaml prints the given data as YAML.
func PrintYaml(data interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlData))
	return nil
}
