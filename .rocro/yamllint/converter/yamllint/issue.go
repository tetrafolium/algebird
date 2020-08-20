package yamllint

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	issueRegexp      = `^([^ ]*): (.*)$`
	occurrenceRegexp = `\[([^ ]*)\] (.*) \(([^ ]*)\)$`
)

// Issue is a reported issue by yamllint tool.
type Issue struct {
	Location   Location
	Occurrence Occurrence
}

// Location is code position part of isssue.
type Location struct {
	Filepath string
	Line     int
	Column   int
}

// Occurrence is occurrence part of issue.
type Occurrence struct {
	Level   string
	Rule    string
	Message string
}

var (
	// regular expressions of output by yamllint tool.
	regexpIssue      = regexp.MustCompile(`^([^ ]*) (.*)$`)
	regexpOccurrence = regexp.MustCompile(`^\[([^ ]*)\] (.*) \(([^ ]*)\)$`)
)

// Parse parses outout line by yamllint tool.
func Parse(str string) (*Issue, error) {
	issue := regexpIssue.FindStringSubmatch(str)

	var (
		//_, strLoc, strOcc = issue...
		strLoc = issue[1]
		strOcc = issue[2]

		loc       = strings.Split(strLoc, `:`)
		filepath  = loc[0]
		line, _   = strconv.Atoi(loc[1])
		column, _ = strconv.Atoi(loc[2])

		//_, level, message, rule = regexpOccurrence.FindStringSubmatch(strOcc)...
		occ     = regexpOccurrence.FindStringSubmatch(strOcc)
		level   = occ[1]
		message = occ[2]
		rule    = occ[3]
	)

	return &Issue{
		Location{
			Filepath: filepath,
			Line:     line,
			Column:   column,
		},
		Occurrence{
			Level:   level,
			Rule:    rule,
			Message: message,
		},
	}, nil
}
