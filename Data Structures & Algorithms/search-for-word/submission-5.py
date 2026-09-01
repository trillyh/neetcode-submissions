class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:

        def dfs(i, j, idx):

            if i >= len(board) or j >= len(board[0]) or i < 0 or j < 0 or board[i][j] != word[idx]:
                return False

            if board[i][j] == word[idx] and idx == len(word)-1:
                return True
            
            temp = board[i][j]
            board[i][j] = "#"
            top = dfs(i-1, j, idx+1)
            bottom = dfs(i+1, j, idx+1)
            left = dfs(i, j-1, idx+1)
            right = dfs(i, j+1, idx+1)
            board[i][j] = temp 
            return top or left or right or bottom
        

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == word[0]:
                    res = dfs(i, j, 0)
                    if res:
                        return True

        return False

        


        