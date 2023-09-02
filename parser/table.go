package parser

import (
	"errors"

	"github.com/sog01/modelgen/types"
)

type Table struct {
	Name    string
	Type    types.DDLType
	Columns []*Column
}

func (c Table) Struct() (*GoStruct, error) {
	var ids types.Ids
	var props Properties
	for _, col := range c.Columns {
		prop, idProp := col.Property()
		if !idProp.Empty() {
			ids = append(ids, idProp)
		}
		props = append(props, prop)
	}
	if len(ids) == 0 {
		return nil, errors.New("empty id")
	}
	return NewGoStruct(types.NewGoName(c.Name), ids, props), nil
}

type Column struct {
	Name          string
	Type          string
	Identifier    bool
	AutoIncrement bool
	NotNull       bool
}

func (c Column) Property() (*Property, types.Id) {
	goType, _ := types.NewGoType(c.Type, !c.NotNull)

	prop := &Property{
		Name:       types.NewGoName(c.Name),
		Type:       goType,
		Tag:        types.NewTag(c.Name),
		OriginName: c.Name,
	}

	var id types.Id
	if c.Identifier {
		id = types.NewId(prop.Name, goType, c.AutoIncrement)
		prop.Type = id.Type
	}

	return prop, id
}
