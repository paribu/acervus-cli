package generate

import (
	"fmt"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/paribu/acervus-cli/src/settings"
)

var projectDir, settingsFilePath string
var defaultSettingsFilepath = "settings.yaml"

var projectName, projectDesc string
var network string
var contractAddr, contractName string
var abiPath string
var schemaPath string
var startBlock, endBlock string
var handlers string
var overwrite bool

func validateProjectName() error {
	if projectName != "" {
		minLength := 3
		if !prompt.HasValidMinLength(projectName, minLength) {
			return fmt.Errorf("project name: value is too short. Min length is %v", minLength)
		}
		maxLength := 20
		if !prompt.HasValidMaxLength(projectName, maxLength) {
			return fmt.Errorf("project name: value is too long. Max length is %v", maxLength)
		}
	}
	return nil
}

func validateNetwork() error {
	if network != "" {
		if !prompt.IsInArray(network, networks) {
			return fmt.Errorf("network: invalid network %v", networks)
		}
	}

	return nil
}

func validateContractAddr() error {
	if contractAddr != "" {
		if !prompt.IsEthereumAddress(contractAddr) {
			return fmt.Errorf("contract address: invalid contract address")
		}
		contractAddr = prompt.ConvertHexToAddress(contractAddr)
	}

	return nil
}

func validateContractName() error {
	if contractName != "" {
		if !prompt.IsRegexMatch(contractName, prompt.ContractNameRegex) {
			return fmt.Errorf("contract name: invalid contract name")
		}
	}

	return nil
}

func validateABIPath() error {
	if abiPath != "" {
		if _, err := prompt.ParseAbiFrom(abiPath); err != nil {
			return fmt.Errorf("abi path: invalid abi file " + err.Error())
		}
	}

	return nil
}

func validateSchemaPath() error {
	if schemaPath != "" {
		if _, err := prompt.ParseSchemaFrom(schemaPath); err != nil {
			return fmt.Errorf("schema path: invalid schema file " + err.Error())
		}
	}

	return nil
}

func validateStartBlock() error {
	if startBlock != "" {
		if ok, err := prompt.IsGreaterThanEqualTo(startBlock, 0); !ok || err != nil {
			return fmt.Errorf("start block: invalid value - %v", err)
		}
	}

	return nil
}

func validateEndBlock() error {
	if endBlock != "" {
		if ok, err := prompt.IsGreaterThanEqualTo(endBlock, 0); !ok || err != nil {
			return fmt.Errorf("end block: invalid value - %v", err)
		}
	}

	return nil
}

func validateHandlers() error {
	if handlers != "" {
		eventHandlers, err := prompt.ParseJsonFrom[[]settings.Handler](handlers)
		if err != nil {
			return fmt.Errorf("handlers: invalid handlers - %v", err)
		}
		handlerTypes := []string{settings.EthBlock, settings.EthContract, settings.EthEvent}
		for _, handler := range eventHandlers {
			if !prompt.IsInArray(handler.Type, handlerTypes) {
				return fmt.Errorf("handlers: invalid handler type")
			}
			if ok, err := prompt.IsGreaterThanEqualTo(fmt.Sprintf("%d", handler.StartBlock), 0); !ok || err != nil {
				return fmt.Errorf("handlers: invalid start block - %v", err)
			}
			if ok, err := prompt.IsGreaterThanEqualTo(fmt.Sprintf("%d", handler.EndBlock), 0); !ok || err != nil {
				return fmt.Errorf("handlers: invalid end block - %v", err)
			}
			// TODO: validate event name and function name fields with regex
		}
	}

	return nil
}

func askMissingFlags() {
	if projectName == "" {
		projectName = prompt.GetProjectName()
	}
	if projectDesc == "" {
		projectDesc = prompt.GetProjectDescription()
	}
	if network == "" {
		network = prompt.GetNetwork(networks)
	}
	if contractAddr == "" {
		contractAddr = prompt.GetContractAddress()
	}
	if contractName == "" {
		contractName = prompt.GetContractName()
	}
	if abiPath == "" {
		abiPath = prompt.GetAbiPath()
	}
	if schemaPath == "" {
		schemaPath = prompt.GetSchemaPath()
	}
}

func getEventHandlers() ([]settings.Handler, error) {
	var eventHandlers []settings.Handler
	if handlers == "" {
		abiJson, err := readABIFile(abiPath)
		if err != nil {
			return nil, fmt.Errorf("error while reading ABI file: %s", err.Error())
		}
		abiEvents := make([]string, 0, len(abiJson.Events))
		for _, event := range abiJson.Events {
			abiEvents = append(abiEvents, event.Sig)
		}
		selectedEvents := prompt.GetSelectedEvents(abiEvents)
		crudEvents := []string{}
		if len(selectedEvents) > 0 {
			crudEvents = prompt.GetCrudEvents(selectedEvents)
		}
		if startBlock == "" {
			startBlock = prompt.GetStartBlock()
		}
		if endBlock == "" {
			endBlock = prompt.GetEndBlock()
		}
		startBlockInt := prompt.ConvertStringToInt(startBlock)
		endBlockInt := prompt.ConvertStringToInt(endBlock)
		eventHandlers = createHandlers(selectedEvents, crudEvents, startBlockInt, endBlockInt)
	} else {
		var err error
		eventHandlers, err = prompt.ParseJsonFrom[[]settings.Handler](handlers)
		if err != nil {
			return nil, fmt.Errorf("error while parsing handlers: %s", err.Error())
		}
	}

	return eventHandlers, nil
}

func createHandlers(selectedEvents, crudEvents []string, startBlock, endBlock int) []settings.Handler {
	handlers := make([]settings.Handler, len(selectedEvents))
	eventUsage := make(map[string]int, len(selectedEvents))
	for i, eventSig := range selectedEvents {
		name := getEventNameFrom(eventSig)
		fnSuffix := ""
		if eventUsage[name] > 0 {
			fnSuffix = fmt.Sprintf("%d", eventUsage[name])
		}
		handlers[i] = settings.Handler{
			Type:       settings.EthEvent,
			Name:       eventSig,
			Function:   fmt.Sprintf("handle%sEvent%s", name, fnSuffix),
			CrudMode:   prompt.IsInArray(eventSig, crudEvents),
			StartBlock: startBlock,
			EndBlock:   endBlock,
		}
		eventUsage[name]++
	}
	return handlers
}

func readABIFile(abiPath string) (abi.ABI, error) {
	file, err := os.Open(abiPath)
	if err != nil {
		return abi.ABI{}, err
	}
	defer file.Close()

	abiJson, err := abi.JSON(file)
	if err != nil {
		return abi.ABI{}, err
	}

	return abiJson, nil
}

func getEventNameFrom(sig string) string {
	pattern := `^([^()]+)`
	regex := regexp.MustCompile(pattern)
	res := regex.FindString(sig)
	return res
}
