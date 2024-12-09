package handlers

import (
	"embed"
	"html/template"
	"io"
	"time"

	"github.com/cpustejovsky/personal-site/domain/education"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

var CurrentYear template.FuncMap = template.FuncMap{
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
}

type Renderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewRenderer() (*Renderer, error) {
	templ, err := template.New("base").Funcs(CurrentYear).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &Renderer{templ: templ, mdParser: parser}, nil
}

func (r *Renderer) RenderHTML(w io.Writer, name string, data any) error {
	return r.templ.ExecuteTemplate(w, name, data)
}

func (r *Renderer) RenderIndex(w io.Writer) error {
	return r.RenderHTML(w, "index.gohtml", nil)
}

func (r *Renderer) RenderAbout(w io.Writer) error {
	return r.RenderHTML(w, "about.gohtml", nil)
}

// TODO: naming is bad at the least
func (r *Renderer) RenderEducation(w io.Writer) error {
	categories := education.GetContinuingEducationCategories(education.ContinuingEducationList)
	return r.RenderHTML(w, "educationlist.gohtml", categories)
}

func (r *Renderer) RenderNotFound(w io.Writer) error {
	return r.RenderHTML(w, "notfound.gohtml", nil)
}

func (r *Renderer) RenderLTC(w io.Writer, input any) error {
	return r.RenderHTML(w, "ltc.gohtml", input)
}

// Render renders post into HTML
func (r *Renderer) RenderResourcePage(w io.Writer, body string) error {
	p := Post{
		Body: body,
	}
	return r.templ.ExecuteTemplate(w, "resources.gohtml", newPostVM(p, r))
}

// RenderIndex creates an HTML index page given a collection of posts
func (r *Renderer) RenderBlogIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "blog_index.gohtml", posts)
}

// Render renders post into HTML
func (r *Renderer) RenderBlogPost(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *Renderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
