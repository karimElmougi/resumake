package templates

import (
	"strings"
	"text/template"
)

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
[[- range $eduEntry := .EducationEntries ]]
\simpleeduentry{[[ $eduEntry.School ]]}{}{[[ $eduEntry.Degree ]]}{[[ $eduEntry.TimeSpan.Display ]]}
[[- end ]]




%==== Experience ====%
\header{Experience}
\vspace{1mm}

[[- range $jobEntry := .JobEntries ]]
\workentry
    {[[ $jobEntry.Employer ]]}
    {[[ $jobEntry.Title ]]}
    {[[ $jobEntry.Location ]]}
	{[[ $jobEntry.Skills.Display | escape ]]}
    {[[ $jobEntry.TimeSpan.Display ]]}
\begin{itemize}[leftmargin=\bulletmargin] \itemsep \bulletvsep
[[- range $bullet := $jobEntry.Bullets ]]
	\item [[ $bullet ]] 
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
\newcommand{\area} [2] {
    \vspace*{-9pt}
    \begin{verse}
        \textbf{#1}   #2
    \end{verse}
}

\newcommand{\lineunder} {
    \vspace*{-8pt} \\
    \hspace*{-18pt} \hrulefill \\
}

\newcommand{\header} [1] {
    {\hspace*{-18pt}\vspace*{6pt} \textsc{#1}}
    \vspace*{-6pt} \lineunder
}

%\eduentry{School name}{Location}{Degree}{GPA value}{Dates}
\newcommand*{\eduentry}[5]{
    \textbf{#1}\hfill#2\\#3 \textit{GPA: #4}\hfill#5\\
    \vspace{2mm}
}

%\simpleeduentry{School name}{Location}{Degree}{GPA value}{Dates}
\newcommand*{\simpleeduentry}[4]{
    \textbf{#1}\hfill#2\\#3 \hfill#4\\
    \vspace{2mm}
}

%\workentry{Company}{Location}{Title}{Dates}
%\newcommand*{\workentry}[4]{
%    \textbf{#1}\hfill#2\\\textit{#3}\hfill#4\\\vspace{-1mm}
%}

\newcommand*{\workentry}[5]{
    \textbf{#1 \textbar{} #2}\hfill#3\\\vspace{0.75mm}\textit{#4}\hfill#5\\\vspace{-2.5mm}
}

%\daterange{Start Date}{End Date}
\newcommand*{\daterange}[2]{
    #1\datesep #2
}

\newcommand{\bulletmargin}{10pt}
\newcommand{\bulletvsep}{-1pt}
\newcommand{\datesep}{ - }
`
