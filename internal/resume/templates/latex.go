package templates

import "text/template"

func Latex() *template.Template {
	tmpl, err := template.New("latex").Delims("[[", "]]").Parse(latex)
	if err != nil {
		panic(err)
	}

	return tmpl
}

var latex = `
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

%%%%%%%%%%%%%%%%%%%%%%% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%

\begin{document}
\vspace*{-40pt}

    

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
	{\verb|[[ $jobEntry.Skills.Display ]]|}
    {[[ $jobEntry.TimeSpan.Display ]]}
\begin{itemize}[leftmargin=\bulletmargin] \itemsep \bulletvsep
[[- range $bullet := $jobEntry.Bullets ]]
	\item [[ $bullet ]] 
[[- end ]]
\end{itemize}
[[ end ]]


%==== Skills ====%
\header{Skills}
%C\#, C++, Rust, Go, Erlang, Python, SQL
\begin{tabular}{ l l }
	Languages:    & C\#, C++, Rust, Go, Erlang, Python, PostgreSQL \\
	Technologies: & Git, Docker, Kubernetes, Helm, Linux           \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
{\textbf{Compiler}} {\sl C\#, ANTLR, LLVM} \\
Created the compiler used to teach the Compilers class at the University of Sherbrooke
\vspace*{2mm}

{\textbf{Gameboy emulator}} {\sl Go} \\
Implemented the CPU instruction set along with emulation of most of the original Gameboy\textquotesingle{}s hardware\\
\vspace*{2mm}

{\textbf{rebar3\_todo}} {\sl Erlang} \\
Contributed to the plugin by making it recursively search through directories for TODO tags in files\\
\vspace*{2mm}

%{\textbf{Gravity Simulator}} {\sl Rust} \\
%Designed a two-dimensional sandbox for modeling the gravitational interactions of a star system %of celestial objects\\
%\vspace*{2mm}

%{\textbf{Discord bot}} {\sl Python} \\
%Developed a bot to handle resume reviewing for the CS Career Hackers non-profit
%\vspace*{2mm}

\end{document}
`
