package subgraph

type Subgraph struct {
	SpecVersion string       `yaml:"specVersion"`
	Description string       `yaml:"description"`
	Schema      Schema       `yaml:"schema"`
	DataSources []DataSource `yaml:"dataSources"`
}

type Schema struct {
	File string `yaml:"file"`
}

type DataSource struct {
	Kind    string  `yaml:"kind"`
	Name    string  `yaml:"name"`
	Network string  `yaml:"network"`
	Source  Source  `yaml:"source"`
	Mapping Mapping `yaml:"mapping"`
}

type Source struct {
	Address    string `yaml:"address"`
	Abi        string `yaml:"abi"`
	StartBlock int    `yaml:"startBlock"`
}

type Mapping struct {
	Kind          string         `yaml:"kind"`
	ApiVersion    string         `yaml:"apiVersion"`
	Language      string         `yaml:"language"`
	Entities      []string       `yaml:"entities"`
	Abis          []Abi          `yaml:"abis"`
	EventHandlers []EventHandler `yaml:"eventHandlers"`
	File          string         `yaml:"file"`
}

type Abi struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

type EventHandler struct {
	Event   string `yaml:"event"`
	Handler string `yaml:"handler"`
}
