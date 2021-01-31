package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "html/template"
    "os"
)

type Post struct {
    Title       string
    Contents     string
}

func main() {
    postContents, err := ioutil.ReadFile("first-post.txt")
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
    fmt.Println("Title")
    fmt.Println(newPost.Title)
    fmt.Println("Content")
    fmt.Println(newPost.Contents)
    fmt.Println("\n\nRead in contents of first-post.txt\n\n")

    tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
    err = tmpl.Execute(os.Stdout, newPost)
    if err != nil {
        panic(err)
    }
    fmt.Println("\n\nPrinted first-post.txt with Go Templates\n\n")
    newFile, err := os.Create("first-post.html")
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(newFile, newPost)
    if err != nil {
        panic(err)
    }
    fmt.Println("\n\nWrote HTML template to first-post.html\n\n")
}
