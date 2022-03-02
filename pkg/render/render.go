package render

import "html/template"

func TemplateParser(f string) (*template.Template, error) {
	parsedFile, err := template.ParseFiles(f)
	if err != nil {
		return nil, err
	}

	return parsedFile, nil
}
