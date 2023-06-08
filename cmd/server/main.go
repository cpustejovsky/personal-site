package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// TODO: configure address for deployment
const addr = ":8080"

type Handler struct {
	http.Handler
}

func NewHandler() (*Handler, error) {
	router := mux.NewRouter()
	handler := &Handler{
		Handler: router,
	}
	//TODO: determine the purpoose of this function
	//staticHandler, err := newStaticHandler()
	//if err != nil {
	//	return nil, fmt.Errorf("problem making static resources handler: %w", err)
	//}

	router.HandleFunc("/", handler.index).Methods(http.MethodGet)
	router.HandleFunc("/about", handler.about).Methods(http.MethodGet)
	router.HandleFunc("/projects", handler.projects).Methods(http.MethodGet)
	router.HandleFunc("/reading", handler.reading).Methods(http.MethodGet)
	router.HandleFunc("/resources", handler.resources).Methods(http.MethodGet)
	router.HandleFunc("/mementomori", handler.mementomori).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(handler.notfound)

	return handler, nil
}

func (h *Handler) index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "index")
}

func (h *Handler) about(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "about")
}

func (h *Handler) projects(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "projects")
}

func (h *Handler) reading(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "reading")
}

func (h *Handler) resources(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "resources")
}

func (h *Handler) mementomori(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "mementomori")
}

func (h *Handler) notfound(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "notfound")
}

func main() {
	log.Printf("listening on %s", addr)
	handler, err := NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}
