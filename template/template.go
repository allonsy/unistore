package template

import (
	t "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var templates *t.Template

func ParseTemplates() error {
	templates = t.New("")
	var err error
	parseFunc := func(path string, info os.FileInfo, err_ error) error {
		if info.IsDir() {
			return nil
		}

		if err_ != nil {
			err = err_
			return err
		} else {
			err = parseFile(path)
			return err
		}
	}
	filepath.Walk("template/html", parseFunc)
	filepath.Walk("template/css", parseFunc)
	filepath.Walk("template/js", parseFunc)
	return err
}

func parseFile(path string) error {
	templateName, err := filepath.Rel("template", path)
	if err != nil {
		return err
	}

	newTemplate := templates.New(templateName)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	contentStr := string(contents)
	_, err = newTemplate.Parse(contentStr)
	return nil
}

func GetTemplates() *t.Template {
	return templates
}

func ExecuteTemplate(w io.Writer, name string, v interface{}) error {
	return templates.ExecuteTemplate(w, name, v)
}
