from collections import deque
import random

def is_valid(s):
    count = 0
    for ch in s:
        if ch == "(":
            count = count + 1
        elif ch == ")":
            count = count - 1
        if count < 0:
            return False
    return count == 0

def removeInvalidParentheses(s):
    """
    :type s: str
    :rtype: List[str]
    """
    q, result = deque(), []
    q.append(s)
    while len(q) and len(result) == 0:
        next_level = set([])
        for _ in range(len(q)):
            x = q.popleft()
            if is_valid(x):
                result.append(x)
            elif len(result) == 0:
                for i in range(len(x)):
                    if x[i] in ("(", ")"):
                        next_level.add(x[0:i] + x[i+1:])
        for nl in next_level:
            q.append(nl)
    return result if result else [""]

def main():
    s = "(()()(()(()((()))(()((())(()(()((()())((())))()"
    a = list(s)
    dut = ""
    random.seed(1234)
    for i in range(40):
        dut = dut + a[i % len(a)]
    print(dut)
    removeInvalidParentheses(dut)

main()
