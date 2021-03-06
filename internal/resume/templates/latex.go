package templates

import (
	"regexp"
	"strings"
	"text/template"
)

// Latex returns the go template for the Latex resume template
func Latex(censor *bool) *template.Template {
	fns := template.FuncMap{"escape": latexEscape, "toUpper": strings.ToUpper, "censor": func(text string) string {
		if *censor {
			re := regexp.MustCompile(`\|\|(.*?)\|\|`)
			text = re.ReplaceAllString(text, "\\censor{$1}")
		} else {
			text = strings.ReplaceAll(text, "||", "")
		}

		return text
	}}

	tmpl, err := template.New("latex").Funcs(fns).Delims("[[", "]]").Parse(latexDocument)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func latexEscape(text string) string {
	text = strings.ReplaceAll(text, "&", "\\&")
	text = strings.ReplaceAll(text, "%", "\\%")
	text = strings.ReplaceAll(text, "$", "\\$")
	text = strings.ReplaceAll(text, "#", "\\#")
	text = strings.ReplaceAll(text, "_", "\\_")
	return text
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
    \usepackage[margin=0.75in]{geometry}
    \textheight=10in
    \pagestyle{empty}
    \raggedright
    \usepackage{censor}

%%%%%%%%%%%%%%%%%%%%%%% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\lineunder} {
    \vspace*{-8pt} \\
    \hspace*{-18pt} \hrulefill \\
}

\newcommand{\header} [1] {
    {\hspace*{-18pt}\vspace*{6pt} {#1}}
    \vspace*{-6pt} \lineunder
}

%%%%%%%%%%%%%%%%%%%%%%% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%

\begin{document}
\vspace*{-40pt}

\sffamily

%==== Profile ====%
\vspace*{-25pt}
\begin{center}
    {\Huge [[ .Header.Name | toUpper | censor ]]}\\
    [[ .Header.Email | censor]]\\
\end{center}




%==== Education ====%
\header{Education}
[[ range $eduEntry := .EducationEntries ]]
\textbf{[[ $eduEntry.School | escape | censor]]}
\hfill\\
[[ $eduEntry.Degree | censor]][[- if $eduEntry.GPA ]] \textit{GPA: [[ $eduEntry.GPA | censor]]}[[ end ]]
\hfill [[ $eduEntry.TimeSpan.Display | censor]]\\
\vspace{2mm}
[[ end ]]




%==== Experience ====%
\header{Experience}
\vspace{1mm}
[[ range $jobEntry := .JobEntries ]]
\textbf{[[ $jobEntry.Employer | escape | censor]] \textbar{} [[ $jobEntry.Title | censor]]}
\hfill [[ $jobEntry.Location | censor]]\\
\vspace{0.75mm}
[[ if $jobEntry.Skills ]]\textit{[[ $jobEntry.Skills.Display | escape | censor]]}[[ "\n" ]][[ end -]]
\hfill [[ $jobEntry.TimeSpan.Display | censor]]\\
[[ if $jobEntry.Skills ]]\vspace{-2.5mm}[[ else ]]\vspace{-7mm}[[ end ]]
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
[[- range $bullet := $jobEntry.Bullets ]]
    \item [[ $bullet | escape | censor]] 
[[- end ]]
\end{itemize}
[[ end ]]



%==== Skills ====%
[[- if or .Languages .Technologies ]]
\header{Skills}
\vspace{1mm}
\begin{tabular}{ l l }
[[ if .Languages ]]    Languages:    & [[ .Languages.Display | escape | censor]] \\ [[- end ]]
[[ if .Technologies ]]    Technologies: & [[ .Technologies.Display | escape | censor]] \\ [[- end ]]
\end{tabular}
\vspace{2mm}
[[ end ]]



%==== Projects ====%
\header{Projects}
\vspace{1mm}
[[- range $project := .Projects ]]
{\textbf{[[ $project.Name | escape | censor]]}} \textit{[[ $project.Skills.Display | escape | censor]]} \\
[[ $project.Description | censor]] \\
\vspace*{2mm}
[[ end ]]
\end{document}
`
