package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeperator		= "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator = "Tags: "
)

type Post struct {
	Title	string
	Description	string
	Tags		[]string
	Body		string
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}


	title := readMetaLine(titleSeperator)
	description := readMetaLine(descriptionSeperator)
	tags := strings.Split(readMetaLine(tagsSeperator), ", ")
	body := readBody(scanner)

	return Post{
		Title: title,
		Description: description,
		Tags: tags,
		Body: body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimPrefix(buf.String(), "\n")
}