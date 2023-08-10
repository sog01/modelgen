package parser

import (
	"errors"
	"strings"

	"github.com/sog01/modelgen/types"
)

type Table struct {
	Name    string
	Type    types.DDLType
	Columns []*Column
}

func (c Table) Struct() *GoStruct {
	gs := &GoStruct{
		Name: types.NewGoName(c.Name),
	}
	for _, col := range c.Columns {
		if col.Identifier {
			goType, _ := types.NewGoType(col.Type, !col.NotNull)
			gs.Id = types.NewId(types.NewGoName(col.Name), goType)
		}
		gs.Properties = append(gs.Properties, col.Property())
	}
	return gs
}

type Column struct {
	Name       string
	Type       string
	Identifier bool
	NotNull    bool
}

func (c Column) Property() *Property {
	goType, _ := types.NewGoType(c.Type, !c.NotNull)
	return &Property{
		Name: types.NewGoName(c.Name),
		Type: goType,
		Tag:  types.NewTag(c.Name),
	}
}

func ParseTable(s string) (Table, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	splitted := strings.Split(s, "\n")
	if len(splitted) == 0 {
		return Table{}, errors.New("empty string")
	}
	switch typ := parseDDLType(splitted[0]); typ {
	case types.CreateTable:
		t := Table{
			Type: typ,
			Name: parseCreateTable(splitted[0]),
		}

		for _, s := range splitted[1 : len(splitted)-1] {
			t.Columns = append(t.Columns, parseColumn(s))
		}
		return t, nil
	}
	return Table{}, nil
}

func parseColumn(s string) *Column {
	s = strings.TrimSpace(s)
	splitted := strings.Split(s, " ")
	c := Column{
		Name:       splitted[0],
		Type:       sanitizeColumnType(splitted[1]),
		Identifier: strings.Contains(s, "primary key"),
		NotNull:    strings.Contains(s, "not null"),
	}

	return &c
}

func sanitizeColumnType(s string) string {
	s = strings.ReplaceAll(s, ",", "")
	bracketIndex := strings.Index(s, "(")
	if bracketIndex > -1 {
		return s[:bracketIndex]
	}

	return s
}

func parseCreateTable(s string) string {
	splitted := strings.Split(s, " ")
	return splitted[len(splitted)-2]
}

func parseDDLType(s string) types.DDLType {
	var index int
	for i, ss := range s {
		if string(ss) == " " {
			index = i
			break
		}
	}
	return types.NewDDLType(s[:index])
}
