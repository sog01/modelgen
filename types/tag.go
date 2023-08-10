package types

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

type Tag struct {
	Name       string `json:""`
	Definition string
}

func NewTag(name string) Tag {
	return Tag{
		Name:       strcase.ToLowerCamel(name),
		Definition: "db",
	}
}

func (t Tag) String() string {
	return fmt.Sprintf(`%s%s:"%s"%s`, "`",
		t.Definition,
		t.Name, "`")
}
