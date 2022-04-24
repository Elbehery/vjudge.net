package main


import "fmt"

func KMP(pat string) []int {
	if pat == "" {
		return nil
	}

	prefixes := make([]int, len(pat))
	prefixes[0] = 0
	for i := 1; i < len(pat); i++ {
		l := prefixes[i-1]
		ch := pat[i]
		for l != 0 && ch != pat[l] {
			l = prefixes[l-1]
		}
		if ch == pat[l] {
			l += 1
		}
		prefixes[i] = l
	}
	return prefixes
}

func main() {
	res := KMP("aabaab")
	if res != nil {
		fmt.Println(res)
	}
}
