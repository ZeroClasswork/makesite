package main

import (
	// "fmt"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Post struct {
	Title    string
	Contents string
}

func main() {
	var fileName string
	flagDescription := "a filename of a txt file to be formatted to a post"
	flag.StringVar(&fileName, "file", "first-post.txt", flagDescription)
	flag.Parse()
	save(fileName)
}

func save(fileName string) {
	postContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	newPost := new(Post)
	contentLines := strings.Split(string(postContents), "\n")
	if len(contentLines) > 0 {
		newPost.Title = contentLines[0]
	}
	for line := range contentLines {
		if line != 0 && contentLines[line] != "\n" {
			newPost.Contents += contentLines[line]
		}
	}

	tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, err := os.Create(fileName[0:len(fileName)-4] + ".html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(newFile, newPost)
	if err != nil {
		panic(err)
	}
}
