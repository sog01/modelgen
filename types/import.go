package types

import (
	"fmt"
	"strings"
)

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

func (i Import) Pointer() *Import {
	return &i
}

type Imports []*Import

func (i Imports) String() string {
	imports := []string{}
	for _, ii := range i {
		imports = append(imports, ii.String())
	}

	if len(imports) > 1 {
		s := "import\n(\n" + strings.Join(imports, "\n") + "\n)\n"
		return s
	}

	return "import " + imports[0]
}

func NewImport(path, alias string) Import {
	return Import{
		Path:  path,
		Alias: alias,
	}
}
