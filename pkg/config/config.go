package config

import "text/template"

// config dont import anything from this app, but it can imported from another app files
// i can put anything i need from the app here, any kind of configuration setting
// it will be available for every package that import this particular package

// AppCopnfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}