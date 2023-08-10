package types

import "strings"

type DDLType int

const (
	CreateTable DDLType = iota
	AlterTable
)

func NewDDLType(s string) DDLType {
	switch strings.ToLower(s) {
	case "create":
		return CreateTable
	case "alter":
		return AlterTable
	default:
		return -1
	}
}
