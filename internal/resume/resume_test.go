package resume_test

import (
	"testing"
	"time"

	"resumake/internal/resume"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

func TestResumeDeserialization(t *testing.T) {
	g := NewGomegaWithT(t)

	r := resume.Resume{}
	err := yaml.Unmarshal(resumeYAML, &r)
	if err != nil {
		g.Expect(err).ToNot(HaveOccurred())
	}

	expectedTime := time.Date(2004, 01, 1, 0, 0, 0, 0, time.UTC)
	expected := resume.Resume{
		CensoringEnabled: false,
		Header: resume.Header{
			Name:  "John Smith",
			Email: "john.smith@gmail.com",
		},
		EducationEntries: []resume.EducationEntry{{
			School: "University of Philadelphia",
			Degree: "B.S. in Computer Science",
			GPA:    "3.45",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.BoundedSpan{
					StartDate: expectedTime,
					EndDate:   expectedTime,
				},
			},
		}},
		JobEntries: []resume.JobEntry{{
			Title:    "Senior Software Engineer",
			Employer: "Microsoft",
			Skills:   resume.Skills{"C#", "C++"},
			Bullets:  []string{"did a thing", "did another thing"},
			Location: "Seattle, WA",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.UnboundedSpan{
					StartDate: expectedTime,
				},
			},
		}},
		Languages:    resume.Skills{"C++", "Java", "C#"},
		Technologies: resume.Skills{"git", "Docker"},
		Projects: []resume.Project{
			{
				Name:        "Compiler",
				Description: "Compiles stuff",
				Skills:      resume.Skills{"C#", "ANTLR", "LLVM"},
			},
		},
	}

	g.Expect(r.Header).To(Equal(expected.Header))
	g.Expect(r.EducationEntries).To(Equal(expected.EducationEntries))
	g.Expect(r.JobEntries).To(Equal(expected.JobEntries))
	g.Expect(r.Languages).To(Equal(expected.Languages))
	g.Expect(r.Technologies).To(Equal(expected.Technologies))
	g.Expect(r.Projects).To(Equal(expected.Projects))
}

var resumeYAML = []byte(`
header:
  name: John Smith
  email: john.smith@gmail.com
  
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
    description: Compiles stuff`)
