package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		val := strings.TrimPrefix(scanner.Text(), tagName)
		return val
	}

	readMetaLineList := func(tagName string) []string {
		scanner.Scan()
		val := strings.TrimPrefix(scanner.Text(), tagName)
		return strings.Split(val, ",")
	}

	// TODO: add error handling

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        readMetaLineList(tagsSeparator),
	}, nil
}
