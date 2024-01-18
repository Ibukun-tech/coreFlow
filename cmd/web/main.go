package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
	"github.com/Ibukun-tech/coreFlow/pkg/Render"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFtoken string
	Flash     string
	Warning   string
	Error     string
}

var st *log.Logger

// its the Home Page handler
func main() {
	var app config.AppConfig
	fmt.Println("here")
	ts, err := Render.CreateTemplate()
	fmt.Println(err)
	// fmt.Println(ts)
	if err != nil {
		fmt.Println(err, "This error")
		log.Fatal(nil)
	}
	app.TemplateCache = ts
	fmt.Println("here 1")
	hand := Handler.NewRepo(&app)
	Handler.NewHandlers(hand)
	Render.NewTemplate(&app)
	// Handle.Repo.Home
	// fmt.Println("here 2")
	// http.HandleFunc("/", hand.Home)
	// Handler.Repo.About()
	srv := &http.Server{
		Addr:    ":2000",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	fmt.Println(err)
}
