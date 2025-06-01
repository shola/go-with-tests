package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
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
	Body        string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		val := strings.TrimPrefix(scanner.Text(), tagName)
		return val
	}

	readBodyContent := func() string {
		// throw away the "---" line... I don't love how brittle this is when the
		// functions aren't called in a precise order
		scanner.Scan()

		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}

		return strings.TrimSuffix(buf.String(), "\n")
	}

	// TODO: add error handling

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBodyContent(),
	}, nil
}
