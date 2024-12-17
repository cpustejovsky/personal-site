package blog_test

import (
	"testing"

	"github.com/cpustejovsky/personal-site/domain/blog"
)

func TestAll(t *testing.T) {
	t.Log(blog.AllPosts)
	op := blog.OrderedPosts(blog.AllPosts...)
	t.Log(op)
	for _, post := range op {
		t.Log(post.HumanReadableDate())
	}

}
