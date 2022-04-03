package types

type Header struct {
	Date string   `yaml:"date"`
	Due  string   `yaml:"due"`
	Type string   `yaml:"type"`
	Tags []string `yaml:"tags"`
}

type TodoProperties struct {
	Date        string   `yaml:"date"`
	Due         string   `yaml:"due"`
	Tags        []string `yaml:"tags"`
	Todo        string
	TodoDetails string
}

type EntryProperties struct {
	Date         string   `yaml:"date"`
	Tags         []string `yaml:"tags"`
	Entry        string
	EntryDetails string
	FileName     string
}
