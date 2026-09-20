# An m×n matrix has m rows and n columns
# The transpose A^T flips rows and columns: (A^T){ij} = A{ji}
# The identity matrix I has 1s on the diagonal and 0s elsewhere; it satisfies AI = IA = A for any appropriately-sized A
# Matrix addition is done element-wise: (A + B){ij} = A{ij} + B{ij}

from typing import TypeVar

Number = int | float
Matrix = list[list[Number]]

T = TypeVar("T")

A: Matrix = [[1, 2, 3], [4, 5, 6]]
B: Matrix = [[7, 8, 9], [10, 11, 12]]


def _check_rectangular(m: Matrix) -> None:
    if not m or not m[0]:
        raise ValueError("matrix must be non-empty")
    width = len(m[0])
    for i in range(len(m)):
        if len(m[i]) != width:
            raise ValueError("matrix rows must all have the same length")


def swap(m: T, n: T) -> tuple[T, T]:
    return (n, m)


def transpose(matrix: Matrix) -> Matrix:
    _check_rectangular(matrix)
    result: Matrix = []
    for i in range(len(matrix[0])):
        row: list[Number] = []
        for j in range(len(matrix)):
            row.append(matrix[j][i])
        result.append(row)
    return result


print(transpose(A))


def add(A: Matrix, B: Matrix) -> Matrix:
    _check_rectangular(A)
    _check_rectangular(B)
    if len(A) != len(B) or len(A[0]) != len(B[0]):
        raise ValueError("matrices must have the same dimensions to add")
    result: Matrix = []
    for i in range(len(A)):
        row: list[Number] = []
        for j in range(len(A[0])):
            row.append(A[i][j] + B[i][j])
        result.append(row)
    return result


print(add(A, B))


def multiply(A: Matrix, B: Matrix) -> Matrix:
    _check_rectangular(A)
    _check_rectangular(B)
    if len(A[0]) != len(B):
        raise ValueError("A columns must equal B rows to multiply")
    m, n, p = len(A), len(A[0]), len(B[0])
    result: Matrix = []
    for i in range(m):
        row: list[Number] = []
        for j in range(p):
            row.append(0)
        result.append(row)
    for i in range(m):
        for j in range(p):
            s: Number = 0
            for k in range(n):
                s += A[i][k] * B[k][j]
            result[i][j] = s
    return result


print(multiply(A, transpose(B)))
