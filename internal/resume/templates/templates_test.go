package templates_test

import (
	"time"

	"github.com/karimElmougi/resumake/internal/resume"
)

var testStartTime = time.Date(2004, 01, 1, 0, 0, 0, 0, time.UTC)

var testEndTime = time.Date(2006, 12, 31, 0, 0, 0, 0, time.UTC)

var testResume = resume.Resume{
	Header: resume.Header{
		Name:             "John Smith",
		Email:            "john.smith@gmail.com",
		GitHubUsername:   "josm",
		LinkedInUsername: "johnsmith",
	},
	EducationEntries: []resume.EducationEntry{
		{
			School: "Georgia Institute of Technology",
			Degree: "M.S. in Computer Science",
			GPA:    "3.9",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.UnboundedSpan{
					StartDate: testStartTime,
				},
			},
		},
		{
			School: "University of Philadelphia",
			Degree: "B.S. in Computer Science",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.BoundedSpan{
					StartDate: testStartTime,
					EndDate:   testEndTime,
				},
			},
		},
	},
	JobEntries: []resume.JobEntry{
		{
			Title:    "Senior Software Engineer",
			Employer: "Microsoft",
			Skills:   resume.Skills{"C#", "C++"},
			Bullets:  []string{"did a thing", "did another thing"},
			Location: "Seattle, WA",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.UnboundedSpan{
					StartDate: testStartTime,
				},
			},
		},
		{
			Title:    "Software Engineer",
			Employer: "IBM",
			Skills:   resume.Skills{"Java"},
			Bullets:  []string{"did a thing", "did another thing"},
			Location: "Seattle, WA",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.BoundedSpan{
					StartDate: testStartTime,
					EndDate:   testEndTime,
				},
			},
		},
		{
			Title:    "Software Engineer Intern",
			Employer: "SAP",
			Skills:   resume.Skills{"ABAP"},
			Bullets:  []string{"did a thing", "did another thing"},
			Location: "Seattle, WA",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.SeasonSpan{
					Season: "Winter",
					Year:   2004,
				},
			},
		},
	},
	Languages:    resume.Skills{"C++", "Java", "C#"},
	Technologies: resume.Skills{"git", "Docker"},
	Projects: []resume.Project{
		{
			Name:        "Compiler",
			Description: "Compiles stuff",
			Skills:      resume.Skills{"C#", "ANTLR", "LLVM"},
		},
		{
			Name:        "Gameboy Emulator",
			Description: "Emulates stuff",
			Skills:      resume.Skills{"C++"},
		},
	},
}
