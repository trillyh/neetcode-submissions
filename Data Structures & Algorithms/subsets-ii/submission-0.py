class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = [] 
        nums.sort()
        def dfs(curr, i):
            if i == len(nums):
                res.append(list.copy(curr))
                return

            curr.append(nums[i])
            dfs(curr, i+1)
            curr.pop()

            while i+1 < len(nums) and nums[i] == nums[i+1]:
                i+=1
            dfs(curr, i+1)
        dfs([], 0)
        return res
        
        
