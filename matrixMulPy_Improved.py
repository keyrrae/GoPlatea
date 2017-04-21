"""
    Implementation of Strassen's Algorthim for Matrix Multiplication
    It's only built for (n * n) matrices
"""
import unittest



def Strassen(m1,m2):
    if len(m1) != len(m2):
        return "Sorry Check your Matrices !"
    if len(m1) % 2 != 0:
        return unpaddingMatrix(multiply(paddingMatrix(m1),paddingMatrix(m2)))
    return multiply(m1,m2)

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
    p1 = multiply(A,subTwoMatrices(F,H))
    p2 = multiply(addTwoMatrices(A,B),H)
    p3 = multiply(addTwoMatrices(C,D),E)
    p4 = multiply(D,subTwoMatrices(G,E))
    p5 = multiply(addTwoMatrices(A,D),addTwoMatrices(E,H))
    p6 = multiply(subTwoMatrices(B,D),addTwoMatrices(G,H))
    p7 = multiply(subTwoMatrices(A,C),addTwoMatrices(E,F))

    R1 = addTwoMatrices(subTwoMatrices(addTwoMatrices(p5,p4),p2),p6)
    R2 = addTwoMatrices(p1,p2)
    R3 = addTwoMatrices(p3,p4)
    R4 = subTwoMatrices(subTwoMatrices(addTwoMatrices(p1,p5),p3),p7)
    return merge(R1,R2,R3,R4)


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

def paddingMatrix(m):
    x = [[0 for col in range(len(m)+1)] for row in range(len(m)+1)]
    for i in range(len(m)):
        for j in range(len(m)):
            x[i][j] = m[i][j]
    return x

def unpaddingMatrix(m):
    # we gonna remove padding
    x = [[0 for col in range(len(m)-1)] for row in range(len(m)-1)]
    for i in range(len(x)):
        for j in range(len(x)):
            x[i][j] = m[i][j]
    return x


# we will multiply 1 * 1 and 2 * 2 and 3 * 3 and 4 * 4 and 5 * 5
"""
2
"""
m = [[0 for x in range(1)] for y in range(1)]
m[0] = [2]
m2 = m
print Strassen(m,m2)
"""
4,3
5,2
"""
m = [[0 for x in range(2)] for y in range(2)]
m[0] = [4,3]
m[1] = [5,2]
m2 = m
print Strassen(m,m2)
"""
4,3,2
5,1,2
3,2,0
"""
m = [[0 for x in range(3)] for y in range(3)]
m[0] = [4,3,2]
m[1] = [5,1,2]
m[2] = [3,2,0]
m2 = m
print Strassen(m,m2)
"""
7,32,4,100
5,30,29,823
23,209,312,3
21,9,20,32
"""
m = [[0 for x in range(4)] for y in range(4)]
m[0] = [7,32,4,100]
m[1] = [5,30,29,823]
m[2] = [23,209,312,3]
m[3] = [21,9,20,32]
m2 = m
print Strassen(m,m2)