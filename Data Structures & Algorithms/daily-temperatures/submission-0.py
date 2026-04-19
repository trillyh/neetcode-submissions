class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = [] # pair of (idx, temp)

        for idx, temp in enumerate(temperatures):
            while stack and temp > stack[-1][1]:
                i, t = stack.pop()
                res[i] = idx - i
            stack.append((idx, temp))

        while len(stack) != 0:
            i, t = stack.pop()
            res[i] = 0

        return res

        