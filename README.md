[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/xavier268/z2z)

# z2z
Matrix operations in Z/2Z

# Basic operations

See details in godoc.

Vectors are 1 x c Matrixes, for efficiency.

Multiplication is implemented by transposing one of the terms first, generating significant speed gains.
Use MatMulTr to multiply by a vector.

# Gauss-Jordan reduction :

GaussShort computes the inverse of a square matrix.

GaussFull accepts any matrix, perfoms a Gauss-Jordan transformation, returning the decomposition with a left pseudo-inverse, the rank, and the determinant (a flag to say the matrix in inversible or not).
