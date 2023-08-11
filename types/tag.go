package types

import (
	"fmt"
	"html/template"

	"github.com/iancoleman/strcase"
)

type Tag struct {
	Name       string
	Definition string
	HTML       template.HTML
}

func NewTag(name string) Tag {
	t := Tag{
		Name:       strcase.ToLowerCamel(name),
		Definition: "db",
	}
	t.HTML = template.HTML(t.String())
	return t
}

func (t Tag) String() string {
	name := fmt.Sprintf("%s%s%s", template.HTML(`"`), t.Name, template.HTML(`"`))
	s := fmt.Sprintf(`%s%s:%s%s`, "`", t.Definition, name, "`")
	return s
}
