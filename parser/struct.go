package parser

import (
	"github.com/sog01/modelgen/types"
)

type GoStruct struct {
	Name        types.GoName
	PackageName types.GoName
	Id          types.Id
	Properties  Properties
}

func (g *GoStruct) ToTemplate() *Template {
	return NewTemplate(g)
}

type Property struct {
	Name types.GoName
	Type types.GoType
	Tag  types.Tag
}

type Properties []*Property

func NewGoStruct(name types.GoName, id types.Id, props Properties) *GoStruct {
	return &GoStruct{
		Name:        name,
		Id:          id,
		PackageName: name.ToLower(),
		Properties:  props,
	}
}
