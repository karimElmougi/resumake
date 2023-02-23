pub mod latex;
pub mod plaintext;

use crate::Resume;

use handlebars::{Handlebars, JsonRender};
use regex::Regex;
use serde::Serialize;

#[derive(Serialize)]
struct Context {
    is_censored: bool,
    has_skill_section: bool,

    #[serde(flatten)]
    resume: Resume,
}

fn list(
    h: &handlebars::Helper,
    handlebars: &Handlebars,
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

    let escape_fn = handlebars.get_escape_fn();

    let output = param
        .iter()
        .map(JsonRender::render)
        .map(|s| escape_fn(&s))
        .collect::<Vec<_>>()
        .join(", ");

    write!(out, "{output}")?;

    Ok(())
}

fn censor(param: &str, c: &handlebars::Context, replacement: &str) -> String {
    let is_censored = c.data().get("is_censored").unwrap().as_bool().unwrap();
    if is_censored {
        let re = Regex::new(r"\|\|(.*?)\|\|").unwrap();
        re.replace_all(&param, replacement).to_string()
    } else {
        param.replace("||", "")
    }
}
