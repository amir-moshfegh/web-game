package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Protection    bool
	Session       *scs.SessionManager
}
