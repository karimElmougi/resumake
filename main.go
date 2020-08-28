package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/karimElmougi/resumake/internal/resume"
	"github.com/karimElmougi/resumake/internal/resume/templates"

	"gopkg.in/yaml.v2"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "resumake",
		Short: "CLI tool to generate resumes",
	}

	censor := false

	rootCmd.PersistentFlags().BoolVarP(&censor, "censor", "c", false, "Enable the censoring of text surrounded by || (ex: ||first.last||@gmail.com)")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "plaintext RESUME",
		Short: "Generates a .txt resume from the provided YAML file",
		Aliases: []string{"txt", "text"},
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s, err := makeResume(args[0], templates.Plaintext(&censor))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(s)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "latex RESUME",
		Short: "Generates a .tex file to compile into a resume using a Latex distribution",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s, err := makeResume(args[0], templates.Latex(&censor))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(s)
		},
	})

	customCommand := cobra.Command{
			Use:   "custom TEMPLATE RESUME",
			Short: "Generates a .tex file to compile into a resume using a Latex distribution",
			Args: cobra.ExactArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				templateFile := args[0]
				content, err := ioutil.ReadFile(templateFile)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				delims := args[1]
				if len(delims)%2 != 0 {
					fmt.Println("Delimiters mut have an even number of characters")
					os.Exit(1)
				}

				mid := len(delims) / 2
				openDelim := delims[:mid]
				closeDelim := delims[mid:]

				tmpl, err := template.New("template").Delims(openDelim, closeDelim).Parse(string(content))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				s, err := makeResume(args[0], tmpl)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Println(s)
			},
	}

	rootCmd.AddCommand(&customCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func makeResume(resumeFilename string, tmpl *template.Template) (string, error) {
	buf, err := ioutil.ReadFile(resumeFilename)
	if err != nil {
		return "", err
	}

	r := resume.Resume{}
	err = yaml.Unmarshal(buf, &r)
	if err != nil {
		return "", err
	}

	b := &strings.Builder{}
	err = tmpl.Execute(b, r)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
