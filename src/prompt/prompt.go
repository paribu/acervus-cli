package prompt

import (
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// PromptContent is a struct that contains the information needed to prompt the user for input.
type PromptContent struct {
	Name    string
	Label   string
	Default string
	Help    string
}

// PromptOptions is a struct that contains the options for prompting the user for input.
type PromptOptions struct {
	Validator   survey.Validator
	Transformer survey.Transformer
}

func GetInput(pc PromptContent, opts PromptOptions) string {
	q := &survey.Question{
		Name: pc.Name,
		Prompt: &survey.Input{
			Message: pc.Label,
			Default: pc.Default,
			Help:    pc.Help,
		},
		Validate:  opts.Validator,
		Transform: opts.Transformer,
	}

	var input string
	for input == "" {
		ask(q, &input)
	}
	return input
}

func GetSelect(pc PromptContent, items []string) string {
	if len(items) == 0 {
		items = []string{"Ethereum Mainnet", "Polygon Mumbai"}
	}
	q := &survey.Question{
		Name: pc.Name,
		Prompt: &survey.Select{
			Message: pc.Label,
			Options: items,
			Help:    pc.Help,
		},
	}

	var input string
	for input == "" {
		ask(q, &input)
	}
	return input
}

func GetMultiSelect(pc PromptContent, items []string) []string {
	q := &survey.Question{
		Name: pc.Name,
		Prompt: &survey.MultiSelect{
			Message: pc.Label,
			Options: items,
			Default: strings.Split(pc.Default, "-"),
			Help:    pc.Help,
		},
	}

	inputs := []string{}
	ask(q, &inputs)
	return inputs
}

func ask(q *survey.Question, result interface{}) {
	if err := survey.Ask([]*survey.Question{q}, result); err != nil {
		if err.Error() == "interrupt" {
			log.Fatal("Cancelled.")
		}

		log.Println(err.Error())
	}
}
