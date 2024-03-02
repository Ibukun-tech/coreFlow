package Render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	model "github.com/Ibukun-tech/coreFlow/pkg/Model"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFtoken = nosurf.Token(r)
	return td
}
func ParseTemplate(w http.ResponseWriter, r *http.Request, locat string, td *model.TemplateData) {
	// create my template cache
	tc := app.TemplateCache
	// Get requested template from cache
	t, ok := tc[locat]
	if !ok {
		log.Println("could not get template from template cache")
	}
	buff := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	t.Execute(buff, td)
	// if err != nil {
	// 	log.Println(err)
	// }
	// Render file
	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}
func CreateTemplate() (map[string]*template.Template, error) {
	// fmt.Println("check temp")
	cache := map[string]*template.Template{}
	//
	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		fmt.Println(err)
	}
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.html")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}

// func ParseTemplate(w http.ResponseWriter, st string) {
// 	var templ *template.Template
// 	var err error
// 	// Check if we have already stored in the cache
// 	_, inMap := tc[st]
// 	if !inMap {
// 		fmt.Println("Caching is being created")
// 		// Its not stored in the cache so we have to put it in the cache
// 		err = CreateTemplateCache(st)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		// We already have it stored in the template
// 		log.Println("User cached template already")
// 	}
// 	templ = tc[st]
// 	templ.Execute(w, nil)
// }
// func CreateTemplateCache(t string) error {
// 	tem := []string{
// 		fmt.Sprintf("./template/%s", t),
// 		"./template/base.layout.html",
// 	}
// 	parse, err := template.ParseFiles(tem...)
// 	if err != nil {
// 		return err
// 	}
// 	tc[t] = parse

// 	return nil
// }
