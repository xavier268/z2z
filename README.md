[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/xavier268/z2z)

# z2z
Matrix operations in Z/2Z

# Basic operations

Vectors are 1 x c Matrixes, for efficiency.

Multiplication is implemented by transposing one of the terms first, generating significant speed gains.
Use MatMulTr to multiply by a vector.

Low level operations modify the matrix in plkace, high level operations generate a new matrix. See godoc for details.

# Gauss-Jordan reduction

Inverse uses a Gauss-Jordan reduction to computes the inverse of a square matrix.

Gauss accepts any matrix, perfoms a Gauss-Jordan transformation, returning the decomposition with a left pseudo-inverse (or actual inverse if inversible), the rank, and the determinant (a flag to say the matrix in inversible or not). See Godoc for details.
