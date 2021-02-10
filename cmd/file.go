/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Post struct {
	Title    string
	Contents template.HTML
}

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		save(args)
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func save(fileNames []string) {
	for fileNum := range fileNames {
		fileName := fileNames[fileNum]
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
}
