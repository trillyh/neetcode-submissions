class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        
        res = []
        def dfs(curr, i):
            if i == len(nums):
                res.append(list.copy(curr))
                return

            curr.append(nums[i])
            dfs(curr, i+1)
            curr.pop()
            dfs(curr, i+1)

            return

        dfs([], 0)
        return res
        
                