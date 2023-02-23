use chrono::NaiveDate;
use serde::Deserialize;

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub struct Resume {
    pub header: Header,
    pub education: Vec<EducationEntry>,
    pub experience: Vec<Job>,
    pub languages: Vec<String>,
    pub technologies: Vec<String>,
    pub projects: Vec<Project>,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub struct Header {
    pub name: String,
    pub email: String,
    pub github: String,
    pub linkedin: Option<String>,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub struct EducationEntry {
    pub school: String,
    pub degree: String,
    pub gpa: String,
    pub timespan: Timespan,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub struct Job {
    pub title: String,
    pub employer: String,
    pub location: String,
    pub skills: Vec<String>,
    pub bullets: Vec<String>,
    pub timespan: Timespan,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub struct Project {
    pub name: String,
    pub description: String,
    pub skills: Vec<String>,
}

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
#[serde(untagged)]
pub enum Timespan {
    Season(Season, Year),

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

type Year = u32;

#[derive(Debug, Clone, PartialEq, Eq, Deserialize)]
pub enum Season {
    Winter,
    Spring,
    Summer,
    Fall,
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

#[cfg(test)]
mod tests {
    use super::*;
    use pretty_assertions::assert_eq;

    #[test]
    fn test() {
        let expected = Resume {
            header: Header {
                name: "John Smith".to_string(),
                email: "john.smith@gmail.com".to_string(),
                github: "josm".to_string(),
                linkedin: Some("johnsmith".to_string()),
            },
            education: vec![EducationEntry {
                school: "University of Philadelphia".to_string(),
                degree: "B.S. in Computer Science".to_string(),
                gpa: "3.45".to_string(),
                timespan: Timespan::Bounded {
                    start: NaiveDate::from_ymd_opt(2004, 01, 01).unwrap(),
                    end: NaiveDate::from_ymd_opt(2004, 01, 01).unwrap(),
                },
            }],
            experience: vec![Job {
                title: "Senior Software Engineer".to_string(),
                employer: "Microsoft".to_string(),
                location: "Seattle, WA".to_string(),
                timespan: Timespan::Unbounded {
                    start: NaiveDate::from_ymd_opt(2004, 01, 01).unwrap(),
                },
                skills: vec!["C#".to_string(), "C++".to_string()],
                bullets: vec!["did a thing".to_string(), "did another thing".to_string()],
            }],
            languages: vec!["C++".to_string(), "Java".to_string(), "C#".to_string()],
            technologies: vec!["git".to_string(), "Docker".to_string()],
            projects: vec![Project {
                name: "Compiler".to_string(),
                skills: vec!["C#".to_string(), "ANTLR".to_string(), "LLVM".to_string()],
                description: "Compiles stuff".to_string(),
            }],
        };

        let result: Resume = serde_yaml::from_str(INPUT).unwrap();

        assert_eq!(result, expected);
    }

    const INPUT: &str = "
header:
  name: John Smith
  email: john.smith@gmail.com
  github: josm
  linkedin: johnsmith
  
education:
  - school: University of Philadelphia
    degree: B.S. in Computer Science
    gpa: 3.45
    timespan:
      start: 01/2004
      end: 01/2004
      
experience:
  - title: Senior Software Engineer
    employer: Microsoft
    location: Seattle, WA
    timespan: 
      start: 01/2004
    skills:
      - C#
      - C++
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
    description: Compiles stuff";
}
