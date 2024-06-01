package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/pharsha1995/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet models.Snippet
	Snippets []models.Snippet
	Form any
	Flash string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var funcMap = template.FuncMap{"humanDate": humanDate}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	partials, err := filepath.Glob("./ui/html/partials/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{"./ui/html/base.tmpl.html", page}
		files = append(files, partials...)

		ts, err := template.New(name).Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}