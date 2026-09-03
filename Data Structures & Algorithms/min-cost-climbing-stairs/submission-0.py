class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
       # subproblem is: at each step, we calculate the minimum cost to reach this stair case. 
       # suppose cache is like this[1,4,5,7,6], res = min of last two values

        length = len(cost)
        cache = [0] * length

        cache[0] = cost[0]
        cache[1] = cost[1]

        for i in range(2, length):
            cache[i] = cost[i] + min(cache[i-1], cache[i-2])

        print(cache)
        return min(cache[length-1], cache[length-2])
