package render

import (
	"fmt"
	"html/template"
)

func TemplateParser(f string) *template.Template {
	parsedFile, err := template.ParseFiles(f)
	if err != nil {
		fmt.Println("Couldn't load or parse the file.")
		return template.New("Something went wrong")
	}

	return parsedFile
}
