package parser

import (
	"github.com/sog01/modelgen/types"
)

type GoStruct struct {
	Name       types.GoName
	Id         types.Id
	Properties Properties
}

type Property struct {
	Name types.GoName
	Type types.GoType
	Tag  types.Tag
}

type Properties []*Property
