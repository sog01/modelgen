package parser

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io/fs"
	"text/template"
)

type Template struct {
	data               any
	modelTemplate      string
	repositoryTemplate string
	err                []error
}

//go:embed template/*
var templateSource embed.FS

func (t *Template) ExecModelTemplate() *Template {
	f, _ := fs.ReadFile(templateSource, "template/model.tmpl")
	mt, err := execTmpl(string(f), t.data)
	if err != nil {
		t.err = append(t.err, fmt.Errorf("failed exec model template: %v", err))
	}
	t.modelTemplate = mt
	return t
}

func (t *Template) ExecRepositoryTemplate() *Template {
	f, _ := fs.ReadFile(templateSource, "template/repository.tmpl")
	rt, err := execTmpl(string(f), t.data)
	if err != nil {
		t.err = append(t.err, fmt.Errorf("failed exec repository template: %v", err))
	}
	t.repositoryTemplate = rt
	return t
}

func (t *Template) Get() (string, string) {
	return t.modelTemplate, t.repositoryTemplate
}

func (t *Template) Error() []error {
	return t.err
}

func (t *Template) Format() (string, string, error) {
	mf, err := format.Source([]byte(t.modelTemplate))
	if err != nil {
		return "", "", fmt.Errorf("failed format model template: %v", err)
	}
	rf, err := format.Source([]byte(t.repositoryTemplate))
	if err != nil {
		return "", "", fmt.Errorf("failed format repository template: %v", err)
	}

	return string(mf), string(rf), nil
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
