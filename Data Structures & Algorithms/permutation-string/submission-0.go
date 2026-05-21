func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if freq1[i] == freq2[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}

		index := s2[r] - 'a'
		freq2[index]++
		if freq1[index] == freq2[index] {
			matches++
		} else if freq1[index]+1 == freq2[index] {
			matches --
		}
		index = s2[l] - 'a'
		freq2[index]--
		if freq1[index] == freq2[index] {
			matches++
		} else if freq1[index]-1 == freq2[index] {
			matches--
		}

		l++
	}

	return matches == 26

	

}
