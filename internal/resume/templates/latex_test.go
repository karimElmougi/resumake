package templates_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"resume/internal/resume/templates"
)

func TestWholeLatexTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { templates.Latex() }).ToNot(Panic())
	tmpl := templates.Latex()

	b := &strings.Builder{}
	err := tmpl.Execute(b, testResume)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(b.String()).To(Equal(latexResume))
}

var latexResume = `
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
	{\Huge \scshape {John Smith}}\\
	john.smith@gmail.com\\
\end{center}




%==== Education ====%
\header{Education}
\simpleeduentry{Georgia Institute of Technology}{}{M.S. in Computer Science}{Jan. 2004 - Current}
\simpleeduentry{University of Philadelphia}{}{B.S. in Computer Science}{Jan. 2004 - Dec. 2006}




%==== Experience ====%
\header{Experience}
\vspace{1mm}
\workentry
    {Microsoft}
    {Senior Software Engineer}
    {Seattle, WA}
	{C\#, C++}
    {Jan. 2004 - Current}
\begin{itemize}[leftmargin=\bulletmargin] \itemsep \bulletvsep
	\item did a thing
	\item did another thing
\end{itemize}

\workentry
    {IBM}
    {Software Engineer}
    {Seattle, WA}
	{Java}
    {Jan. 2004 - Dec. 2006}
\begin{itemize}[leftmargin=\bulletmargin] \itemsep \bulletvsep
	\item did a thing
	\item did another thing
\end{itemize}

\workentry
    {SAP}
    {Software Engineer Intern}
    {Seattle, WA}
	{ABAP}
    {Winter 2004}
\begin{itemize}[leftmargin=\bulletmargin] \itemsep \bulletvsep
	\item did a thing
	\item did another thing
\end{itemize}




%==== Skills ====%
\header{Skills}
\begin{tabular}{ l l }
	Languages:    & C++, Java, C\# \\
	Technologies: & git, Docker \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
{\textbf{Compiler}} {\sl C\#, ANTLR, LLVM} \\
Compiles stuff \\
\vspace*{2mm}

{\textbf{Gameboy Emulator}} {\sl C++} \\
Emulates stuff \\
\vspace*{2mm}


\end{document}
`
