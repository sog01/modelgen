package types

import (
	"fmt"
	"html/template"

	"github.com/iancoleman/strcase"
)

type Tag struct {
	Name       string
	Definition string
}

func NewTag(name string) Tag {
	t := Tag{
		Name:       strcase.ToLowerCamel(name),
		Definition: "db",
	}
	return t
}

func (t Tag) String() string {
	name := fmt.Sprintf("%s%s%s", template.HTML(`"`), t.Name, template.HTML(`"`))
	s := fmt.Sprintf(`%s%s:%s%s`, "`", t.Definition, name, "`")
	return s
}
