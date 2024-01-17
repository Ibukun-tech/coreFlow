package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
	"github.com/Ibukun-tech/coreFlow/pkg/Render"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
)

// its the Home Page handler
func main() {
	var app config.AppConfig

	ts, err := Render.CreateTemplate()
	if err != nil {
		log.Fatal(nil)
	}
	app.TemplateCache = ts
	http.HandleFunc("/", Handler.Home)
	http.HandleFunc("/about", Handler.About())
	fmt.Println("Its running on port 2000")
	http.ListenAndServe(":2000", nil)
}
