package templates

import (
	"regexp"
	"strings"
	"text/template"
)

// Plaintext returns the go template of the plaintext resume template
func Plaintext() *template.Template {
	fns := template.FuncMap{"censor": plaintextCensor}
	tmpl, err := template.New("plaintext").Funcs(fns).Parse(plaintext)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func plaintextCensor(s string) string {
	re := regexp.MustCompile(`\|\|.*?\|\|`)
	s = re.ReplaceAllStringFunc(s, func(_s string) string { return strings.Repeat("#", len(_s)-4) })
	return s
}

var plaintext = `
{{- .Header.Name | censor }}
==============================
{{ .Header.Email | censor }}

EDUCATION
==============================
{{- range $eduEntry := .EducationEntries }}
{{ $eduEntry.Degree | censor }}, {{ $eduEntry.School | censor }}, {{ $eduEntry.TimeSpan.Display | censor }}
{{ if $eduEntry.GPA }}GPA: {{ $eduEntry.GPA | censor }}{{ "\n" }}{{ end }}
{{- end }}
PROFESSIONAL EXPERIENCE
===============================
{{- range $jobEntry := .JobEntries }}
{{ $jobEntry.Employer | censor }}, {{ $jobEntry.Location | censor }}, {{ $jobEntry.TimeSpan.Display | censor }}
{{ $jobEntry.Title | censor }}
{{- range $bullet := $jobEntry.Bullets }}
* {{ $bullet | censor }} 
{{- end }}
{{ if $jobEntry.Skills }}Technologies used: {{ $jobEntry.Skills.Display | censor }}{{ "\n" }}{{ end }}
{{- end }}
SKILLS
==============================
Languages: {{ .Languages.Display | censor }}
Technologies: {{ .Technologies.Display | censor }}

PROJECTS
==============================
{{- range $project := .Projects }}
{{ $project.Name | censor }}
{{ $project.Description | censor }}
Technologies used: {{ $project.Skills.Display | censor }}{{ "\n" }}
{{- end }}`
