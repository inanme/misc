package morestrings

import (
	"strings"
)

func Join(elems []string, sep, prefix, suffix string) string {
	joined := strings.Join(elems, sep)
	if joined == "" {
		return joined
	}
	return prefix + joined + suffix
}

func ContainsAll(str, sub string, rest ...string) bool {
	if !strings.Contains(str, sub) {
		return false
	}
	for _, sub := range rest {
		if !strings.Contains(str, sub) {
			return false
		}
	}
	return true
}

func ContainsAny(str, sub string, rest ...string) bool {
	if strings.Contains(str, sub) {
		return true
	}
	for _, sub := range rest {
		if strings.Contains(str, sub) {
			return true
		}
	}
	return false
}

func ContainAll(strs []string, sub string, rest ...string) bool {
	for _, str := range strs {
		if !strings.Contains(str, sub) {
			return false
		}
		for _, sub := range rest {
			if !strings.Contains(str, sub) {
				return false
			}
		}
	}
	return true
}

func ContainAny(strs []string, sub string, rest ...string) bool {
	for _, str := range strs {
		if strings.Contains(str, sub) {
			return true
		}
		for _, sub := range rest {
			if strings.Contains(str, sub) {
				return true
			}
		}
	}
	return false
}
