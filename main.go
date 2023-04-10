package main

import (
	"fmt"
	"regexp"
	"strings"
)

func MarkdownToHTML(input string) string {
	lines := strings.Split(input, "\n")

	headerRegex := regexp.MustCompile(`^#{1,6}`)
	boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
	italicRegex := regexp.MustCompile(`\*(.*?)\*`)
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	var output strings.Builder
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Headers
		trimmed = convertHeaders(trimmed, headerRegex)

		// Bold
		trimmed = convertBold(trimmed, boldRegex)

		// Italic
		trimmed = convertItalic(trimmed, italicRegex)

		// Links
		trimmed = convertLinks(trimmed, linkRegex)

		output.WriteString(trimmed)
		output.WriteString("\n")
	}

	return output.String()
}

func convertHeaders(input string, headerRegex *regexp.Regexp) string {
	if matches := headerRegex.FindString(input); len(matches) > 0 {
		headerLevel := len(matches)
		input = headerRegex.ReplaceAllString(input, "")
		return fmt.Sprintf("<h%d>%s</h%d>", headerLevel, strings.TrimSpace(input), headerLevel)
	}
	return input
}

func convertBold(input string, boldRegex *regexp.Regexp) string {
	return boldRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<strong>" + boldRegex.ReplaceAllString(str, "$1") + "</strong>"
	})
}

func convertItalic(input string, italicRegex *regexp.Regexp) string {
	return italicRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<em>" + italicRegex.ReplaceAllString(str, "$1") + "</em>"
	})
}

func convertLinks(input string, linkRegex *regexp.Regexp) string {
	return linkRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<a href=\"" + linkRegex.ReplaceAllString(str, "$2") + "\">" + linkRegex.ReplaceAllString(str, "$1") + "</a>"
	})
}

func main() {
	md := `# Header 1
**Bold text**
*Italic text*
[Link to Google](https://www.google.com)`

	fmt.Println(MarkdownToHTML(md))
}
