package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/convert"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/sarif"
	"github.com/tetrafolium/algebird/.rocro/yamllint/converter/yamllint"
)

const (
	jsonPrefix = ``
	jsonIndent = `  `
)

func main() {
	var results []*sarif.Result

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

		results = append(results, result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	byteslice, err := json.MarshalIndent(results, jsonPrefix, jsonIndent)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json.Marshal failed:", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, bytes.NewBuffer(byteslice).String())
}
