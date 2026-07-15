func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	var dfs func(i int, total int, curr []int) 

	dfs = func(i int, total int, curr []int) {
		if total == target {
			fmt.Println(curr)
			snapshot := make([]int, len(curr))
			copy(snapshot, curr)
			res = append(res, snapshot)
			return
		}

		if i >= len(nums) || total > target {
			return
		}

		curr = append(curr, nums[i])
		dfs(i, total + nums[i], curr)

		curr = curr[:len(curr)-1]
		dfs(i+1, total, curr)
	}
	dfs(0, 0, []int{})
	return res
}
