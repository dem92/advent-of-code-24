package utils

import "regexp"

func FindRegexMatches(regex, line string) []string {
	r := regexp.MustCompile(regex)
	matches := r.FindAllString(line, -1)
	return matches
}
