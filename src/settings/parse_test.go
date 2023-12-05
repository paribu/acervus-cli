package settings_test

import (
	"testing"

	"github.com/paribu/acervus-cli/src/settings"
)

func TestNewProjectFromFile(t *testing.T) {
	filepath := "../../sync_service/project/settings.yaml"

	project, err := settings.NewProjectFromFile(filepath)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if project.Project != "Example" {
		t.Fatalf("Expected project name to be 'Example', got '%s'", project.Project)
	}

	expectedSchemaPath := "./schema.graphql"
	if project.Schema != expectedSchemaPath {
		t.Fatalf("Expected schema path in schema.Project to be '%s', got '%s'", expectedSchemaPath, project.Schema)
	}

}
