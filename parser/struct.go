package parser

import (
	"github.com/sog01/modelgen/types"
)

type GoStruct struct {
	Name                types.GoName
	PluralName          types.GoName
	PackageName         types.GoName
	PrivateName         types.GoName
	OriginName          string
	FirstId             types.Id
	RemainsIds          types.Ids
	Ids                 types.Ids
	Properties          Properties
	PropertiesWithoutId Properties
	Imports             Imports
}

func (g *GoStruct) ToTemplate() *Template {
	return NewTemplate(g)
}

type Property struct {
	Name       types.GoName
	Type       types.GoType
	Tag        types.Tag
	OriginName string
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

func (p Properties) ExcludeId(ids []types.Id) Properties {
	props := Properties{}
	for _, pp := range p {
		var idFound bool
		for _, id := range ids {
			if pp.Name == id.Name {
				idFound = true
				continue
			}
		}
		if idFound {
			continue
		}
		props = append(props, pp)
	}
	return props
}

type Imports struct {
	Model      types.Imports
	Repository types.Imports
}

func NewGoStruct(name types.GoName, ids types.Ids, props Properties) *GoStruct {
	g := &GoStruct{
		Name:                name,
		PluralName:          name.Plural(),
		FirstId:             ids[0],
		Ids:                 ids,
		PackageName:         name.ToLower(),
		PrivateName:         name.ToLower(),
		OriginName:          name.Origin(),
		Properties:          props,
		PropertiesWithoutId: props.ExcludeId(ids),
		Imports: Imports{
			Model: props.GetModelImports(),
			Repository: types.Imports{
				types.NewImport("github.com/jmoiron/sqlx", "").Pointer(),
				types.NewImport("strings", "").Pointer(),
				types.NewImport("context", "").Pointer(),
				types.NewImport("fmt", "").Pointer(),
			},
		},
	}
	if len(ids) > 1 {
		g.RemainsIds = ids[1:]
	}
	return g
}
