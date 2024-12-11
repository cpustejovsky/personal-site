package blog

import (
	"os"
)

// Post is a representation of a post
type Post struct {
	Title       string
	Description string
	Body        string
	//TODO: Add date
	Tags     []string
	FileName string
}

// TODO: How to generate description and tags from markdown body
func (p *Post) GetBody(path string) error {
	path = path + "/handlers/static/posts/" + p.FileName + ".md"
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	p.Body = string(b)
	return nil
}

type Posts []Post

type PostMap map[string]Post

func NewPostMap() PostMap {
	m := make(PostMap, len(AllPosts))
	for _, post := range AllPosts {
		m[post.FileName] = post
	}
	return m
}

//TODO: pagination
