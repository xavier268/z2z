# z2z
Matrix operations in Z/2Z

Basic matrix operations, see details in godoc.
Multiplication is implemented by transposing one of the terms first, generating significant speed gains.

Vectors are 1 x c Matrixes, for efficiency.

GaussShort computes the inverse of a square matrix.

GaussFull accepts any matrix, perfoms a Gauss-Jordan transformation, returning the decomposition with a left pseudo-inverse, the rank, and the determinant (a flag to say the matrix in inversible or not).
