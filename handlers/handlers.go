package handlers

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cpustejovsky/personal-site/domain/education"
	"github.com/cpustejovsky/personal-site/domain/lifetogether"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"time"
)

const ResourcesURl = "https://dev.to/api/articles/281175"

var (
	//go:embed "templates/*"
	templates embed.FS
	//go:embed "static/*"
	static embed.FS
)

type Renderer struct {
	templ *template.Template
}

type Handler struct {
	http.Handler
	Renderer Renderer
}

func New() (*Handler, error) {
	r, err := NewRenderer()
	if err != nil {
		return nil, err
	}

	staticHandler, err := newStaticHandler()
	if err != nil {
		return nil, fmt.Errorf("creating static handler failed:\t%w", err)
	}

	router := http.NewServeMux()
	handler := &Handler{
		Handler:  router,
		Renderer: *r,
	}
	router.HandleFunc("/", handler.index)
	router.HandleFunc("/about", handler.about)
	router.HandleFunc("/education", handler.education)
	router.HandleFunc("/resources", handler.resources)
	router.HandleFunc("/ltc", handler.ltc)
	router.HandleFunc("/ltc/calculate", handler.updateltc)
	router.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	return handler, nil
}

func NewRenderer() (*Renderer, error) {
	fm := template.FuncMap{
		"currentYear": func() int {
			return time.Now().UTC().Year()
		},
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

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notfound(w, r)
		return
	}
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

func (h *Handler) ltc(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "ltc.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateltc(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var dateDating, dateMarried time.Time
	yourName := r.PostForm.Get("yourName")
	otherName := r.PostForm.Get("otherName")
	yourBirthday, err := time.Parse(time.DateOnly, r.PostForm.Get("yourBirthday"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	otherBirthday, err := time.Parse(time.DateOnly, r.PostForm.Get("otherBirthday"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	dateMet, err := time.Parse(time.DateOnly, r.PostForm.Get("dateMet"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.PostForm.Get("dateDating") != "" {
		dateDating, err = time.Parse(time.DateOnly, r.PostForm.Get("dateDating"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	if r.PostForm.Get("dateMarried") != "" {
		dateMarried, err = time.Parse(time.DateOnly, r.PostForm.Get("dateMarried"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	in := lifetogether.Input{
		YourName:      yourName,
		OtherName:     otherName,
		YourBirthday:  yourBirthday,
		OtherBirthday: otherBirthday,
		DateMet:       dateMet,
		DateDating:    &dateDating,
		DateMarried:   &dateMarried,
	}
	out, err := lifetogether.CalculateNow(in)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = h.Renderer.RenderHTML(w, "ltc.gohtml", out)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) education(w http.ResponseWriter, _ *http.Request) {
	categories := education.GetContinuingEducationCategories(education.ContinuingEducationList)
	err := h.Renderer.RenderHTML(w, "educationlist.gohtml", categories)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) resources(w http.ResponseWriter, _ *http.Request) {
	body, err := GetResourcesPage()
	if err != nil {
		log.Println("error getting resource page", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = h.Renderer.RenderHTML(w, "resources.gohtml", template.HTML(body))
	if err != nil {
		log.Println("error rendering resource page", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetResourcesPage() (string, error) {
	res, err := http.Get(ResourcesURl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var m map[string]any
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	body, ok := m["body_html"]
	if !ok {
		return "", errors.New("body not found")
	}
	htmlbody, ok := body.(string)
	if !ok {
		return "", nil
	}
	return htmlbody, nil
}

func (h *Handler) notfound(w http.ResponseWriter, _ *http.Request) {
	err := h.Renderer.RenderHTML(w, "notfound.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func newStaticHandler() (http.Handler, error) {
	lol, err := fs.Sub(static, "static")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(lol)), nil
}
