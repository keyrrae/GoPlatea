import random
import sys

random.seed(1234)

def createRandomMatrix(n):
    matrix = []
    for i in xrange(n):
        matrix.append([random.random() for el in xrange(n)])
    return matrix

def saveMatrix(matrixA, matrixB, filename):
    f = open(filename, 'w')
    for i, matrix in enumerate([matrixA, matrixB]):
        if i != 0:
            f.write("\n")
        for line in matrix:
            f.write("\t".join(map(str, line)) + "\n")

def ikjMatrixProduct(A, B):
    n = len(A)
    C = [[0 for i in xrange(n)] for j in xrange(n)]
    for i in xrange(n):
        for k in xrange(n):
            for j in xrange(n):
                C[i][j] += A[i][k] * B[k][j]
    return C

dimension = 100
if len(sys.argv) > 1:
    dimension = int(sys.argv[1])

matrixA = createRandomMatrix(dimension)
matrixB = createRandomMatrix(dimension)
matrixC = ikjMatrixProduct(matrixA, matrixB)
