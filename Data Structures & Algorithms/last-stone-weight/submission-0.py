class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        heapq.heapify_max(stones) 
       
        while len(stones) >= 2:
            first = heapq.heappop_max(stones)
            second = heapq.heappop_max(stones)
            if first - second != 0:
                heapq.heappush_max(stones, (first-second))
        if len(stones) == 1:
            return stones[0]

        return 0
    