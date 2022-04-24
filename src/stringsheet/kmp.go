package stringsheet

import (
	"errors"
)

var (
	EmptyPatternErr = errors.New("input pattern is empty")
)

func ComputeKMPLPS(pattern string) ([]int, error) {
	if len(pattern) <= 0 {
		return nil, EmptyPatternErr
	}
	lps := make([]int, len(pattern))
	lps[0] = 0
	for i := 1; i < len(pattern); i++ {
		ch := pattern[i]
		l := lps[i-1]
		for l != 0 && pattern[l] != ch {
			l = lps[l-1]
		}
		if ch == pattern[l] {
			l += 1
		}
		lps[i] = l
	}
	return lps, nil
}
