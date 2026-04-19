func lengthOfLongestSubstring(s string) int {
	// sliding window
	// make a set storing all characters in the window
	// if duplicate, move l++ until no longer duplicate
	set := make(map[byte]struct{})
	res := 0
	l := 0

	for r := 0; r < len(s); r++ {
		for {
			if _, ok := set[s[r]]; !ok {
				break
			}
			delete(set, s[l])
			l++
		}
		set[s[r]] = struct{}{}
		res = max(res, r-l+1)
	}
	return res
}
