package templates_test

import (
	"time"

	"resume/internal/resume"
)

var sampleStartTime = time.Date(2004, 01, 1, 0, 0, 0, 0, time.UTC)

var sampleEndTime = time.Date(2006, 12, 31, 0, 0, 0, 0, time.UTC)

var sampleResume = resume.Resume{
	Header: resume.Header{
		Name:  "John Smith",
		Email: "john.smith@gmail.com",
	},
	EducationEntries: []resume.EducationEntry{
		{
			School: "Georgia Institute of Technology",
			Degree: "M.S. in Computer Science",
			GPA:    "3.9",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.UnboundedSpan{
					StartDate: sampleStartTime,
				},
			},
		},
		{
			School: "University of Philadelphia",
			Degree: "B.S. in Computer Science",
			TimeSpan: resume.TimeSpan{
				TimeSpanVariant: resume.BoundedSpan{
					StartDate: sampleStartTime,
					EndDate:   sampleEndTime,
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
					StartDate: sampleStartTime,
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
					StartDate: sampleStartTime,
					EndDate:   sampleEndTime,
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
