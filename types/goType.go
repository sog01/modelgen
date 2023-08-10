package types

import (
	"errors"
	"strings"
)

type GoType int

const (
	String GoType = iota
	Int
	Int8
	Int64
	Float
	Float64
	Time
	Bool
	Decimal

	NullString
	NullInt
	NullInt8
	NullInt64
	NullFloat
	NullFloat64
	NullTime
	NullBool
	NullDecimal
)

var imports = map[GoType]string{
	Time:        "time",
	NullString:  "gopkg.in/guregu/null.v4",
	NullInt:     "gopkg.in/guregu/null.v4",
	NullInt8:    "gopkg.in/guregu/null.v4",
	NullInt64:   "gopkg.in/guregu/null.v4",
	NullFloat:   "gopkg.in/guregu/null.v4",
	NullFloat64: "gopkg.in/guregu/null.v4",
	NullTime:    "gopkg.in/guregu/null.v4",
	NullBool:    "gopkg.in/guregu/null.v4",
	NullDecimal: "gopkg.in/guregu/null.v4",
}

func (g GoType) Import() string {
	return imports[g]
}

func (g GoType) String() string {
	switch g {
	case String:
		return "string"
	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int64:
		return "int64"
	case Float:
		return "float"
	case Float64:
		return "float64"
	case Time:
		return "time.Time"
	case Bool:
		return "bool"
	case Decimal:
		return "float64"
	case NullString:
		return "null.String"
	case NullInt:
		return "null.Int"
	case NullInt8:
		return "null.Int"
	case NullInt64:
		return "null.Int"
	case NullFloat:
		return "null.Float"
	case NullFloat64:
		return "null.Float"
	case NullTime:
		return "null.Time"
	case NullBool:
		return "null.Bool"
	case NullDecimal:
		return "null.Float"
	default:
		return "unknown"
	}
}

func NewGoType(s string, nullable bool) (GoType, error) {
	if nullable {
		return newNullableGoType(strings.ToLower(s))
	}
	switch strings.ToLower(s) {
	case "bigint":
		return Int64, nil
	case "int":
		return Int, nil
	case "text", "varchar", "enum", "char", "longtext", "mediumblob":
		return String, nil
	case "float":
		return Float64, nil
	case "tinyint":
		return Int64, nil
	case "datetime", "date", "timestamp":
		return Time, nil
	case "decimal":
		return Float64, nil
	default:
		return -1, errors.New("unknown type")
	}
}

func newNullableGoType(s string) (GoType, error) {
	switch s {
	case "bigint":
		return NullInt64, nil
	case "int":
		return NullInt, nil
	case "text", "varchar", "enum", "char", "longtext", "mediumblob":
		return NullString, nil
	case "float":
		return NullFloat64, nil
	case "tinyint":
		return NullInt64, nil
	case "datetime", "date", "timestamp":
		return NullTime, nil
	case "decimal":
		return NullDecimal, nil
	default:
		return -1, errors.New("unknown type")
	}
}
