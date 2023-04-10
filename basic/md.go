package basic

import (
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
