package main

import (
	"regexp"
	"strings"

	"golang.design/x/clipboard"
)

var regularExpression = regexp.MustCompile("((https?://(www|smile)?.?)?amazon.de/)[^/]*/?((dp|gp/product)/[^/^?]*)")

func ReduceAmazonUrls(input string) (string, error) {
	results := []string{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		matches := regularExpression.FindStringSubmatch(line)

		if matches == nil {
			// fmt.Printf("No match found\n")
			results = append(results, line)
			continue
		}

		// fmt.Printf("Found matches: `%q`\n", matches)
		if matches == nil || len(matches) < 4 {
			continue
		}

		reducedUrl := strings.Join([]string{matches[1], matches[4]}, "")
		results = append(results, reducedUrl)
		// fmt.Printf("reducedUrl: `%q`\n", reducedUrl)
	}

	if results == nil {
		return "", nil
	}

	return strings.Join(results, "\n"), nil
}

func ReadClipboard() string {
	content := clipboard.Read(clipboard.FmtText)
	contentString := string(content[:])
	return contentString
}

func WriteClipboard(output string) {
	clipboard.Write(clipboard.FmtText, []byte(output))
}

func main() {
	clipboardContent := ReadClipboard()
	if clipboardContent == "" {
		return
	}
	shortenedResultString, err := ReduceAmazonUrls(clipboardContent)
	if err == nil && shortenedResultString != "" {
		WriteClipboard(shortenedResultString)
	}
}
