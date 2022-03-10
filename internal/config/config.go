package config

type Appconfig struct {
	Port              int
	TemplateDirectory string
}

func (a *Appconfig) SetPort(port int) {
	a.Port = port
}

func (a *Appconfig) SetTemplateDirectory(path string) {
	a.TemplateDirectory = path
}
