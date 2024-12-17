package blog

import (
	"fmt"
	"os"
	"time"
)

// Post is a representation of a post
type Post struct {
	Title       string
	Description string
	Body        string
	//TODO: add tags in when ready
	Date     time.Time
	FileName string
}

func (p Post) HumanReadableDate() string {
	return fmt.Sprintf("%d %s %d", p.Date.Year(), p.Date.Month().String(), p.Date.Day())
}

// TODO: How to generate description and tags from markdown body
func (p Post) GetBody(path string) (Post, error) {
	path = path + "/handlers/static/posts/" + p.FileName + ".md"
	b, err := os.ReadFile(path)
	if err != nil {
		return p, err
	}
	p.Body = string(b)
	return p, nil
}

type PostMap map[string]Post

func NewPostMap() PostMap {
	m := make(PostMap, len(AllPosts))
	for _, post := range AllPosts {
		m[post.FileName] = post
	}
	return m
}

//TODO: pagination
