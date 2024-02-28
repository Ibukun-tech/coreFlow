package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	SessionStore  *scs.SessionManager
}
