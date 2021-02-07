package main

import (
	"flag"
	// "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Post struct {
	Title    string
	Contents template.HTML
}

func main() {
	var fileName string
	var dirName string
	fileFlagDescription := "a filename of a txt file to be formatted to a post"
	flag.StringVar(&fileName, "file", "", fileFlagDescription)
	dirFlagDescription := "a directory with txt files to be formatted to posts"
	flag.StringVar(&dirName, "dir", "", dirFlagDescription)
	flag.Parse()
	if fileName != "" {
		save(fileName)
	}
	if dirName != "" {
		files, err := ioutil.ReadDir(dirName)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.Name()[len(file.Name())-4:] == ".txt" {
				// fileContents, err := ioutil.ReadFile(file.Name())
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// contentLines := strings.Split(string(fileContents), "\n")
				// for line := range contentLines {
				// 	fmt.Println(contentLines[line])
				// }
				save(file.Name())
			}
		}
	}
}

func save(fileName string) {
	postContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	newPost := new(Post)
	contentLines := strings.Split(string(postContents), "\n")
	if len(contentLines) > 0 {
		newPost.Title = contentLines[0]
	}
	for line := range contentLines {
		if line != 0 && contentLines[line] != "\n" {
			newPost.Contents += template.HTML("<p>" + contentLines[line] + "</p>\n")
		}
	}

	tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile, err := os.Create(fileName[0:len(fileName)-4] + ".html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(newFile, newPost)
	if err != nil {
		log.Fatal(err)
	}
}
