package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"resumake/internal/resume"
	"resumake/internal/resume/templates"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	app := cli.App{
		Name:  "resumake",
		Usage: "CLI tool to generate resumes",
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "plaintext",
				Aliases:   []string{"txt", "text"},
				Usage:     "Generates a .txt resume",
				ArgsUsage: "resume.yaml",
				Action:    makeTextResume,
			},
			&cli.Command{
				Name:      "latex",
				Usage:     "Generates a .tex file to compile into a resume using a Latex distribution",
				ArgsUsage: "resume.yaml",
				Action:    makeLatexResume,
			},
			&cli.Command{
				Name:      "custom",
				Usage:     "Renders a custom template",
				ArgsUsage: "templateFile resume.yaml",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "delimiters",
						Aliases: []string{"d", "delim"},
						Value:   "{{}}",
						Usage:   "go template delimiters",
					},
				},
				Action: makeCustomResume,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func makeCustomResume(c *cli.Context) error {
	if c.NArg() < 2 {
		cli.ShowCommandHelpAndExit(c, "custom", 1)
	}

	templateFile := c.Args().First()
	content, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	delims := c.String("delimiters")
	if len(delims)%2 != 0 {
		println("Delimiters must have an even number of characters")
		os.Exit(1)
	}

	mid := len(delims) / 2
	openDelim := delims[:mid]
	closeDelim := delims[mid:]

	tmpl, err := template.New("template").Delims(openDelim, closeDelim).Parse(string(content))
	if err != nil {
		return err
	}

	s, err := makeResume(c.Args().Get(1), tmpl)
	if err != nil {
		return err
	}

	fmt.Println(s)

	return nil
}

func makeLatexResume(c *cli.Context) error {
	if c.NArg() < 1 {
		cli.ShowCommandHelpAndExit(c, "latex", 1)
	}

	s, err := makeResume(c.Args().First(), templates.Latex())
	if err != nil {
		return err
	}

	fmt.Println(s)

	return nil
}

func makeTextResume(c *cli.Context) error {
	if c.NArg() < 1 {
		cli.ShowCommandHelpAndExit(c, "text", 1)
	}

	s, err := makeResume(c.Args().First(), templates.Plaintext())
	if err != nil {
		return err
	}

	fmt.Println(s)

	return nil
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
