class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        res = 0
        lowest = prices[0]
        for price in prices:
            profit = price - lowest
            res = max(res, profit)
            lowest = min(lowest, price)
        return res