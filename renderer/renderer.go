package renderer

import (
	"embed"
	"html/template"
	"io"
	"time"

	"github.com/cpustejovsky/personal-site/domain/blog"
	"github.com/cpustejovsky/personal-site/domain/education"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed "templates/*"
var templates embed.FS

var CurrentYear template.FuncMap = template.FuncMap{
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
}

type Renderer struct {
	templ *template.Template
}

func newParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	return parser.NewWithExtensions(extensions)
}
func New() (*Renderer, error) {
	templ, err := template.New("base").Funcs(CurrentYear).ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &Renderer{templ: templ}, nil
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
	p := blog.Post{
		Body: body,
	}
	return r.templ.ExecuteTemplate(w, "resources.gohtml", newPostVM(p))
}

// RenderIndex creates an HTML index page given a collection of posts
func (r *Renderer) RenderBlogIndex(w io.Writer, posts blog.Posts) error {
	return r.templ.ExecuteTemplate(w, "blog_index.gohtml", posts)
}

// Render renders post into HTML
func (r *Renderer) RenderBlogPost(w io.Writer, p blog.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p))
}

type postViewModel struct {
	blog.Post
	HTMLBody template.HTML
}

func newPostVM(p blog.Post) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), newParser(), nil))
	return vm
}
