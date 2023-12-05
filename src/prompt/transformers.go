package prompt

import (
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/common"
)

func HexToAddress(val interface{}) interface{} {
	transformer := survey.TransformString(func(hex string) string {
		return ConvertHexToAddress(hex)
	})
	return transformer(val)
}

func ConvertHexToAddress(hex string) string {
	return common.HexToAddress(hex).String()
}

func ConvertStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
