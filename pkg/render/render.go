package render

import (
	"html/template"
	"log"
)

var Render *Renderers

type Renderers struct {
	templatesDir string
}

func InitRenderers(templatesDir string) {
	Render = &Renderers{
		templatesDir: templatesDir,
	}
}

func (rd Renderers) TemplateParser(f string) *template.Template {
	parsedFile, err := template.ParseFiles(rd.templatesDir + f)
	if err != nil {
		log.Println("Couldn't load or parse the file.", err)
		parsedFile = handleParsingError(err)
		return parsedFile
	}

	// apply layout
	parsedFile, err = parsedFile.ParseGlob(rd.templatesDir + "*.layout.tmpl")
	if err != nil {
		log.Println("Error while applying template", err)
		parsedFile = handleParsingError(err)
		return parsedFile
	}

	return parsedFile
}

func handleParsingError(err error) *template.Template {
	errorPage, _ := template.New("tmp").Parse("Something went wrong")
	return errorPage
}
