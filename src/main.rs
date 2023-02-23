use resumake::Resume;

use std::error::Error;
use std::fs::File;
use std::io::Write;
use std::path::{Path, PathBuf};

use clap::{Parser, Subcommand};
use resumake::templates;

#[derive(Parser, Debug)]
#[command(author, version, about)]
struct Args {
    #[command(subcommand)]
    command: Command,

    #[arg(short, long)]
    censor: bool,

    #[arg(short, long)]
    output: Option<PathBuf>,
}

#[derive(Subcommand, Debug)]
enum Command {
    Latex {
        resume: PathBuf,
        #[arg(short, long)]
        pdf: bool,
    },
    Plaintext {
        resume: PathBuf,
    },
}

fn main() {
    let args = Args::parse();

    let resume_path = match &args.command {
        Command::Latex { resume, .. } => resume,
        Command::Plaintext { resume } => resume,
    };
    let resume = match read_resume(&resume_path) {
        Ok(resume) => resume,
        Err(err) => {
            eprintln!("Invalid resume file: {err}");
            std::process::exit(1);
        }
    };

    let result = match args.command {
        Command::Latex { .. } => templates::latex::render(resume),
        Command::Plaintext { .. } => templates::plaintext::render(resume),
    };

    let output = match result {
        Ok(output) => output,
        Err(err) => {
            eprintln!("Unabe to render template: {err}");
            std::process::exit(1);
        }
    };

    match args.output {
        Some(path) => File::create(path)
            .and_then(|mut f| f.write_all(output.as_bytes()))
            .unwrap(),
        None => println!("{}", output),
    }
}

fn read_resume(path: &Path) -> Result<Resume, Box<dyn Error>> {
    serde_yaml::from_reader(File::open(path)?).map_err(Into::into)
}
