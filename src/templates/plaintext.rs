use crate::Resume;

use std::error::Error;

use handlebars::{Handlebars, JsonRender};
use regex::Regex;
use serde::Serialize;

const TEMPLATE: &str = include_str!("template.txt");

pub fn render(resume: Resume) -> Result<String, Box<dyn Error>> {
    let mut handlebars = Handlebars::new();

    handlebars
        .register_template_string("plaintext", TEMPLATE)
        .unwrap();

    handlebars.register_helper("censor", Box::new(censor));
    handlebars.register_helper("list", Box::new(list));

    #[derive(Serialize)]
    struct Context {
        is_censored: bool,
        has_skill_section: bool,

        #[serde(flatten)]
        resume: Resume,
    }

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

fn censor(
    h: &handlebars::Helper,
    _: &Handlebars,
    c: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for censor helper.",
    ))?;

    let param = param.value().render();

    let is_censored = c.data().get("is_censored").unwrap().as_bool().unwrap();
    if is_censored {
        let re = Regex::new(r"\|\|.*?\|\|").unwrap();
        write!(out, "{}", re.replace_all(&param, "######"))?;
    } else {
        write!(out, "{}", param.replace("||", ""))?;
    }

    Ok(())
}

fn list(
    h: &handlebars::Helper,
    _: &Handlebars,
    _: &handlebars::Context,
    _: &mut handlebars::RenderContext,
    out: &mut dyn handlebars::Output,
) -> Result<(), handlebars::RenderError> {
    let param = h.param(0).ok_or(handlebars::RenderError::new(
        "Param 0 is required for list helper.",
    ))?;

    let param = param
        .value()
        .as_array()
        .ok_or(handlebars::RenderError::new(
            "Param 0 to list helper must be array.",
        ))?;

    let output = param
        .iter()
        .map(JsonRender::render)
        .collect::<Vec<_>>()
        .join(", ");
    write!(out, "{output}")?;

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    use pretty_assertions::assert_eq;

    #[test]
    fn censor_test() {
        #[derive(Serialize)]
        struct Context {
            is_censored: bool,
            value: String,
        }

        let mut handlebars = Handlebars::new();

        handlebars
            .register_template_string("test", "{{ censor value }}")
            .unwrap();

        handlebars.register_helper("censor", Box::new(censor));

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
            "######"
        );
    }

    #[test]
    fn test() {
        let resume = crate::tests::test_resume();
        assert_eq!(render(resume).unwrap(), EXPECTED);
    }

    const EXPECTED: &str = "John Smith
john.smith@gmail.com
https://github.com/josm
https://linkedin.com/in/johnsmith
==============================

PROFESSIONAL EXPERIENCE
===============================
Microsoft, Seattle, WA, Jan. 2009 - Current
Senior Software Engineer
* did a thing
* did another thing
Technologies used: C#, C++

IBM, Seattle, WA, Jan. 2007 - Dec. 2008
Software Engineer
* did a thing
* did another thing
Technologies used: Java

SAP, Seattle, WA, Summer 2005
Software Engineer Intern
* did a thing
* did another thing
Technologies used: ABAP

SKILLS
==============================
Languages: C++, Java, C#
Technologies: git, Docker

PROJECTS
==============================
Compiler
Compiles stuff
Technologies used: C#, ANTLR, LLVM

Gameboy Emulator
Emulates stuff
Technologies used: C++

EDUCATION
==============================
M.S. in Computer Science, Georgia Institute of Technology, Jan. 2007 - Current
GPA: 3.9

B.S. in Computer Science, University of Philadelphia, Jan. 2004 - Dec. 2006


";
}
