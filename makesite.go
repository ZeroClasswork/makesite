package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

type Post struct {
    Title       string
    Content     string
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
            newPost.Content += contentLines[line]
        }
    }
    fmt.Println("Title")
    fmt.Println(newPost.Title)
    fmt.Println("Content")
    fmt.Println(newPost.Content)
}
