"""
    Implementation of Straight Forward Algorthim for Matrix Multiplication
    It's only built for (n * n) matrices
"""

def multiply(m1,m2):
    if len(m1) <= 2 or len(m2) <= 2:
        C = [[0 for row in range(len(m2[0]))] for col in range(len(m1))]
        for i in range(len(m1)):
            for j in range(len(m2[0])):
                for k in range(len(m1[0])):
                    C[i][j] += m1[i][k] * m2[k][j]
        return C


    A = getTheFirstPartOfMatrix(m1)
    B = getTheSecondPartOfMatrix(m1)
    C = getTheThirdPartOfMatrix(m1)
    D = getTheForthPartOfMatrix(m1)
    E = getTheFirstPartOfMatrix(m2)
    F = getTheSecondPartOfMatrix(m2)
    G = getTheThirdPartOfMatrix(m2)
    H = getTheForthPartOfMatrix(m2)
    p1 = multiply(A,E)
    p2 = multiply(B,G)
    p3 = multiply(A,F)
    p4 = multiply(B,H)
    p5 = multiply(C,E)
    p6 = multiply(D,G)
    p7 = multiply(C,F)
    p8 = multiply(D,H)
    return merge(addTwoMatrices(p1,p2),addTwoMatrices(p3,p4),addTwoMatrices(p5,p6),addTwoMatrices(p7,p8))


def merge(A,B,C,D):
    n = len(A)+len(C)
    m = [[0 for x in range(n)] for y in range(n)]
    for i in range(len(A)):
        m[i] = A[i] + B[i]
    for i in range(len(C)):
        m[i+len(A)] = C[i] + D[i]

    return m



"""
    Helper Functions
"""
def simpleMultiply(m1,m2):
    C = [[0 for row in range(len(m2[0]))] for col in range(len(m1))]
    for i in range(len(m1)):
        for j in range(len(m2[0])):
            for k in range(len(m1[0])):
                C[i][j] += m1[i][k] * m2[k][j]
    return C


def getTheFirstPartOfMatrix(m):
    A = []
    n = len(m)
    for i in range(0,n/2):
        A.append(m[i][:n/2])
    return A

def getTheSecondPartOfMatrix(m):
    B = []
    n = len(m)
    for i in range(0,n/2):
        B.append(m[i][n/2:])
    return B


def getTheThirdPartOfMatrix(m):
    C = []
    n = len(m)
    for i in range(n/2,n):
        C.append(m[i][:n/2])
    return C

def getTheForthPartOfMatrix(m):
    D = []
    n = len(m)
    for i in range(n/2,n):
        D.append(m[i][n/2:])
    return D

def addTwoMatrices(m1,m2):
    m3 = [ [ 0 for x in range(len(m1[0]))] for y in range(len(m1))]
    for i in range(len(m1)):
        m3[i] = [m1[i][j] + m2[i][j] for j in range(len(m1[i]))]
    return m3

def subTwoMatrices(m1,m2):
    m3 = [ [ 0 for x in range(len(m1[0]))] for y in range(len(m1))]
    for i in range(len(m1)):
        m3[i] = [m1[i][j] - m2[i][j] for j in range(len(m1[i]))]
    return m3



"""
2,3,4
4,3,2
3,4,5
"""
m = [[0 for x in range(3)] for y in range(3)]
m[0] = [2,3,4]
m[1] = [4,3,2]
m[2] = [3,4,5]

m2 = m
print multiply(m,m2)