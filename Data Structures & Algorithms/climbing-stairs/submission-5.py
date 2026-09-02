class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n
        cache = [0] * (n + 1) # n = 3 cache=[0,1,2,3]

        cache[1] = 1
        cache[2] = 2

        for i in range(3,n+1): # n = 3, so range(n+1) goes up to 3
            print(f"{i}")
            cache[i] = cache[i-1] + cache[i-2]
        
        print(f"{cache}")
        return cache[-1]



        

        