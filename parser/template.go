package parser

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
)

type Template struct {
	goStruct *GoStruct
}

//go:embed template/*
var templateSource embed.FS

func (t *Template) ToModelTemplate() string {
	f, _ := fs.ReadFile(templateSource, "template/model.tmpl")
	mt, _ := execTmpl(string(f), t.goStruct)
	return mt
}

func NewTemplate(goStruct *GoStruct) *Template {
	return &Template{goStruct}
}

func execTmpl(s string, data interface{}) (string, error) {
	t, err := template.New("").Parse(s)

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
