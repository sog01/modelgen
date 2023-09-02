package types

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type GoName struct {
	name       string
	originName string
}

func (g GoName) ToLower() GoName {
	return GoName{
		name:       strings.ToLower(string(g.name)),
		originName: g.originName,
	}
}

func (g GoName) Plural() GoName {
	return GoName{
		name:       pluralize.NewClient().Plural(string(g.name)),
		originName: g.originName,
	}
}

func (g GoName) String() string {
	return g.name
}

func (g GoName) Origin() string {
	return g.originName
}

func NewGoName(s string) GoName {
	s = strings.ReplaceAll(s, "`", "")
	c := strcase.ToCamel(s)
	return GoName{
		name:       pluralize.NewClient().Singular(c),
		originName: s,
	}
}
