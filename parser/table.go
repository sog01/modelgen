package parser

import (
	"errors"
	"strings"

	"github.com/sog01/modelgen/types"
)

/*
CREATE TABLE IF NOT EXISTS tasks (
    task_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    due_date DATE,
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB;
*/

type Table struct {
	Type    types.DDLType
	Table   string
	Columns []*Column
}

type Column struct {
	Name       string
	Type       string
	Identifier bool
	Nullable   bool
}

func ParseTable(s string) (Table, error) {
	s = strings.TrimSpace(s)
	splitted := strings.Split(s, "\n")
	if len(splitted) == 0 {
		return Table{}, errors.New("empty string")
	}
	switch typ := parseDDLType(splitted[0]); typ {
	case types.CreateTable:
		return Table{
			Type:  typ,
			Table: parseCreateTable(splitted[0]),
		}, nil
	}
	return Table{}, nil
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
