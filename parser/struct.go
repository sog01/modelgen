package parser

import (
	"github.com/sog01/modelgen/types"
)

type GoStruct struct {
	Name        types.GoName
	PackageName types.GoName
	Id          types.Id
	Properties  Properties
	Imports     Imports
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

func (p Properties) GetModelImports() types.Imports {
	modelImportExist := make(map[string]struct{})
	modelImports := types.Imports{}
	for _, pp := range p {
		typ, needImport := pp.Type.Import()
		if !needImport {
			continue
		}
		if _, ok := modelImportExist[typ.String()]; !ok {
			modelImports = append(modelImports, &typ)
		}
		modelImportExist[typ.String()] = struct{}{}
	}

	return modelImports
}

type Imports struct {
	Model      types.Imports
	Repository types.Imports
}

func NewGoStruct(name types.GoName, id types.Id, props Properties) *GoStruct {
	return &GoStruct{
		Name:        name,
		Id:          id,
		PackageName: name.ToLower(),
		Properties:  props,
		Imports: Imports{
			Model: props.GetModelImports(),
		},
	}
}
