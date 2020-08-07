package convert

import (
	"path/filepath"

	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/sarif"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/yamllint"
)

func IssueToResult(issue *yamllint.Issue) (*sarif.Result, error) {
	locations := issueLocationToResultLocations(&issue.Location)
	result := sarif.Result{
		Level: issue.Occurrence.Level,
		Message: sarif.Message{
			Text: issue.Occurrence.Message,
		},
		Locations: locations,
		Rank:      levelToRank(issue.Occurrence.Level),
	}
	return &result, nil
}

func issueLocationToResultLocations(issueLoc *yamllint.Location) sarif.Locations {
	fileExt := filepath.Ext(issueLoc.Filepath)
	artifactLocation := sarif.ArtifactLocation{
		URI:       issueLoc.Filepath,
		URIBaseID: `REPOROOT`,
	}
	region := sarif.Region{
		StartLine:      issueLoc.Line,
		StartColumn:    issueLoc.Column,
		SourceLanguage: fileExtToLanguage(fileExt),
	}
	physicalLocation := sarif.PhysicalLocation{
		ArtifactLocation: artifactLocation,
		Region:           region,
	}
	location := sarif.Location{
		PhysicalLocation: physicalLocation,
	}
	locations := sarif.Locations{&location}
	return locations
}

var (
	mapLevelToRank = map[string]float64{
		"info":     10.0,
		"warning":  50.0,
		"error":    90.0,
		"critical": 100.0,
	}
	mapFileExtentionToLanguage = map[string]string{
		".yml":  "YAML",
		".yaml": "YAML",
	}
)

func levelToRank(level string) float64 {
	rank, ok := mapLevelToRank[level]
	if ok {
		return rank
	}
	return 0.0
}

func fileExtToLanguage(ext string) string {
	lang, ok := mapFileExtentionToLanguage[ext]
	if ok {
		return lang
	}
	return ""
}
