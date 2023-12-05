package prompt

import (
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

func NewSchema(graphQLStr string) (*ast.SchemaDocument, error) {
	source := &ast.Source{
		Name:  "schema.graphql",
		Input: graphQLStr,
	}

	doc, err := parser.ParseSchema(source)

	if err != nil {
		log.Println("Could not parse schema. Error:", err.Error())
		return nil, err
	}

	return doc, nil
}

func NewSchemaFromFile(filepath string) (*ast.SchemaDocument, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	graphQLStr := string(data)
	return NewSchema(graphQLStr)
}
