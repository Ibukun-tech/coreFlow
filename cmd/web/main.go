package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ibukun-tech/coreFlow/pkg/Handler"
	model "github.com/Ibukun-tech/coreFlow/pkg/Model"
	"github.com/Ibukun-tech/coreFlow/pkg/Render"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
	"github.com/alexedwards/scs/v2"
)

// type TemplateData struct {
// 	StringMap map[string]string
// 	IntMap    map[string]int
// 	FloatMap  map[string]float64
// 	Data      map[string]interface{}
// 	CSRFtoken string
// 	Flash     string
// 	Warning   string
// 	Error     string
// }

var st *log.Logger
var session *scs.SessionManager
var app config.AppConfig

func Run() error {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.SessionStore = session

	// session
	ts, err := Render.CreateTemplate()

	if err != nil {
		return err
	}
	app.TemplateCache = ts
	hand := Handler.NewRepo(&app, &model.TemplateData{})
	Handler.NewHandlers(hand)
	Render.NewTemplate(&app)

	return nil
}

// its the Home Page handler
func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    ":2000",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	fmt.Println(err)
}
