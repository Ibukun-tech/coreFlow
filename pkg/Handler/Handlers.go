package Handler

import (
	"fmt"
	"log"
	"net/http"

	model "github.com/Ibukun-tech/coreFlow/pkg/Model"
	"github.com/Ibukun-tech/coreFlow/pkg/Render"
	"github.com/Ibukun-tech/coreFlow/pkg/config"
	"github.com/Ibukun-tech/coreFlow/pkg/forms"
	"github.com/justinas/nosurf"
)

// Repo the repository use by the handlers
var Repo *Repository

// This is of the type repository
type Repository struct {
	App *config.AppConfig
	Td  *model.TemplateData
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, td *model.TemplateData) *Repository {
	return &Repository{
		App: a,
		Td:  td,
	}
}

// NewHandler set the repositiry for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.SessionStore.Put(r.Context(), "remote_ip", remoteIp)
	Render.ParseTemplate(w, r, "home.page.html", &model.TemplateData{})
}

type handle func(http.ResponseWriter, *http.Request)
type JsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := model.Reservation{
		FirstName: r.Form.Get("firstName"),
		LastName:  r.Form.Get("lastName"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)
	form.Has("firstName", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		return
	}
}
func (m *Repository) AvailableJson(w http.ResponseWriter, r *http.Request) {

}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	token := nosurf.Token(r)
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	fmt.Println(start, end)
	fmt.Println(token)
	w.Write([]byte("This is the post available handler"))
}
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
		Render.ParseTemplate(w, r, "about.page.html", &model.TemplateData{
			StringMap: St,
		})
	}
}
