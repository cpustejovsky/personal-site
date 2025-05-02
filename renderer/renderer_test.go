package renderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/cpustejovsky/personal-site/domain/blog"
	"github.com/cpustejovsky/personal-site/renderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blog.Post{
			Title: "hello world",
			Body: `# First recipe!
Welcome to my **amazing blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions.`,
			Description: "This is a description",
			// Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := renderer.New(renderer.CurrentYear)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.RenderBlogPost(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blog.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderBlogIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blog.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			// Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := renderer.New(renderer.CurrentYear)

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.RenderBlogPost(io.Discard, aPost)
	}
}
