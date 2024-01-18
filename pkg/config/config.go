package config

import (
	"log"
	"text/template"
)

// Holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
