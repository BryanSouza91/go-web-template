package main

import (
	"embed"
	"html/template"
	"log/slog"
	"os"
)

var (
	port string

	//go:embed templates
	tmplFS embed.FS

	//go:embed assets
	assetsFS embed.FS

	baseTmpl   *template.Template
	indexTmpl  *template.Template
	err404Tmpl *template.Template
)

func initTemplates() {
	var err error
	baseTmpl = template.Must(template.New("base").ParseFS(tmplFS, "templates/base.html"))

	indexTmpl, err = template.Must(baseTmpl.Clone()).ParseFS(tmplFS, "templates/index.html")
	if err != nil {
		slog.Error("Failed to parse index template", "error", err)
		os.Exit(1)
	}

	err404Tmpl, err = template.Must(baseTmpl.Clone()).ParseFS(tmplFS, "templates/404.html")
	if err != nil {
		slog.Error("Failed to parse 404 template", "error", err)
		os.Exit(1)
	}
}

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	initTemplates()
}
