package templates_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/karimElmougi/resumake/internal/resume/templates"
)

func TestWholeLatexTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	censor := false
	g.Expect(func() { templates.Latex(&censor) }).ToNot(Panic())
	tmpl := templates.Latex(&censor)

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
    \usepackage[margin=0.75in]{geometry}
    \textheight=10in
    \pagestyle{empty}
    \raggedright
    \usepackage{censor}
    \usepackage{fontawesome}
    \usepackage{hyperref}

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
    {\Huge JOHN SMITH}\\
    \vspace{2.5pt}
    \faEnvelope \ john.smith@gmail.com
    $|$ \faLinkedinSquare \ \href{https://linkedin.com/in/johnsmith}{linkedin.com/in/johnsmith}
    $|$ \faGithub \ \href{https://github.com/josm}{github.com/josm}\\
\end{center}




%==== Experience ====%
\header{Experience}
\vspace{1mm}

\textbf{Microsoft}
\hfill Jan. 2004 - Current\\
\textbf{Senior Software Engineer} - \textit{C\#, C++}
\vspace{-1.75mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{IBM}
\hfill Jan. 2004 - Dec. 2006\\
\textbf{Software Engineer} - \textit{Java}
\vspace{-1.75mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{SAP}
\hfill Winter 2004\\
\textbf{Software Engineer Intern} - \textit{ABAP}
\vspace{-1.75mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}




%==== Skills ====%
\header{Skills}
\vspace{1mm}
\begin{tabular}{ l l }
    Languages:    & C++, Java, C\# \\
    Technologies: & git, Docker \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
\vspace{1mm}
{\textbf{Compiler}} \textit{C\#, ANTLR, LLVM} \\
Compiles stuff \\
\vspace*{2mm}

{\textbf{Gameboy Emulator}} \textit{C++} \\
Emulates stuff \\
\vspace*{2mm}




%==== Education ====%
\header{Education}

\textbf{Georgia Institute of Technology}
\hfill\\
M.S. in Computer Science \textit{GPA: 3.9}
\hfill Jan. 2004 - Current\\
\vspace{2mm}

\textbf{University of Philadelphia}
\hfill\\
B.S. in Computer Science
\hfill Jan. 2004 - Dec. 2006\\
\vspace{2mm}

\end{document}
`
