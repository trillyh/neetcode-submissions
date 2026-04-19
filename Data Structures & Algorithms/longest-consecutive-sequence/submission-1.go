func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	// build the array
	for _, i := range(nums) {
		set[i] = true
	}

	// second pass, start building the sequen if and only if starting exits
	res := 0
	for _, i := range(nums) {
		currLen := 1
		if set[i-1] {
			j := i
			for set[j]{
				fmt.Printf("%d\n", j)
				currLen++
				j++
			}	
		}
		res = max(res, currLen)
	}
	return res

}