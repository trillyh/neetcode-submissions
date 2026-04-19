class Solution:
    def isValid(self, s: str) -> bool:
        closing = {")": "(", "]": "[", "}": "{"}

        stack = []
        for c in s: 
            if c in closing and len(stack) != 0:
                val = stack.pop()
                if val != closing[c]:
                    return False
            
            else:
                stack.append(c)
        
        return len(stack) == 0
            

        

            




        