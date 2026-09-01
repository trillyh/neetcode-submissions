class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        
        res = []

        def dfs(candidates, curr):
            if not candidates: # length of candidates is 0
                res.append(curr)
                return

            
            for i in range(len(candidates)):
                dfs(candidates[:i] + candidates[i+1:],curr + [candidates[i]])

        dfs(nums, [])
        return res

            
        