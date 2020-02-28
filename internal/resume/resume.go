package resume

import (
	"fmt"
	"strings"
	"time"
)

type Resume struct {
	Header           Header           `yaml:"header"`
	EducationEntries []EducationEntry `yaml:"education"`
	JobEntries       []JobEntry       `yaml:"experience"`
	Languages        Skills           `yaml:"languages"`
	Technologies     Skills           `yaml:"technologies"`
	Projects         []Project        `yaml:"projects"`
}

type Header struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type EducationEntry struct {
	School   string   `yaml:"school"`
	Degree   string   `yaml:"degree"`
	GPA      string   `yaml:"gpa"`
	TimeSpan TimeSpan `yaml:"timespan"`
}

type JobEntry struct {
	Title    string   `yaml:"title"`
	Employer string   `yaml:"employer"`
	Skills   Skills   `yaml:"skills"`
	Bullets  []string `yaml:"bullets"`
	Location string   `yaml:"location"`
	TimeSpan TimeSpan `yaml:"timespan"`
}

type Project struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Skills      Skills `yaml:"skills"`
}

type Skills []string

func (s Skills) Display() string {
	return strings.Join(s, ", ")
}

type TimeSpan struct {
	TimeSpanVariant
}

func (t *TimeSpan) UnmarshalYAML(unmarshal func(interface{}) error) error {
	data := make(map[string]interface{})
	if err := unmarshal(&data); err != nil {
		return err
	}

	startDate, hasStart := data["start"].(string)
	endDate, hasEnd := data["end"].(string)

	if hasStart && hasEnd {
		start, err := time.Parse("01/2006", startDate)
		if err != nil {
			return err
		}

		end, err := time.Parse("01/2006", endDate)
		if err != nil {
			return err
		}

		t.TimeSpanVariant = BoundedSpan{
			StartDate: start,
			EndDate:   end,
		}
	} else if hasStart {
		start, err := time.Parse("01/2006", startDate)
		if err != nil {
			return err
		}

		t.TimeSpanVariant = UnboundedSpan{StartDate: start}
	} else {
		season := data["season"].(string)
		year := data["year"].(int)

		t.TimeSpanVariant = SeasonSpan{Season: season, Year: year}
	}

	return nil
}

type TimeSpanVariant interface {
	Display() string
}

type SeasonSpan struct {
	Season string
	Year   int
}

func (s SeasonSpan) Display() string {
	return fmt.Sprintf("%v %v", s.Season, s.Year)
}

type UnboundedSpan struct {
	StartDate time.Time
}

func (u UnboundedSpan) Display() string {
	startYear, startMonth, _ := u.StartDate.Date()
	return fmt.Sprintf("%v - Current", formatMonthYear(startMonth, startYear))
}

type BoundedSpan struct {
	StartDate time.Time
	EndDate   time.Time
}

func (b BoundedSpan) Display() string {
	return fmt.Sprintf("%v - %v", b.displayStartDate(), b.displayEndDate())
}

func (b BoundedSpan) displayStartDate() string {
	startYear, startMonth, _ := b.StartDate.Date()
	return formatMonthYear(startMonth, startYear)
}

func (b BoundedSpan) displayEndDate() string {
	endYear, endMonth, _ := b.EndDate.Date()
	return formatMonthYear(endMonth, endYear)
}

func formatMonthYear(month time.Month, year int) string {
	return fmt.Sprintf("%v %v", trimMonth(month), year)
}

func trimMonth(month time.Month) string {
	monthStr := month.String()
	if len(monthStr) > 3 {
		return monthStr[0:3] + "."
	}
	return monthStr
}
