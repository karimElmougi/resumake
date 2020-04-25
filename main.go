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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(app)
	buf, err := ioutil.ReadFile("test_resume.yaml")
	if err != nil {
		panic(err)
	}

	r := resume.Resume{}
	err = yaml.Unmarshal(buf, &r)
	if err != nil {
		panic(err)
	}

	tmpl := templates.Latex()
	b := &strings.Builder{}
	err = tmpl.Execute(b, r)
	if err != nil {
		panic(err)
	}

	s := b.String()
	fmt.Println(s)
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
