package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"resume/internal/resume/templates"

	"gopkg.in/yaml.v2"

	"resume/internal/resume"
)

func main() {
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
