package morestrings

import (
	"fmt"
	"strings"
)

const (
	DefaultCapacity = 1024
)

func JoinT[T fmt.Stringer](elems []T, sep, prefix, suffix string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return prefix + elems[0].String() + suffix
	}

	var b strings.Builder
	b.Grow(DefaultCapacity)
	b.WriteString(prefix)
	b.WriteString(elems[0].String())
	for _, e := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(e.String())
	}
	b.WriteString(suffix)
	return b.String()
}

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

func ContainsAllCaseInsensitive(str, sub string, rest ...string) bool {
	str = strings.ToLower(str)
	sub = strings.ToLower(sub)
	if !strings.Contains(str, sub) {
		return false
	}
	for _, sub := range rest {
		if !strings.Contains(str, strings.ToLower(sub)) {
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

func ContainsAnyCaseInsensitive(str, sub string, rest ...string) bool {
	str = strings.ToLower(str)
	sub = strings.ToLower(sub)
	if strings.Contains(str, sub) {
		return true
	}
	for _, sub := range rest {
		if strings.Contains(str, strings.ToLower(sub)) {
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

func ContainAllCaseInsensitive(strs []string, sub string, rest ...string) bool {
	sub = strings.ToLower(sub)
	for _, str := range strs {
		str = strings.ToLower(str)
		if !strings.Contains(str, sub) {
			return false
		}
		for _, sub := range rest {
			if !strings.Contains(str, strings.ToLower(sub)) {
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

func ContainAnyCaseInsensitive(strs []string, sub string, rest ...string) bool {
	sub = strings.ToLower(sub)
	for _, str := range strs {
		str = strings.ToLower(str)
		if strings.Contains(str, sub) {
			return true
		}
		for _, sub := range rest {
			if strings.Contains(str, strings.ToLower(sub)) {
				return true
			}
		}
	}
	return false
}
