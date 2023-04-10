package tailwind

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertHeaders(input string, headerRegex *regexp.Regexp) string {
	if matches := headerRegex.FindString(input); len(matches) > 0 {
		headerLevel := len(matches)
		input = headerRegex.ReplaceAllString(input, "")
		return fmt.Sprintf("<h%d class=\"text-4xl font-bold\">%s</h%d>", headerLevel, strings.TrimSpace(input), headerLevel)
	}
	return input
}

func ConvertBold(input string, boldRegex *regexp.Regexp) string {
	return boldRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<strong class=\"font-semibold\">" + boldRegex.ReplaceAllString(str, "$1") + "</strong>"
	})
}

func ConvertItalic(input string, italicRegex *regexp.Regexp) string {
	return italicRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<em class=\"italic\">" + italicRegex.ReplaceAllString(str, "$1") + "</em>"
	})
}

func ConvertLinks(input string, linkRegex *regexp.Regexp) string {
	return linkRegex.ReplaceAllStringFunc(input, func(str string) string {
		return "<a class=\"text-blue-500 hover:text-blue-700\" href=\"" + linkRegex.ReplaceAllString(str, "$2") + "\">" + linkRegex.ReplaceAllString(str, "$1") + "</a>"
	})
}
