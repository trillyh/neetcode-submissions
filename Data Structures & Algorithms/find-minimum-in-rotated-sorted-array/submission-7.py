class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1
        res = nums[0]
        
        while l <= r:
            m = l + (r-l)//2
            res = min(nums[m], res) # in case mid is exactly where the value is
            if nums[m] < nums[r]:
                r = m
            else: # nums[m] >= nums[r]
                l = m + 1

        return res



