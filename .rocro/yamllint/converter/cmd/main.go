package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tetrafolium/algebird/.rocro/sarif"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/convert"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/yamllint"
)

const (
	jsonPrefix = ``
	jsonIndent = `  `
)

func main() {
	var (
		results   []sarif.Result
		artifacts []sarif.Artifact
	)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		issue, err := yamllint.Parse(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot parse standard input:", err)
			os.Exit(1)
		}

		result, err := convert.IssueToResult(issue)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot convert sarif result:", err)
			os.Exit(1)
		}

		results = append(results, *result)
	}

	artifacts, err := convert.ResultsToArtifacts(results)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in resultsToArtifacts:", err)
		os.Exit(1)
	}

	run := sarif.Run{
		OriginalURIBaseIDs: convert.OriginalURIBaseIDs(),
		Artifacts:          artifacts,
		Results:            &results,
	}

	sarifLog := sarif.Log{
		Version: sarif.VERSION,
		Schema:  sarif.SCHEMA,
		Runs:    []sarif.Run{run},
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	// jsonBytes, err := json.Marshal(results)
	jsonBytes, err := json.MarshalIndent(sarifLog, jsonPrefix, jsonIndent)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json.Marshal failed:", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, bytes.NewBuffer(jsonBytes).String())
}
