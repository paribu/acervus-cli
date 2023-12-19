package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/paribu/acervus-cli/src/settings"
)

const (
	SkipOnConflict  = "skip"  // keep the old schema and skip the new one
	ForceOnConflict = "force" // replace the old schema with the new one
	MergeOnConflict = "merge" // merge attributes from the new schema with the old one
)

type GenerateGraphQLRequest struct {
	GqlProjectDir     string `json:"gqlProjectDir"`
	AbiFile           string `json:"abiFile"`
	GraphqlFile       string `json:"graphqlFile"`
	ConflictBehaviour string `json:"conflictBehaviour"`
}

type GenerateGraphQLResponse struct {
	Files []struct {
		Path     string `json:"path"`
		Contents string `json:"contents"`
	}
}

func (a *ProjectManagerAPI) GraphQL(gqlProjectDir, yamlFilepath string, autoSkipMode bool) (*GenerateGraphQLResponse, error) {
	yamlFile, err := settings.NewProjectFromFile(yamlFilepath)
	if err != nil {
		return nil, err
	}

	abiFile, err := os.ReadFile(yamlFile.Sources[0].Source.Abi)
	if err != nil {
		return nil, err
	}

	var schemaFileStr string
	schemaFilePath := yamlFile.Schema
	isOldSchemaExists := isFileExists(schemaFilePath)
	if isOldSchemaExists {
		// read old schema file when it exists
		schemaFile, err := os.ReadFile(yamlFile.Schema)
		if err != nil {
			return nil, err
		}

		schemaFileStr = string(schemaFile)
	} else {
		// create empty schema file when it does not exist
		err = os.WriteFile(schemaFilePath, []byte(""), 0644)
		if err != nil {
			return nil, err
		}

		schemaFileStr = ""
	}

	conflictBehavior := SkipOnConflict
	if isOldSchemaExists && !autoSkipMode {
		conflictBehavior = getConflictBehaviorChoice()
	}

	body, err := json.Marshal(GenerateGraphQLRequest{
		GqlProjectDir:     gqlProjectDir,
		AbiFile:           string(abiFile),
		GraphqlFile:       schemaFileStr,
		ConflictBehaviour: conflictBehavior,
	})
	if err != nil {
		return nil, err
	}

	resp, err := a.makeAuthenticatedAPIRequest(
		http.MethodPost,
		endpoints.generate.graphql,
		RequestData{Body: body},
	)
	if err != nil {
		return nil, err
	}

	var graphqlResp GenerateGraphQLResponse
	err = json.Unmarshal(resp, &graphqlResp)
	if err != nil {
		return nil, err
	}

	return &graphqlResp, nil
}

func getConflictBehaviorChoice() string {
	resetChoicePrompt := prompt.PromptContent{
		Name: "resetChoice",
		Label: fmt.Sprintf(
			"You already have a schema file. What do you want to do when conflicts occur? (%s/%s/%s)",
			SkipOnConflict,
			ForceOnConflict,
			MergeOnConflict,
		),
		Default: SkipOnConflict,
		Help:    "Choose 'reset' to overwrite the existing schema, 'merge' to merge with existing schema, or 'cancel' to cancel the operation.",
	}

	resetChoiceOptions := prompt.PromptOptions{
		Validator: func(val interface{}) error {
			choice := val.(string)
			if choice != SkipOnConflict && choice != ForceOnConflict && choice != MergeOnConflict {
				return fmt.Errorf(
					"invalid choice, please choose between %s, %s, or %s",
					SkipOnConflict,
					ForceOnConflict,
					MergeOnConflict,
				)
			}
			return nil
		},
	}

	return prompt.GetInput(resetChoicePrompt, resetChoiceOptions)
}

func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
