class Solution:
    def search(self, nums: List[int], target: int) -> int:
        # two pass, first one find where the pivot is
        # then perform a regular binary search

        l, r = 0, len(nums)-1
        
        while l < r:
            m = l + (r-l)//2
            if nums[m] < nums[r]:
                r = m
            else: # nums[m] >= nums[r] [2L,3M,1R]
                l = m + 1
        pivot = l
        print(pivot)
        l, r = 0, len(nums)-1
        
        # determine where the target lines [l-->pivot-->r]
        if nums[pivot] <= target and target <= nums[r]:
            l = pivot
        else: # must in the left segment: pivot > target or target > nums[r]
        # but pivot > target cannot be true (b/c at least target = pivot)
        # that left us with target > nums[r] and only be in the right segment
            r = pivot - 1 # we also exclude the pivot, because pivot belongs to the right segment
            # and target is not in right segment based on our condition above.
        
        print(f"l: {l} r: {r}")
        while l <= r:
            m = l + (r-l)//2
            print(f"m: {m}")
            if nums[m] == target:
                return m
            elif nums[m] < target:
                l = m + 1
            else:
                r = m - 1
        return -1
            
        