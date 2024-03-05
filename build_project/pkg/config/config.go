package config

import (
	"html/template"
)

//禁止任何的import,只定义全局变量

// App Config holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	//InfoLog       *log.Logger
}
