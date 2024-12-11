package blog

var AllPosts = Posts{
	HelloWorld,
	Testing,
}

var HelloWorld = Post{
	Title:       "Hello, World",
	Description: "First Blog Posts",
	Tags:        []string{"foo", "bar"},
	FileName:    "hello-world",
}

var Testing = Post{
	Title:       "Teting 1,2,3",
	Description: "Second Blog Posts",
	Tags:        []string{"bar", "baz"},
	FileName:    "test",
}
