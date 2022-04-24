package stringsheet

import (
	"errors"
)

var (
	EmptyPatternErr = errors.New("input pattern is empty")
	EmptyTextErr    = errors.New("input text is empty")
)

// KMPMatch computes the number of occurrences of the given pattern within the given text
// returns []int which contains the indices of the pattern occurrences
func KMPMatch(pattern, txt string) ([]int, error) {
	if len(txt) <= 0 {
		return nil, EmptyTextErr
	}
	lps, err := ComputeKMPLPS(pattern)
	if err != nil {
		return nil, err
	}
	occurrences := make([]int, 0)
	i, j := 0, 0
	for i < len(txt) {
		if pattern[j] == txt[i] {
			i++
			j++
		}
		if j >= len(pattern) {
			// match found
			occurrences = append(occurrences, i-j)
			j = lps[j-1]
		} else if pattern[j] != txt[i] {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	return occurrences, nil
}

// ComputeKMPLPS computes the longest prefix that is also suffix in the input pattern
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
