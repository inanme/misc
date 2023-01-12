package morestrings

import "strings"

func Join(elems []string, sep, prefix, suffix string) string {
	joined := strings.Join(elems, sep)
	if joined == "" {
		return joined
	}
	return prefix + joined + suffix
}
