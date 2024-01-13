package Handler

import (
	"fmt"
	"net/http"

	"github.com/Ibukun-tech/coreFlow/pkg/Render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	Render.ParseTemplate(w, "home.html")
}

type handle func(http.ResponseWriter, *http.Request)

func About() handle {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("About Page")
		Render.ParseTemplate(w, "about.html")
	}
}
