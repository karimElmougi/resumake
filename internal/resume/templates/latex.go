package templates

import (
	"strings"
	"text/template"
)

// Latex returns the go template for the Latex resume template
func Latex() *template.Template {
	tmpl, err := template.New("base").Delims("[[", "]]").Parse(latexDocument)
	if err != nil {
		panic(err)
	}

	b := &strings.Builder{}
	err = tmpl.Execute(b, map[string]string{
		"Definitions":    latexDefinitions,
		"ResumeTemplate": latexResumeTemplate,
	})
	if err != nil {
		panic(err)
	}

	fns := template.FuncMap{"escape": latexEscape}
	tmpl, err = template.New("latex").Funcs(fns).Delims("[[", "]]").Parse(b.String())
	if err != nil {
		panic(err)
	}

	return tmpl
}

func latexEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "\\&")
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "$", "\\$")
	s = strings.ReplaceAll(s, "#", "\\#")
	s = strings.ReplaceAll(s, "_", "\\_")
	return s
}

var latexDocument = `
\documentclass[letterpaper]{article}
    \usepackage{fullpage}
    \usepackage{amsmath}
    \usepackage{amssymb}
    \usepackage{textcomp}
    \usepackage{enumitem}
    \usepackage[utf8]{inputenc}
    \usepackage[T1]{fontenc}
    \usepackage[margin=1in]{geometry}
    \textheight=10in
    \pagestyle{empty}
    \raggedright

%%%%%%%%%%%%%%%%%%%%%%% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%
[[ .Definitions ]]
%%%%%%%%%%%%%%%%%%%%%%% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%

\begin{document}
\vspace*{-40pt}


[[ .ResumeTemplate ]]
\end{document}
`

var latexResumeTemplate = `
%==== Profile ====%
\vspace*{-10pt}
\begin{center}
    {\Huge \scshape {[[ .Header.Name ]]}}\\
    [[ .Header.Email ]]\\
\end{center}




%==== Education ====%
\header{Education}
[[ range $eduEntry := .EducationEntries ]]
\textbf{[[ $eduEntry.School ]]}
\hfill\\
[[ $eduEntry.Degree ]][[- if $eduEntry.GPA ]] \textit{GPA: [[ $eduEntry.GPA ]]}[[ end ]]
\hfill [[ $eduEntry.TimeSpan.Display ]]\\
\vspace{2mm}
[[ end ]]




%==== Experience ====%
\header{Experience}
\vspace{1mm}
[[ range $jobEntry := .JobEntries ]]
\textbf{[[ $jobEntry.Employer ]] \textbar{} [[ $jobEntry.Title ]]}
\hfill [[ $jobEntry.Location ]]\\
\vspace{0.75mm}
[[ if $jobEntry.Skills ]]\textit{[[ $jobEntry.Skills.Display | escape ]]}[[ "\n" ]][[ end -]]
\hfill [[ $jobEntry.TimeSpan.Display]]\\
[[ if $jobEntry.Skills ]]\vspace{-2.5mm}[[ else ]]\vspace{-7mm}[[ end ]]
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
[[- range $bullet := $jobEntry.Bullets ]]
    \item [[ $bullet | escape ]] 
[[- end ]]
\end{itemize}
[[ end ]]



%==== Skills ====%
\header{Skills}
\begin{tabular}{ l l }
    Languages:    & [[ .Languages.Display | escape ]] \\
    Technologies: & [[ .Technologies.Display | escape ]] \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
[[- range $project := .Projects ]]
{\textbf{[[ $project.Name | escape ]]}} {\sl [[ $project.Skills.Display | escape ]]} \\
[[ $project.Description ]] \\
\vspace*{2mm}
[[ end ]]
`

var latexDefinitions = `
\newcommand{\lineunder} {
    \vspace*{-8pt} \\
    \hspace*{-18pt} \hrulefill \\
}

\newcommand{\header} [1] {
    {\hspace*{-18pt}\vspace*{6pt} \textsc{#1}}
    \vspace*{-6pt} \lineunder
}
`
