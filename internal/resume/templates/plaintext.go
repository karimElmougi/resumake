package templates

import (
	"regexp"
	"strings"
	"text/template"
)

var censoringEnabled bool = false

// Plaintext returns the go template of the plaintext resume template
func Plaintext() *template.Template {
	fns := template.FuncMap{"censor": plaintextCensor, "setCensoringEnabled": setCensoringEnabled}
	tmpl, err := template.New("plaintext").Funcs(fns).Parse(plaintext)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func setCensoringEnabled(b bool) string {
	censoringEnabled = b
	return ""
}

func plaintextCensor(text string) string {
	if censoringEnabled {
		re := regexp.MustCompile(`\|\|.*?\|\|`)
		text = re.ReplaceAllStringFunc(text, func(s string) string {
			return strings.Repeat("#", len(s)-4)
		})
	} else {
		text = strings.ReplaceAll(text, "||", "")
	}
	return text
}

var plaintext = `
{{- .CensoringEnabled | setCensoringEnabled -}}
{{- .Header.Name | censor }}
{{ .Header.Email | censor }}
==============================

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
