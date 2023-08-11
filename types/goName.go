package types

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type GoName string

func (g GoName) ToLower() GoName {
	return GoName(strings.ToLower(string(g)))
}

func NewGoName(s string) GoName {
	c := strcase.ToCamel(s)
	return GoName(c)
}
