package types

import "github.com/iancoleman/strcase"

type GoName string

func NewGoName(s string) GoName {
	c := strcase.ToCamel(s)
	return GoName(c)
}
