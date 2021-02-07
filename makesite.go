package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Post struct {
	Title    string
	Contents template.HTML
}

var Green = "\033[32m"
var Bold = "\033[1m"
var Reset = "\033[0m"

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
		numSaved := saveDir(dirName)
		fmt.Printf("%s%sSuccess!%s Generated %s%d%s page(s).\n",
			Green, Bold, Reset, Bold, numSaved, Reset)
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

func saveDir(dirName string) int {
	numSaved := 0
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".txt" {
			save(file.Name())
			numSaved += 1
		}
		if file.IsDir() {
			err = filepath.Walk(file.Name(), func(path string, info os.FileInfo, err error) error {
				if err == nil && len(info.Name()) > 4 && info.Name()[len(info.Name())-4:] == ".txt" {
					save(path)
					numSaved += 1
				} else {
					return err
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return numSaved
}
