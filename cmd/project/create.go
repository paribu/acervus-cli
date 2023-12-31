package project

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/paribu/acervus-cli/src/settings"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  "Create a new project with default files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID, err := createProject()
		if err != nil {
			return fmt.Errorf("error creating project: %s", err)
		}

		cmd.Printf("Created a new project with id: %s\n", projectID)
		cmd.Printf("Generating default files...\n")

		api := api.NewProjectManagerAPI()

		boilerplateRes, err := api.GenerateBoilerplate(projectID, defaultSettingsFilepath)
		if err != nil {
			return fmt.Errorf("error generating files: %s", err)
		}

		projectDir = getProjectDir()
		schemaFilepath, err := getSchemaFilepath()
		if err != nil {
			return err
		}

		if err := createProjectFiles(boilerplateRes.Files, projectDir, projectID, schemaFilepath); err != nil {
			return err
		}

		if isNpmInstalled() {
			prepareProjectFiles(projectDir, projectID)
		}

		cmd.Printf("Updated existing schema file: %s\n", schemaFilepath)
		cmd.Printf("Created files at: %s/%s/\n", projectDir, projectID)

		return nil
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&projectDir, "dir", "d", "", "Directory where the project will be created")
}

func createProject() (string, error) {
	api := api.NewProjectManagerAPI()

	resp, err := api.CreateProject(defaultSettingsFilepath)
	if err != nil {
		return "", fmt.Errorf("error when creating project: %s", err)
	}

	return resp.ProjectId, nil
}

func createProjectFiles(files []api.File, projectDir, projectID, schemaFilepath string) error {
	for _, fileInfo := range files {
		fileName := filepath.Join(projectDir, projectID, fileInfo.Path)
		fileDir := filepath.Dir(fileName)

		if filepath.Base(schemaFilepath) == filepath.Base(fileInfo.Path) {
			fileName = schemaFilepath
		}

		if _, err := os.Stat(fileDir); errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(fileDir, os.ModePerm)
		}

		if err := os.WriteFile(fileName, []byte(fileInfo.Contents), os.ModePerm); err != nil {
			return fmt.Errorf("error when writing files: %s", err)
		}
	}

	return nil
}

func prepareProjectFiles(projectDir, projectID string) error {
	fmt.Printf("Installing dependencies. This may take a while...\n")

	cmdExec := exec.Command("npm", "install")
	cmdExec.Dir = filepath.Join(projectDir, projectID)
	err := cmdExec.Run()
	if err != nil {
		return fmt.Errorf("error when installing dependencies: %s", err)
	}

	cmdExec = exec.Command("npm", "run", "format")
	cmdExec.Dir = filepath.Join(projectDir, projectID)
	err = cmdExec.Run()
	if err != nil {
		return fmt.Errorf("error when formatting files: %s", err)
	}

	fmt.Printf("Files formatted.\n")

	return nil
}

func getSchemaFilepath() (string, error) {
	settingsFile, err := settings.NewProjectFromFile(defaultSettingsFilepath)
	if err != nil {
		return "", err
	}

	return settingsFile.Schema, nil
}

func getProjectDir() string {
	if projectDir == "" {
		projectDir = prompt.GetProjectDirectory()
	}
	return projectDir
}

func isNpmInstalled() bool {
	_, err := exec.LookPath("npm")

	return err == nil
}
