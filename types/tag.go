package types

import (
	"fmt"
	"html/template"
)

type Tag struct {
	Name       string
	Definition string
}

func NewTag(name string) Tag {
	t := Tag{
		Name:       name,
		Definition: "db",
	}
	return t
}

func (t Tag) String() string {
	name := fmt.Sprintf("%s%s%s", template.HTML(`"`), t.Name, template.HTML(`"`))
	s := fmt.Sprintf(`%s%s:%s%s`, "`", t.Definition, name, "`")
	return s
}
