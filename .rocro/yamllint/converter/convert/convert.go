package convert

import (
	"path/filepath"
	"reflect"

	"github.com/tetrafolium/algebird/.rocro/sarif"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/yamllint"
)

func IssueToResult(issue *yamllint.Issue) (*sarif.Result, error) {
	locations := issueLocationToResultLocations(&issue.Location)
	result := sarif.Result{
		Level: issue.Occurrence.Level,
		Message: &sarif.Message{
			Text: issue.Occurrence.Message,
		},
		Locations: locations,
		Rank:      levelToRank(issue.Occurrence.Level),
	}
	return &result, nil
}

func issueLocationToResultLocations(issueLoc *yamllint.Location) []sarif.Location {
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
		ArtifactLocation: &artifactLocation,
		Region:           &region,
	}
	location := sarif.Location{
		PhysicalLocation: &physicalLocation,
	}
	locations := []sarif.Location{location}
	return locations
}

func ResultsToArtifacts(results []sarif.Result) ([]sarif.Artifact, error) {
	var locations []sarif.Location
	var artifacts []sarif.Artifact

	for _, result := range results {
		for _, location := range result.Locations {
			if locationInList(location, locations) {
				continue
			}

			artifact := sarif.Artifact{
				Location: &location,
			}
			locations = append(locations, location)
			artifacts = append(artifacts, artifact)
		}
	}

	return artifacts, nil
}

func locationInList(loc sarif.Location, list []sarif.Location) bool {
	for _, elem := range list {
		if equalLocation(loc, elem) {
			return true
		}
	}
	return false
}

func equalLocation(a, b sarif.Location) bool {
	return reflect.DeepEqual(a, b)
}

var (
	mapLevelToRank = map[string]float64{
		"info":     10.0,
		"warning":  50.0,
		"error":    90.0,
		"critical": 100.0,
	}

	// NOTE: tractrix/codelift should have this table.
	mapFileExtentionToLanguage = map[string]string{
		".go":   "Go",
		".json": "JSON",
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
