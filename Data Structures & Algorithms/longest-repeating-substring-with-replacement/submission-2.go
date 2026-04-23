func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0
	l := 0

	mostC := 0
	for r := 0; r < len(s); r++ {
		count[s[r]] += 1
		mostC = max(mostC, count[s[r]])
		for (r - l + 1) - mostC > k {
			count[s[l]]	-= 1
			l++
		}
		res = max(res, r-l+1)
	}

	return res
}
