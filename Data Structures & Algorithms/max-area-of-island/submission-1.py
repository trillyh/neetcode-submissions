class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]):
        def dfs(i, j):
            
            if i >= len(grid) or j >= len(grid[0]) or i < 0 or j < 0 or grid[i][j] == 0:
                return 0

            grid[i][j] = 0
            curr = dfs(i-1, j) + dfs(i, j-1) + dfs(i+1,j) + dfs(i, j+1)
            return 1 + curr

        res = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == 1:
                    res = max(res, dfs(i, j))

        return res
                    
            