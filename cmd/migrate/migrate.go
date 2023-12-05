package migrate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/paribu/acervus-cli/cmd/generate"
	"github.com/paribu/acervus-cli/cmd/migrate/subgraph"
	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/paribu/acervus-cli/src/settings"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate your project from a specified source to the current platform.",
	Long: `The migrate command enables users to effortlessly transfer their projects
from a designated source platform to the current system.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if sourcePlatform != "" {
			if !prompt.IsInArray(sourcePlatform, prompt.MigratePlatforms) {
				return fmt.Errorf("source platform: invalid platform %v", prompt.MigratePlatforms)
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if sourcePlatform == "" {
			sourcePlatform = prompt.GetSourcePlatform()
		}

		if sourcePlatform == "Subgraph" {
			return migrateFromSubgraph()
		}

		return nil
	},
}

func init() {
	MigrateCmd.Flags().StringVarP(&sourcePlatform, "sourcePlatform", "s", "", "Platform you will be migrating from")
	MigrateCmd.Flags().StringVarP(&projectDir, "dir", "d", "", "Directory where the project will be created")
}

func migrateFromSubgraph() error {
	if projectDir == "" {
		projectDir = prompt.GetProjectDirectory()
	}

	// TODO: remove this. this is just for testing purposes.
	if err := os.Chdir("example-subgraph"); err != nil {
		return fmt.Errorf("failed to change the working directory: %s", err)
	}

	projectFullPath := "subgraph.yaml"
	project, err := subgraph.NewSubgraphFromFile(projectFullPath)
	if err != nil {
		return fmt.Errorf("error opening subgraph file: %s", err)
	}

	ds := project.DataSources[0]
	schemaFullPath := project.Schema.File
	if _, err := prompt.ParseSchemaFrom(schemaFullPath); err != nil {
		return fmt.Errorf("schema path: invalid schema file %s", err)
	}

	abiFullPath := ds.Mapping.Abis[0].File
	if _, err := prompt.ParseAbiFrom(abiFullPath); err != nil {
		return fmt.Errorf("abi path: invalid abi file %s", err)
	}

	if _, err := os.Stat(projectDir); err != nil {
		if err := os.Mkdir(projectDir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory '%s': %v", projectDir, err)
		}
	}

	copyFile(
		abiFullPath,
		filepath.Join(projectDir, filepath.Base(abiFullPath)),
	)
	copyFile(
		schemaFullPath,
		filepath.Join(projectDir, filepath.Base(schemaFullPath)),
	)
	if err := os.Chdir(projectDir); err != nil {
		return fmt.Errorf("failed to change the working directory: %s", err)
	}

	handlers := make([]settings.Handler, 0, len(ds.Mapping.EventHandlers))
	if strings.Contains(ds.Mapping.Kind, settings.EthEvent) {
		for _, handler := range ds.Mapping.EventHandlers {
			handlers = append(handlers, settings.Handler{
				Type:       settings.EthEvent,
				Name:       strings.ReplaceAll(handler.Event, "indexed ", ""),
				Function:   handler.Handler,
				StartBlock: ds.Source.StartBlock,
			})
		}
	}
	handlersBytes, err := json.Marshal(handlers)
	if err != nil {
		return fmt.Errorf("error converting to JSON: %v", err)
	}
	handlersStr := string(handlersBytes)

	generate.GenerateSettingsCmd.Flags().Set("projectDesc", project.Description)
	generate.GenerateSettingsCmd.Flags().Set("contractName", ds.Name)
	generate.GenerateSettingsCmd.Flags().Set("contractAddress", ds.Source.Address)
	generate.GenerateSettingsCmd.Flags().Set("abiPath", filepath.Base(abiFullPath))
	generate.GenerateSettingsCmd.Flags().Set("schemaPath", filepath.Base(schemaFullPath))
	generate.GenerateSettingsCmd.Flags().Set("startBlock", fmt.Sprintf("%d", ds.Source.StartBlock))
	generate.GenerateSettingsCmd.Flags().Set("handlers", handlersStr)

	return generate.GenerateSettingsCmd.RunE(generate.GenerateSettingsCmd, []string{})
}
