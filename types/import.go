package types

import "fmt"

type Import struct {
	Alias string
	Path  string
}

func (i Import) String() string {
	if i.Alias != "" {
		return fmt.Sprintf(`%s "%s"`, i.Alias, i.Path)
	}
	return fmt.Sprintf(`"%s"`, i.Path)
}

func NewImport(path, alias string) Import {
	return Import{
		Path:  path,
		Alias: alias,
	}
}
