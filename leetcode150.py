from typing import List

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:

        if len(tokens) == 1:
            return int(tokens[0])
        else:

            operators = ['+', '-', '*', '/']

            left_pos = 0
            right_pos = 1
            op_pos = 2

            while tokens[op_pos] not in operators:
                left_pos+=1
                right_pos+=1
                op_pos+=1

            left = int(tokens[left_pos])
            right = int(tokens[right_pos])
            op = tokens[op_pos]

            if op == '+':
                res = left + right
            elif op == '-':
                res = left - right
            elif op == '*':
                res = left * right
            else: # Division
                if right != 0:
                    res = int(left / right)
                else:
                    res = 0

            tokens = tokens[:left_pos] + [res] + tokens[op_pos+1:]

        return self.evalRPN(tokens)



if __name__ == "__main__":

    sol = Solution()
    tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
    resultado = sol.evalRPN(tokens)
    print(f"Resultado: {resultado}")
