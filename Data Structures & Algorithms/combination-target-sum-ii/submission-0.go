func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(i int, total int, curr []int)
	dfs = func(i int, total int, curr []int) {
		if total == target {
			snapshot := make([]int, len(curr))
			copy(snapshot, curr)
			res = append(res, snapshot)
			return
		}	

		if total > target || i >= len(candidates) {
			return	
		}
		
		curr = append(curr, candidates[i])
		dfs(i+1, total + candidates[i], curr)	

		curr = curr[:len(curr)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, total, curr)
	}	
	dfs(0, 0, []int{})
	return res
}
