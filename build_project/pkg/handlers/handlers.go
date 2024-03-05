package handlers

import (
	"mygithub/pkg/config"
	"mygithub/pkg/models"
	"mygithub/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the hanlders
func NewHandlers(r *Repository) {
	Repo = r
}

// 使用指针 *http.Request 的原因是效率和约定。在 Go 中，使用指针可以避免复制整个请求结构体，特别是对于较大的结构体来说，这可以提高性能
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpt(w, "home.page.tmpl", &models.TemplateData{})
}

// about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again!!!"
	//send the data to the template
	render.RenderTmpt(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
