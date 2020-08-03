package templates

import (
	"regexp"
	"strings"
	"text/template"
)

// Latex returns the go template for the Latex resume template
func Latex() *template.Template {
	fns := template.FuncMap{"escape": latexEscape, "toUpper": strings.ToUpper, "censor": latexCensor}
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

func latexCensor(text string) string {
	re := regexp.MustCompile(`\|\|(.*?)\|\|`)
	text = re.ReplaceAllString(text, "\\censor{$1}")
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
    \usepackage[margin=1in]{geometry}
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
[[ if not .CensoringEnabled ]]\StopCensoring[[ end ]]

%==== Profile ====%
\vspace*{-10pt}
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
\header{Skills}
\begin{tabular}{ l l }
    Languages:    & [[ .Languages.Display | escape | censor]] \\
    Technologies: & [[ .Technologies.Display | escape | censor]] \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
[[- range $project := .Projects ]]
{\textbf{[[ $project.Name | escape | censor]]}} \textit{[[ $project.Skills.Display | escape | censor]]} \\
[[ $project.Description | censor]] \\
\vspace*{2mm}
[[ end ]]
\end{document}
`
