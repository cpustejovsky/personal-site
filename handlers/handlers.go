package handlers

import (
	"embed"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

type Handler struct {
	http.Handler
	Renderer Renderer
}

func New() (*Handler, error) {
	router := mux.NewRouter()
	r, err := NewRenderer()
	if err != nil {
		return nil, err
	}
	handler := &Handler{
		Handler:  router,
		Renderer: *r,
	}
	//TODO: determine the purpose of this function
	//staticHandler, err := newStaticHandler()
	//if err != nil {
	//	return nil, fmt.Errorf("problem making static resources handler: %w", err)
	//}

	router.HandleFunc("/", handler.index).Methods(http.MethodGet)
	router.HandleFunc("/about", handler.about).Methods(http.MethodGet)
	router.HandleFunc("/projects", handler.projects).Methods(http.MethodGet)
	router.HandleFunc("/reading", handler.reading).Methods(http.MethodGet)
	router.HandleFunc("/resources", handler.resources).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(handler.notfound)

	return handler, nil
}

func currentYear() int {
	return time.Now().UTC().Year()
}

type Renderer struct {
	templ *template.Template
}

func NewRenderer() (*Renderer, error) {
	fm := template.FuncMap{
		"currentYear": currentYear,
	}

	templ, err := template.New("base").Funcs(fm).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &Renderer{
		templ: templ,
	}, nil
}

func (r *Renderer) RenderHTML(w io.Writer, name string, data any) error {
	return r.templ.ExecuteTemplate(w, name, data)
}

func (h *Handler) index(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) about(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) projects(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "projects.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) reading(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "reading.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) resources(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "resources.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) notfound(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "notfound.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
