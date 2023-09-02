package types

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name       string
	Definition string
}

func NewTag(name string) Tag {
	t := Tag{
		Name:       strings.ReplaceAll(name, "`", ""),
		Definition: "db",
	}
	return t
}

func (t Tag) String() string {
	name := fmt.Sprintf("%s%s%s", `"`, t.Name, `"`)
	s := fmt.Sprintf(`%s%s:%s%s`, "`", t.Definition, name, "`")
	return s
}
