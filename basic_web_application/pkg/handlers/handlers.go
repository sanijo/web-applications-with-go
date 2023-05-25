package handlers

import (
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/render"
)

// Repository is the repository type
type Repository struct {
    App *config.AppConfig
}

// Repo repository used by the handlers
var Repo *Repository

// NewRepo sets a new repository
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository {
        App: a,
    }
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
    Repo = r
}

// Home is homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.page.html")
}

// About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "about.page.html")
}

