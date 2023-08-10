package parser

import (
	"github.com/sog01/modelgen/types"
)

type GoStruct struct {
	Name       string
	Id         types.Id
	Properties Properties
}

type Property struct {
	Name string
	Type types.GoType
	Tag  types.Tag
}

type Properties []*Property
