package main

import (
	"regexp"
	"strings"

	"golang.design/x/clipboard"
)

var amazonRegex = regexp.MustCompile(`((https?://(www|smile)?.?)?amazon.de/).*/?((dp|gp/product)/[^/^?]*)`)
var etsyRegex = regexp.MustCompile(`(https?://www\.etsy\.com/[^/]+/listing/[0-9]+)`)

func ReduceUrls(input string) (string, error) {
	results := []string{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Skip empty lines
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}

		// Try Amazon regex first
		amazonMatches := amazonRegex.FindStringSubmatch(line)
		if len(amazonMatches) >= 5 {
			reducedUrl := strings.Join([]string{amazonMatches[1], amazonMatches[4]}, "")
			results = append(results, reducedUrl)
			continue
		}

		// Try Etsy regex
		etsyMatches := etsyRegex.FindStringSubmatch(line)
		if len(etsyMatches) >= 1 {
			reducedUrl := etsyMatches[1]
			results = append(results, reducedUrl)
			continue
		}

		// No match found, keep original line
		results = append(results, line)
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
	shortenedResultString, err := ReduceUrls(clipboardContent)
	if err == nil && shortenedResultString != "" {
		WriteClipboard(shortenedResultString)
	}
}
