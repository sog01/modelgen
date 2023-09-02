package parser

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"github.com/sog01/modelgen/types"
)

func Parse(s string) ([]*Table, error) {
	p := parser.New()
	stmtNodes, _, err := p.Parse(s, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed parse sql: %v", err)
	}

	tables := []*Table{}
	for _, node := range stmtNodes {
		tables = append(tables, extract(&node))
	}

	return tables, nil
}

type visitTable struct {
	Name        string
	Type        types.DDLType
	Columns     map[string]*Column
	ColumnNames []string
}

func (v *visitTable) Enter(in ast.Node) (ast.Node, bool) {
	if v.Type != types.EmptyDDL &&
		v.Type != types.CreateTable {
		return nil, true
	}

	if table, ok := in.(*ast.CreateTableStmt); ok {
		v.Name = table.Table.Name.String()
		v.Type = types.CreateTable
	}
	if columnDef, ok := in.(*ast.ColumnDef); ok {
		columnName := columnDef.Name.String()
		var (
			autoIncrement bool
			notNull       bool
			primaryKey    bool
		)

		v.ColumnNames = append(v.ColumnNames, columnName)
		for _, opt := range columnDef.Options {
			if opt.Tp == ast.ColumnOptionPrimaryKey {
				primaryKey = true
			}
			if opt.Tp == ast.ColumnOptionAutoIncrement {
				autoIncrement = true
			}
			if opt.Tp == ast.ColumnOptionNotNull {
				notNull = true
			}
		}
		if _, ok := v.Columns[columnName]; !ok {
			v.Columns[columnName] = &Column{
				Name:          columnName,
				Type:          sanitizeColumnType(columnDef.Tp.String()),
				AutoIncrement: autoIncrement,
				Identifier:    primaryKey,
				NotNull:       notNull,
			}
		} else {
			v.Columns[columnName].Type = sanitizeColumnType(columnDef.Tp.String())
			v.Columns[columnName].Identifier = primaryKey
			v.Columns[columnName].AutoIncrement = autoIncrement
			v.Columns[columnName].NotNull = notNull
		}
	}
	if constraint, ok := in.(*ast.Constraint); ok {
		if constraint.Tp == ast.ConstraintPrimaryKey && len(constraint.Keys) > 0 {
			for _, key := range constraint.Keys {
				columnName := key.Column.String()
				if _, ok := v.Columns[columnName]; !ok {
					v.Columns[columnName] = &Column{
						Name:       columnName,
						Identifier: constraint.Tp == ast.ConstraintPrimaryKey,
					}
				} else {
					v.Columns[columnName].Identifier = constraint.Tp == ast.ConstraintPrimaryKey
				}
			}
		}
	}
	return in, false
}

func (v *visitTable) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode) *Table {
	v := &visitTable{
		Columns: make(map[string]*Column),
	}
	(*rootNode).Accept(v)

	columns := []*Column{}
	for _, col := range v.ColumnNames {
		columns = append(columns, v.Columns[col])
	}
	return &Table{
		Name:    v.Name,
		Type:    v.Type,
		Columns: columns,
	}
}

func sanitizeColumnType(s string) string {
	s = strings.ReplaceAll(s, ",", "")
	bracketIndex := strings.Index(s, "(")
	if bracketIndex > -1 {
		return s[:bracketIndex]
	}

	return s
}
