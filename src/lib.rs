pub mod templates;

use chrono::{Datelike, NaiveDate};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Resume {
    pub header: Header,
    pub education: Vec<EducationEntry>,
    pub experience: Vec<Job>,
    pub languages: Vec<String>,
    pub technologies: Vec<String>,
    pub projects: Vec<Project>,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Header {
    pub name: String,
    pub email: String,
    pub github: String,
    pub linkedin: Option<String>,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct EducationEntry {
    pub school: String,
    pub degree: String,

    #[serde(skip_serializing_if = "Option::is_none")]
    pub gpa: Option<String>,

    pub timespan: Timespan,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Job {
    pub title: String,
    pub employer: String,
    pub location: String,
    pub skills: Vec<String>,
    pub bullets: Vec<String>,
    pub timespan: Timespan,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Project {
    pub name: String,
    pub description: String,
    pub skills: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
#[serde(untagged)]
pub enum Timespan {
    Season {
        season: Season,
        year: u32,
    },

    Bounded {
        #[serde(deserialize_with = "deserialize_date")]
        start: NaiveDate,

        #[serde(deserialize_with = "deserialize_date")]
        end: NaiveDate,
    },

    Unbounded {
        #[serde(deserialize_with = "deserialize_date")]
        start: NaiveDate,
    },
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub enum Season {
    Winter,
    Spring,
    Summer,
    Fall,
}

impl std::fmt::Display for Season {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        std::fmt::Debug::fmt(self, f)
    }
}

fn deserialize_date<'de, D>(deserializer: D) -> Result<NaiveDate, D::Error>
where
    D: serde::Deserializer<'de>,
{
    const FMT: &str = "%d/%m/%Y";

    let date = String::deserialize(deserializer)?;
    NaiveDate::parse_from_str(&date, FMT)
        .or_else(|_| {
            let date = format!("01/{date}");
            NaiveDate::parse_from_str(&date, FMT)
        })
        .map_err(|_| {
            serde::de::Error::custom(format!(
                "Expected date format dd/mm/yyyy or mm/yyyy, got `{date}`"
            ))
        })
}

impl Serialize for Timespan {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let output = match self {
            Timespan::Season { season, year } => format!("{season} {year}"),
            Timespan::Bounded { start, end } => {
                format!("{} - {}", format(start), format(end))
            }
            Timespan::Unbounded { start } => format!("{} - Current", format(start)),
        };

        serializer.serialize_str(&output)
    }
}

fn format(date: &NaiveDate) -> String {
    if date.month() == 5 {
        date.format("%b %Y").to_string()
    } else {
        date.format("%b. %Y").to_string()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use pretty_assertions::assert_eq;

    #[test]
    fn season_serialize_test() {
        let span = Timespan::Season {
            season: Season::Winter,
            year: 2001,
        };
        let s = serde_yaml::to_string(&span).unwrap();

        assert_eq!(s, "Winter 2001\n")
    }

    #[test]
    fn bounded_serialize_test() {
        let span = Timespan::Bounded {
            start: NaiveDate::from_ymd_opt(2001, 04, 01).unwrap(),
            end: NaiveDate::from_ymd_opt(2001, 12, 01).unwrap(),
        };

        let s = serde_yaml::to_string(&span).unwrap();

        assert_eq!(s, "Apr. 2001 - Dec. 2001\n")
    }

    #[test]
    fn unbounded_serialize_test() {
        let span = Timespan::Unbounded {
            start: NaiveDate::from_ymd_opt(2001, 05, 01).unwrap(),
        };

        let s = serde_yaml::to_string(&span).unwrap();

        assert_eq!(s, "May 2001 - Current\n")
    }

    #[test]
    fn test() {
        let result: Resume = serde_yaml::from_str(INPUT).unwrap();

        assert_eq!(result, test_resume());
    }

    pub fn test_resume() -> Resume {
        Resume {
            header: Header {
                name: "John Smith".to_string(),
                email: "john.smith@gmail.com".to_string(),
                github: "josm".to_string(),
                linkedin: Some("johnsmith".to_string()),
            },
            education: vec![
                EducationEntry {
                    school: "Georgia Institute of Technology".to_string(),
                    degree: "M.S. in Computer Science".to_string(),
                    gpa: Some("3.9".to_string()),
                    timespan: Timespan::Unbounded {
                        start: NaiveDate::from_ymd_opt(2007, 01, 01).unwrap(),
                    },
                },
                EducationEntry {
                    school: "University of Philadelphia".to_string(),
                    degree: "B.S. in Computer Science".to_string(),
                    gpa: None,
                    timespan: Timespan::Bounded {
                        start: NaiveDate::from_ymd_opt(2004, 01, 01).unwrap(),
                        end: NaiveDate::from_ymd_opt(2006, 12, 01).unwrap(),
                    },
                },
            ],
            experience: vec![
                Job {
                    title: "Senior Software Engineer".to_string(),
                    employer: "Microsoft".to_string(),
                    location: "Seattle, WA".to_string(),
                    timespan: Timespan::Unbounded {
                        start: NaiveDate::from_ymd_opt(2009, 01, 01).unwrap(),
                    },
                    skills: vec!["C#".to_string(), "C++".to_string()],
                    bullets: vec!["did a thing".to_string(), "did another thing".to_string()],
                },
                Job {
                    title: "Software Engineer".to_string(),
                    employer: "IBM".to_string(),
                    location: "Seattle, WA".to_string(),
                    timespan: Timespan::Bounded {
                        start: NaiveDate::from_ymd_opt(2007, 01, 01).unwrap(),
                        end: NaiveDate::from_ymd_opt(2008, 12, 01).unwrap(),
                    },
                    skills: vec!["Java".to_string()],
                    bullets: vec!["did a thing".to_string(), "did another thing".to_string()],
                },
                Job {
                    title: "Software Engineer Intern".to_string(),
                    employer: "SAP".to_string(),
                    location: "Seattle, WA".to_string(),
                    timespan: Timespan::Season {
                        year: 2005,
                        season: Season::Summer,
                    },
                    skills: vec!["ABAP".to_string()],
                    bullets: vec!["did a thing".to_string(), "did another thing".to_string()],
                },
            ],
            languages: vec!["C++".to_string(), "Java".to_string(), "C#".to_string()],
            technologies: vec!["git".to_string(), "Docker".to_string()],
            projects: vec![
                Project {
                    name: "Compiler".to_string(),
                    skills: vec!["C#".to_string(), "ANTLR".to_string(), "LLVM".to_string()],
                    description: "Compiles stuff".to_string(),
                },
                Project {
                    name: "Gameboy Emulator".to_string(),
                    skills: vec!["C++".to_string()],
                    description: "Emulates stuff".to_string(),
                },
            ],
        }
    }

    const INPUT: &str = "
header:
  name: John Smith
  email: john.smith@gmail.com
  github: josm
  linkedin: johnsmith
  
education:
  - school: Georgia Institute of Technology
    degree: M.S. in Computer Science
    gpa: 3.9
    timespan:
      start: 01/2007
  - school: University of Philadelphia
    degree: B.S. in Computer Science
    timespan:
      start: 01/2004
      end: 12/2006
      
experience:
  - title: Senior Software Engineer
    employer: Microsoft
    location: Seattle, WA
    timespan: 
      start: 01/2009
    skills:
      - C#
      - C++
    bullets:
      - did a thing
      - did another thing
  - title: Software Engineer
    employer: IBM
    location: Seattle, WA
    timespan: 
      start: 01/2007
      end: 12/2008
    skills:
      - Java
    bullets:
      - did a thing
      - did another thing
  - title: Software Engineer Intern
    employer: SAP
    location: Seattle, WA
    timespan: 
        season: Summer
        year: 2005
    skills:
      - ABAP
    bullets:
      - did a thing
      - did another thing

languages:
  - C++
  - Java
  - C#

technologies:
  - git
  - Docker

projects:
  - name: Compiler
    skills: 
      - C#
      - ANTLR
      - LLVM
    description: Compiles stuff
  - name: Gameboy Emulator
    description: Emulates stuff
    skills:
      - C++";
}
