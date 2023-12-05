package settings

var EthContract = "ethereum/contract"
var EthBlock = "ethereum/block"
var EthEvent = "ethereum/event"

type Project struct {
	Project     string   `yaml:"project"`
	Description string   `yaml:"description"`
	Schema      string   `yaml:"schema"`
	Sources     []Source `yaml:"sources"`
}

type Source struct {
	Track   string       `yaml:"track"`
	Name    string       `yaml:"name"`
	Network string       `yaml:"network"`
	Source  SourceDetail `yaml:"source"`
	Code    CodeDetail   `yaml:"code"`
}

func (s *Source) IsEthContract() bool {
	return s.Track == EthContract
}

func (s *Source) IsEthBlock() bool {
	return s.Track == EthBlock
}

type SourceDetail struct {
	Address string `yaml:"address"`
	Abi     string `yaml:"abi"`
}

type CodeDetail struct {
	File     string    `yaml:"file"`
	Handlers []Handler `yaml:"handlers"`
}

type Handler struct {
	Type       string `yaml:"type"`
	Function   string `yaml:"function"`
	Name       string `yaml:"name,omitempty"`
	CrudMode   bool   `yaml:"crudMode,omitempty"`
	StartBlock int    `yaml:"startBlock,omitempty"`
	EndBlock   int    `yaml:"endBlock,omitempty"`
	To         string `yaml:"to,omitempty"`
	From       string `yaml:"from,omitempty"`
}

func (h *Handler) IsEthEvent() bool {
	return h.Type == EthEvent
}
