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

// Package cmd holds the commands for the makesite CLI
package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Post is a struct that holds information to be added into an html file
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
		for argNum := range args {
			arg := args[argNum]
			outputFile, err := save(arg)
			if err != nil {
				fmt.Printf("Error transforming %s to .html file!\n", arg)
			} else {
				fmt.Printf("Successfully created %s based on %s\n", outputFile, arg)
			}
		}
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

func save(fileName string) (outputFileName string, err error) {
	postContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
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
	newFileName := fileName[0:len(fileName)-4] + ".html"
	newFile, err := os.Create(newFileName)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(newFile, newPost)
	if err != nil {
		return "", err
	}
	return newFileName, nil
}
