package prompt

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var MigratePlatforms = []string{"Subgraph"}
var ContractNameRegex = `^[a-zA-Z_][a-zA-Z0-9_]*$`
var defaultSchemaFilepath = "./schema.graphql"

func GetContinue(filename string) string {
	continuePc := PromptContent{
		Name:  "Continue",
		Label: fmt.Sprintf("%s file already exists. Do you want to overwrite it?", filename),
	}

	continuePcItems := []string{"yes", "no"}
	return GetSelect(continuePc, continuePcItems)
}

func GetProjectName() string {
	projectNamePc := PromptContent{
		Name:  "Project Name",
		Label: "Project name",
		Help: `Enter a unique name for your project.
This name will be used to identify your project in the other tools.`,
	}

	projectOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			survey.Required,
			survey.MinLength(3),
			survey.MaxLength(20),
		),
	}

	projectName := GetInput(projectNamePc, projectOpts)
	return projectName
}

func GetProjectDescription() string {
	projectDescPc := PromptContent{
		Name:  "Project Description",
		Label: "Project description",
		Help: `Provide a brief description of your project.
This should summarize the purpose and functionalities of your project.`,
	}
	projectDescOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			survey.Required,
		),
	}

	projectDesc := GetInput(projectDescPc, projectDescOpts)
	return projectDesc
}

func GetProjectDirectory() string {
	projectDirPc := PromptContent{
		Name:    "Directory to create",
		Label:   "Directory to create",
		Help:    `This is the directory where your project will be created.`,
		Default: "project",
	}

	projectDir := GetInput(projectDirPc, PromptOptions{})
	return projectDir
}

func GetNetwork(networks []string) string {
	networkPc := PromptContent{
		Name:  "Network",
		Label: "Network",
		Help:  `Select the network your smart contract is deployed on.`,
	}
	network := GetSelect(networkPc, networks)
	return network
}

func GetSourcePlatform() string {
	sourcePlatformPc := PromptContent{
		Name:  "Source Platform",
		Label: "Source platform",
		Help:  `Select the platform you will be migrating from.`,
	}
	sourcePlatform := GetSelect(sourcePlatformPc, MigratePlatforms)
	return sourcePlatform
}

func GetContractAddress() string {
	contractAddrPc := PromptContent{
		Name:  "Contract Address",
		Label: "Contract address",
		Help: `Enter the address of the smart contract you will be indexing.
The address should start with 0x and be in hexadecimal format.`,
	}
	contractAddrOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			survey.Required,
			EthereumAddress,
		),
		Transformer: survey.ComposeTransformers(
			HexToAddress,
		),
	}
	contractAddr := GetInput(contractAddrPc, contractAddrOpts)
	return contractAddr
}

func GetContractName() string {
	contractNamePc := PromptContent{
		Name:  "Contract Name",
		Label: "Contract name",
		Help: `Enter the name of your smart contract.
This name will be used in your queries and help identify the contract.`,
	}
	contractNameOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			survey.Required,
			Regex(ContractNameRegex, "invalid contract name"),
		),
	}

	contractName := GetInput(contractNamePc, contractNameOpts)
	return contractName
}

func GetAbiPath() string {
	abiPathPc := PromptContent{
		Name:    "ABI Path",
		Label:   "ABI path",
		Default: "./abi.json",
		Help: `Enter the path to your smart contract's Application Binary Interface (ABI) file.
The ABI defines how your project will interact with the contract.
Example path: './abi.json'.`,
	}
	abiPathOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			survey.Required,
			AbiFile,
		),
	}
	abiPath := GetInput(abiPathPc, abiPathOpts)
	return abiPath
}

func GetSchemaPath() string {
	schemaPathPc := PromptContent{
		Name:    "Schema Path",
		Label:   "Schema path",
		Default: defaultSchemaFilepath,
		Help: `Enter the path to your project's GraphQL schema file.
The schema defines the structure of your data and how it can be queried.
Example path: './schema.graphql'.`,
	}
	schemaPathOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			SchemaFile,
		),
	}
	schemaPath := GetInput(schemaPathPc, schemaPathOpts)
	return schemaPath
}

func GetStartBlock() string {
	startBlockPc := PromptContent{
		Name:    "Start Block",
		Label:   "Start block",
		Default: "0",
		Help: `Enter the block number where your indexing process will begin.
If your contract was deployed at a specific block, use that block number.
This can expedite the processing.`,
	}
	startBlockOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			GreaterThanEqualTo(0),
		),
	}
	startBlock := GetInput(startBlockPc, startBlockOpts)
	return startBlock
}

func GetEndBlock() string {
	endBlockPc := PromptContent{
		Name:    "End Block",
		Label:   "End block",
		Default: "0",
		Help: `Enter the block number where your indexing process will end.
If you want to index all blocks, leave this field empty.`,
	}
	endBlockOpts := PromptOptions{
		Validator: survey.ComposeValidators(
			GreaterThanEqualTo(0),
		),
	}
	endBlock := GetInput(endBlockPc, endBlockOpts)
	return endBlock
}

func GetSelectedEvents(events []string) []string {
	selectedEventsPc := PromptContent{
		Name:    "Select Events",
		Label:   "Select events",
		Default: strings.Join(events, "-"),
		Help: `Select the events you want to index.
You can select multiple events by pressing the space bar.`,
	}
	selectedEvents := GetMultiSelect(selectedEventsPc, events)
	return selectedEvents
}

func GetCrudEvents(events []string) []string {
	crudModeEventsPc := PromptContent{
		Name:    "CRUD Events",
		Label:   "CRUD events",
		Default: strings.Join(events, "-"),
		Help: `Select the events you want to index in CRUD mode.
You can select multiple events by pressing the space bar.`,
	}
	crudEvents := GetMultiSelect(crudModeEventsPc, events)
	return crudEvents
}
