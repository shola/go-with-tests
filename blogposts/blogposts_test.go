package blogposts_test

// best practice to put tests in their own package

import (
	"errors"
	"io/fs"
	"reflect"
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
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: go, tdd`
		secondBody = `Title: Post 2
Description: Description 2
Tags: engineering, reference`
	)
	// setup in-memory FS for testing purposes
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        "go, tdd",
	}

	assertPost(t, got, want)
}

func assertPost(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
