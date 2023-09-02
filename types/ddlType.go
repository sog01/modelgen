package types

import "strings"

type DDLType int

const (
	CreateTable DDLType = iota + 1
	AlterTable
	EmptyDDL = 0
)

func NewDDLType(s string) DDLType {
	switch strings.ToLower(s) {
	case "create":
		return CreateTable
	case "alter":
		return AlterTable
	default:
		return 0
	}
}
