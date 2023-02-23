use super::{list, Context};
use crate::Resume;

use std::error::Error;

use handlebars::{Handlebars, JsonRender};

const TEMPLATE: &str = include_str!("template.tex");

pub fn render(resume: Resume) -> Result<String, Box<dyn Error>> {
    let mut handlebars = Handlebars::new();

    handlebars
        .register_template_string("plaintext", TEMPLATE)
        .unwrap();

    handlebars.register_escape_fn(Box::new(latex_escape));
    handlebars.register_helper("censor", Box::new(censor));
    handlebars.register_helper("censorText", Box::new(censor_text));
    handlebars.register_helper("toUpper", Box::new(to_upper));
    handlebars.register_helper("list", Box::new(list));
    handlebars.register_helper("wrap", Box::new(wrap_in_braces));

    let ctx = handlebars::Context::wraps(Context {
        is_censored: false,
        has_skill_section: !resume.languages.is_empty() || !resume.technologies.is_empty(),
        resume,
    })
    .unwrap();

    handlebars
        .render_with_context("plaintext", &ctx)
        .map_err(Into::into)
}

fn latex_escape(input: &str) -> String {
    input
        .replace('&', r"\&")
        .replace('%', r"\%")
        .replace('$', r"\$")
        .replace('#', r"\#")
        .replace('_', r"\_")
}

fn wrap_in_braces(
    h: &handlebars::Helper,
    _: &Handlebars,
    _: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for wrap helper.",
    ))?;

    let param = &param.value().render();
    write!(out, "{{{}}}", param)?;

    Ok(())
}

fn to_upper(
    h: &handlebars::Helper,
    handlebars: &Handlebars,
    _: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for to_upper helper.",
    ))?;

    let param = handlebars.get_escape_fn()(&param.value().render().to_uppercase());

    write!(out, "{}", param)?;

    Ok(())
}

fn censor(
    h: &handlebars::Helper,
    handlebars: &Handlebars,
    c: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for censor helper.",
    ))?;

    let param = handlebars.get_escape_fn()(&param.value().render());

    write!(out, "{}", super::censor(&param, c, "######"))?;

    Ok(())
}

fn censor_text(
    h: &handlebars::Helper,
    handlebars: &Handlebars,
    c: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for censor helper.",
    ))?;

    let param = handlebars.get_escape_fn()(&param.value().render());

    write!(out, "{}", super::censor(&param, c, r"\censor{$1}"))?;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    use pretty_assertions::assert_eq;
    use serde::Serialize;

    #[test]
    fn censor_test() {
        #[derive(Serialize)]
        struct Context {
            is_censored: bool,
            value: String,
        }

        let mut handlebars = Handlebars::new();

        handlebars
            .register_template_string("test", "{{ censorText value }}")
            .unwrap();

        handlebars.register_helper("censorText", Box::new(censor_text));

        let mut ctx = Context {
            is_censored: false,
            value: String::from("||hello||"),
        };

        let hctx = handlebars::Context::wraps(&ctx).unwrap();
        assert_eq!(
            handlebars.render_with_context("test", &hctx).unwrap(),
            "hello"
        );

        ctx.is_censored = true;
        let hctx = handlebars::Context::wraps(&ctx).unwrap();
        assert_eq!(
            handlebars.render_with_context("test", &hctx).unwrap(),
            r"\censor{hello}"
        );
    }

    #[test]
    fn test() {
        let resume = crate::tests::test_resume();
        assert_eq!(render(resume).unwrap(), EXPECTED);
    }

    const EXPECTED: &str = r"\documentclass[letterpaper]{article}
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
\hfill Jan. 2009 - Current\\
\textbf{Senior Software Engineer} - \textit{C\#, C++}
\vspace{-1.75mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{IBM}
\hfill Jan. 2007 - Dec. 2008\\
\textbf{Software Engineer} - \textit{Java}
\vspace{-1.75mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{SAP}
\hfill Summer 2005\\
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
\textbf{Compiler} \textit{C\#, ANTLR, LLVM} \\
Compiles stuff \\
\vspace*{2mm}

\textbf{Gameboy Emulator} \textit{C++} \\
Emulates stuff \\
\vspace*{2mm}




%==== Education ====%
\header{Education}
\textbf{Georgia Institute of Technology}
\hfill\\
M.S. in Computer Science \textit{GPA: 3.9}
\hfill Jan. 2007 - Current\\
\vspace{2mm}

\textbf{University of Philadelphia}
\hfill\\
B.S. in Computer Science
\hfill Jan. 2004 - Dec. 2006\\
\vspace{2mm}

\end{document}
";
}
