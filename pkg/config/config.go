package config

import (
	"log"
	"text/template"
)

// config dont import anything from this app, but it can imported from another app files
// i can put anything i need from the app here, any kind of configuration setting
// it will be available for every package that import this particular package

// why the application wide config is so useful
// example lets add a logger so we can use it for error handling
// that allow me to write information into a log file or the terminal or write wherever we tell it to write
// it available to every part of this application that has access to the app config 
//  thats what make config is so useful

// AppCopnfig holds the application config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog log.Logger
}