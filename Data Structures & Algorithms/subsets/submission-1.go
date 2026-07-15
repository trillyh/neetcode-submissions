// add: append(s, e)
// pop: s = s[:len(s)-1]



func subsets(nums []int) [][]int {
	res := [][]int{}
	
	var dfs func(curr []int, i int)
	dfs = func(curr[]int, i int) {
		if i >= len(nums) {
			snapshot := make([]int, len(curr))
			copy(snapshot, curr)
			res = append(res, snapshot)
			return
		}
		// add the current element and go next
		curr = append(curr, nums[i])
		dfs(curr, i+1)

		// pop the last element and go next
		curr = curr[:len(curr)-1]
		dfs(curr, i+1)
	}
	dfs([]int{}, 0)
	return res
}
