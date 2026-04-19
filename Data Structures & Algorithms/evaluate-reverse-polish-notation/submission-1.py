class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        # use a stack, push two value in, when encouter expression:
        # pop the stack 2 times two get two value out
        # push in the result of the operation between two value

        stack = []

        for c in tokens:
            if c == '+':
                stack.append(stack.pop() + stack.pop())
            elif c == '-':
                a, b = stack.pop(), stack.pop()
                stack.append(b - a)
            elif c == '*':
                stack.append(stack.pop() * stack.pop())
            elif c == '/':
                a, b = stack.pop(), stack.pop()
                stack.append(int(float(b) / a))
            else:
                stack.append(int(c))
        
        return stack[-1]
        