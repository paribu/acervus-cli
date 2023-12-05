package subgraph

import (
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"gopkg.in/yaml.v2"
)

func NewSubgraphFromFile(filepath string) (*Subgraph, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return NewSubgraph(string(data))
}

func NewSubgraph(yamlStr string) (*Subgraph, error) {
	var subgraph Subgraph
	err := yaml.Unmarshal([]byte(yamlStr), &subgraph)
	if err != nil {
		return nil, err
	}

	return &subgraph, nil
}

func NewSubgraphSchema(graphQLStr string) (*ast.SchemaDocument, error) {
	source := &ast.Source{
		Name:  "subgraph.graphql",
		Input: graphQLStr,
	}

	doc, err := parser.ParseSchema(source)

	if err != nil {
		log.Println("Could not parse schema. Error:", err.Error())
		return nil, err
	}

	return doc, nil
}

func NewSubgraphSchemaFromFile(filepath string) (*ast.SchemaDocument, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	graphQLStr := string(data)
	return NewSubgraphSchema(graphQLStr)
}
