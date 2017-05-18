def minDrop(s, si, oc, cache, pseq):
    N = len(s)

    if oc < 0:
        return N - si + 1

    if si == N :
        if oc == 0:
            pseq[si][oc] = {''}
        return oc

    if cache[si][oc] != -1:
        return cache[si][oc]

    
    if s[si] in '()':
        dc0 = 1 + minDrop(s, si + 1, oc, cache, pseq)
        pseq0 = pseq[si + 1][oc]

        if s[si] == '(':
            dc1 = minDrop(s, si + 1, oc + 1, cache, pseq)
            pseq1 = ['(' + x for x in pseq[si + 1][oc + 1]]
        else:
            dc1 = minDrop(s, si + 1, oc - 1, cache, pseq)
            pseq1 = [')' + x for x in pseq[si + 1][oc - 1]]

        cache[si][oc] = min(dc0, dc1)

        # note '=' - in case of eqaulity we keep both combination sets
        if dc0 >= dc1 :
            pseq[si][oc] = pseq[si][oc].union(pseq1)

        if dc0 <= dc1 :
            pseq[si][oc] = pseq[si][oc].union(pseq0) 

    else:
        cache[si][oc] = minDrop(s, si + 1, oc, cache, pseq)
        pseq[si][oc] = [s[si] + x for x in pseq[si + 1][oc]]

    return cache[si][oc]

class Solution(object):
    def removeInvalidParentheses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        N = len(s)
        cache = [[-1 for x in range(N)] for x in range(N)]
        pseq = [[set() for x in range(N + 1)] for x in range(N + 1)]

        c = minDrop(s, 0, 0, cache, pseq)

        return list(pseq[0][0])