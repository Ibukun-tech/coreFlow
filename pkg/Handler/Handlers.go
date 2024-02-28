package Handler

import (
	"fmt"
	"net/http"

	model "github.com/Ibukun-tech/coreFlow/pkg/Model"
	"github.com/Ibukun-tech/coreFlow/pkg/Render"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
)

// Repo the repository use by the handlers
var Repo *Repository

// This is of the type repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler set the repositiry for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.SessionStore.Put(r.Context(), "remote_ip", remoteIp)
	Render.ParseTemplate(w, "home.page.html", &model.TemplateData{})
}

type handle func(http.ResponseWriter, *http.Request)

func (m *Repository) About() handle {
	return func(w http.ResponseWriter, r *http.Request) {
		e := map[string]string{
			"jjsjs": "hshsh",
		}
		fmt.Println(e)

		St := make(map[string]string)
		St["About"] = "You should work noew"
		getIp, _ := m.App.SessionStore.Get(r.Context(), "remote_ip").(byte)

		St["remote_ip"] = string(getIp)
		fmt.Println("About Page")
		Render.ParseTemplate(w, "about.page.html", &model.TemplateData{
			StringMap: St,
		})
	}
}
