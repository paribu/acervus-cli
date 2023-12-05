package generate

import (
	"fmt"
	"os"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/paribu/acervus-cli/src/settings"
	"github.com/spf13/cobra"
)

var networks []string

var GenerateSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Generates a settings template file.",
	Long: `This command generates a new settings file with a default template.
This file will serve as a configuration for your application, outlining all the
necessary settings. Users can then customize the settings to fit their requirements.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := fetchNetworks()
		if err != nil {
			return fmt.Errorf("error while fetching networks: %s", err)
		}

		return validateFields()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkExistingProjects(defaultSettingsFilepath); err != nil {
			return err
		}

		askMissingFlags()
		eventHandlers, err := getEventHandlers()
		if err != nil {
			return err
		}

		project := &settings.Project{
			Project:     projectName,
			Description: projectDesc,
			Schema:      schemaPath,
			Sources: []settings.Source{
				{
					Track:   settings.EthContract,
					Name:    contractName,
					Network: network,
					Source: settings.SourceDetail{
						Address: contractAddr,
						Abi:     abiPath,
					},
					Code: settings.CodeDetail{
						File:     contractName,
						Handlers: eventHandlers,
					},
				},
			},
		}

		err = project.ToFile(defaultSettingsFilepath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	GenerateSettingsCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "Enter a unique name for your project")
	GenerateSettingsCmd.Flags().StringVarP(&projectDesc, "projectDesc", "d", "", "Provide a brief description of your project")
	GenerateSettingsCmd.Flags().StringVarP(&network, "network", "w", "", "Select the network your smart contract is deployed on")
	GenerateSettingsCmd.Flags().StringVarP(&contractAddr, "contractAddress", "a", "", "Enter the address of the smart contract you will be indexing")
	GenerateSettingsCmd.Flags().StringVarP(&contractName, "contractName", "m", "", "Enter the name of your smart contract")
	GenerateSettingsCmd.Flags().StringVarP(&abiPath, "abiPath", "f", "", "Enter the path to your smart contract's ABI file")
	GenerateSettingsCmd.Flags().StringVarP(&schemaPath, "schemaPath", "c", "", "Enter the path to your project's GraphQL schema file")
	GenerateSettingsCmd.Flags().StringVarP(&startBlock, "startBlock", "s", "", "Enter the block number where your indexing process will begin")
	GenerateSettingsCmd.Flags().StringVarP(&endBlock, "endBlock", "e", "", "Enter the block number where your indexing process will end")
	GenerateSettingsCmd.Flags().StringVarP(&handlers, "handlers", "z", "", "Provide the handlers for your project")
}

func fetchNetworks() error {
	api := api.NewProjectManagerAPI()

	networksResponse, err := api.GetNetworks()
	if err != nil {
		return err
	}

	for _, network := range networksResponse {
		networks = append(networks, network.Name)
	}

	return nil
}

func validateFields() error {
	validators := []func() error{
		validateProjectName,
		validateNetwork,
		validateContractAddr,
		validateContractName,
		validateABIPath,
		validateSchemaPath,
		validateStartBlock,
		validateEndBlock,
		validateHandlers,
	}

	for _, validator := range validators {
		if err := validator(); err != nil {
			return err
		}
	}

	return nil
}

func checkExistingProjects(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return nil
	}

	return askOverrideSettings(filename)
}

func askOverrideSettings(filename string) error {
	isContinue := prompt.GetContinue(filename)
	if isContinue == "no" {
		return fmt.Errorf(
			"project creation cancelled by user due to file %s already exists",
			filename,
		)
	}
	return nil
}
