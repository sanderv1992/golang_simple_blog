package page

import (
	"net/http"

	"blog2/site"
)


type Server struct {
	*site.Renderer
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "home.tmpl", nil)
}

func (server *Server) About(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "about.tmpl", nil)
}

func (server *Server) Contacts(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "contacts.tmpl", nil)
}