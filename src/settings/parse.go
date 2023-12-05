package settings

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func NewProjectFromFile(filepath string) (*Project, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return NewProject(string(data))
}

func NewProject(yamlStr string) (*Project, error) {
	var project Project
	err := yaml.Unmarshal([]byte(yamlStr), &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *Project) ToFile(filename string) error {
	fileBytes, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Errorf("error while creating YAML: %s", err.Error())
	}

	err = os.WriteFile(filename, fileBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error while writing file: %s", err.Error())
	}

	return nil
}

func (p *Project) ToString() (string, error) {
	projectYaml, err := yaml.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(projectYaml), nil
}
