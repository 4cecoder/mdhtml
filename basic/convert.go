package basic

import (
	"fmt"
	"regexp"
	"strings"
)

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
