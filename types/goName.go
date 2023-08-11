package types

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type GoName string

func (g GoName) ToLower() GoName {
	return GoName(strings.ToLower(string(g)))
}

func (g GoName) ToSingular() GoName {
	return GoName(pluralize.NewClient().Singular(string(g)))
}

func (g GoName) ToPlural() GoName {
	return GoName(pluralize.NewClient().Plural(string(g)))
}

func NewGoName(s string) GoName {
	c := strcase.ToCamel(s)
	return GoName(c)
}
