package Render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func ParseTemplate(w http.ResponseWriter, locat string) {
	// create my template cache
	tc, err := CreateTemplate()
	if err != nil {
		log.Fatal(err)
	}
	// Get requested template from cache
	t, ok := tc[locat]
	if !ok {
		log.Println(err)
	}
	buff := new(bytes.Buffer)
	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}
	// Render file
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}
func CreateTemplate() (map[string]*template.Template, error) {
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
