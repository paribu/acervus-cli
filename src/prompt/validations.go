package prompt

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vektah/gqlparser/v2/ast"
)

func GreaterThanEqualTo(gteValue int) survey.Validator {
	return func(val interface{}) error {
		if ok, err := IsGreaterThanEqualTo(val.(string), gteValue); !ok || err != nil {
			return fmt.Errorf("value must be greater than or equal to %v", gteValue)
		}
		return nil
	}
}

func GreaterThan(gtValue int) survey.Validator {
	return func(val interface{}) error {
		if ok, err := IsGreaterThan(val.(string), gtValue); !ok || err != nil {
			return fmt.Errorf("value must be greater than %v", gtValue)
		}
		return nil
	}
}

func AbiFile(val interface{}) error {
	if _, err := ParseAbiFrom(val.(string)); err != nil {
		return fmt.Errorf("error while parsing ABI")
	}
	return nil
}

func SchemaFile(val interface{}) error {
	if _, err := ParseSchemaFrom(val.(string)); err != nil {
		return fmt.Errorf("error while parsing schema")
	}
	return nil
}

func Regex(regex string, msg string) survey.Validator {
	return func(val interface{}) error {
		if !IsRegexMatch(val.(string), regex) {
			return fmt.Errorf(msg)
		}
		return nil
	}
}

func EthereumAddress(val interface{}) error {
	if !IsEthereumAddress(val.(string)) {
		return errInvalidAddress
	}
	return nil
}

func ConvertToInt(str string) (value int, err error) {
	value, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func IsGreaterThan(str string, gtValue int) (result bool, err error) {
	value, err := ConvertToInt(str)
	if err != nil {
		return false, err
	}
	if value <= gtValue {
		return false, fmt.Errorf("value must be greater than %d", gtValue)
	}
	return true, nil
}

func IsGreaterThanEqualTo(str string, gteValue int) (result bool, err error) {
	value, err := ConvertToInt(str)
	if err != nil {
		return false, err
	}
	if value < gteValue {
		return false, fmt.Errorf("value must be greater than or equal to %d", gteValue)
	}
	return true, nil
}

func ParseAbiFrom(filePath string) (parsedAbi *abi.ABI, err error) {
	data, err := ReadFileContent(filePath)
	if err != nil {
		return nil, err
	}
	abiObj, err := abi.JSON(strings.NewReader(string(*data)))
	if err != nil {
		return nil, err
	}
	return &abiObj, nil
}

func ParseSchemaFrom(filePath string) (parsedSchema *ast.SchemaDocument, err error) {
	data, err := ReadFileContent(filePath)
	if err != nil {
		return handleGQLFileReadError(filePath, err)
	}

	parsedSchema, err = NewSchema(string(*data))
	if err != nil {
		return nil, err
	}

	return parsedSchema, nil
}

func ParseJsonFrom[T any](str string) (parsedJson T, err error) {
	err = json.Unmarshal([]byte(str), &parsedJson)
	if err != nil {
		var zero T
		return zero, err
	}
	return parsedJson, nil
}

func ReadFileContent(str string) (content *[]byte, err error) {
	file, err := os.ReadFile(str)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func IsEthereumAddress(str string) bool {
	return common.IsHexAddress(str)
}

func HasValidMinLength(str string, length int) bool {
	return !(len([]rune(str)) < length)
}

func HasValidMaxLength(str string, length int) bool {
	return !(len([]rune(str)) > length)
}

func IsInArray(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func IsRegexMatch(str string, regex string) bool {
	matched, _ := regexp.MatchString(regex, str)
	return matched
}

func handleGQLFileReadError(filePath string, err error) (*ast.SchemaDocument, error) {
	if !errors.Is(err, os.ErrNotExist) {
		// failed with unknown error
		return nil, err
	}

	// failed with (file not found) error
	if err := createEmptyFile(filePath); err != nil {
		return nil, fmt.Errorf(
			"failed to create empty graphql file with message: %s",
			err.Error(),
		)
	}

	return NewSchema("")
}

func createEmptyFile(filePath string) error {
	// create all directories in path
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	// create an empty file
	err = os.WriteFile(filePath, []byte(""), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
