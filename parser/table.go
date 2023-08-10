package parser

import (
	"errors"
	"strings"

	"github.com/sog01/modelgen/types"
)

type Table struct {
	Type    types.DDLType
	Table   string
	Columns []*Column
}

type Column struct {
	Name       string
	Type       string
	Identifier bool
	NotNull    bool
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
			Type:  typ,
			Table: parseCreateTable(splitted[0]),
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
