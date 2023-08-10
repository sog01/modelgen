package types

import "fmt"

type Tag struct {
	Name       string `json:""`
	Definition string
}

func NewTag(name string) Tag {
	return Tag{
		Name:       name,
		Definition: "db",
	}
}

func (t Tag) String() string {
	return fmt.Sprintf(`%s%s:"%s"%s`, "`",
		t.Definition,
		t.Name, "`")
}
