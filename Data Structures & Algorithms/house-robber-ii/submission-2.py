class Solution:
    def rob(self, nums: List[int]) -> int:

        if len(nums) == 1:
            return nums[0]
        if len(nums) <= 2:
            return max(nums[0], nums[1])

        def dp(v: List[int]):
            cache = [0] * len(v)
            cache[0] = v[0]
            cache[1] = max(v[0],v[1])
            
            for i in range(2, len(v)):
                cache[i] = max(v[i] + cache[i-2], cache[i-1])
            
            return cache[-1]

        return max(dp(nums[1:]), dp(nums[:-1]))