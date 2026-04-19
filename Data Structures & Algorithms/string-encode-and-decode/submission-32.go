type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var decoded_strs string
	for _, str := range strs {
		decoded_strs += fmt.Sprintf("%d#%s", len(str),str)
	}
	return decoded_strs	
}

func (s *Solution) Decode(encoded string) []string {
	decoded_strs := []string{}
	l := 0
	for l < len(encoded) {
		r := l
		for encoded[r] != '#' {
			r++
		}

		length, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		word := encoded[l:l+length]
		decoded_strs = append(decoded_strs, word)

		l += length		
		fmt.Println(word)
		fmt.Println(decoded_strs)
	}
	return decoded_strs
}