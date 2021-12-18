package importers

type Metrics struct {
	total int
}

type Importer interface {
	Import(string) (<-chan []string, error)
}
