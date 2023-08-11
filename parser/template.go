package parser

import (
	"bytes"
	"embed"
	"io/fs"
	"text/template"
)

type Template struct {
	data any
}

//go:embed template/*
var templateSource embed.FS

func (t *Template) ToModelTemplate() string {
	f, _ := fs.ReadFile(templateSource, "template/model.tmpl")
	mt, _ := execTmpl(string(f), t.data)
	return mt
}

func NewTemplate(data any) *Template {
	return &Template{data: data}
}

func execTmpl(s string, data any) (string, error) {
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
