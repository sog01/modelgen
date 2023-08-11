package parser

import (
	"strings"

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

func (p Properties) GetModelImports() ModelImports {
	modelImportExist := make(map[string]struct{})
	modelImports := ModelImports{}
	for _, pp := range p {
		typ, needImport := pp.Type.Import()
		if !needImport {
			continue
		}
		if _, ok := modelImportExist[typ.String()]; !ok {
			modelImports = append(modelImports, typ)
		}
		modelImportExist[typ.String()] = struct{}{}
	}

	return modelImports
}

type Imports struct {
	Model      ModelImports
	Repository RepositoryImports
}

type ModelImports []types.Import

func (i ModelImports) String() string {
	imports := []string{}
	for _, ii := range i {
		imports = append(imports, ii.String())
	}

	if len(imports) > 1 {
		s := "import\n(\n" + strings.Join(imports, "\n") + ")\n"
		return s
	}

	return "import " + imports[0]
}

type RepositoryImports []types.Import

func (i RepositoryImports) String() string {
	imports := []string{}
	for _, ii := range i {
		imports = append(imports, ii.String())
	}

	if len(imports) > 1 {
		return "(\n" + strings.Join(imports, "\n") + ")\n"
	}

	return imports[0]
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
