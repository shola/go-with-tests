package blogposts_test

// best practice to put tests in their own package

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/shola/blogposts"
)

type StubFailingFS struct {
}

func (s *StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("on no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
