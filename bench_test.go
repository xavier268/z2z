package z2z

import "testing"

// --- MatMul ---

func BenchmarkMatMulNaive5(b *testing.B) {
	l := 5
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulNaive(m)
	}
}
func BenchmarkMatMulNaive50(b *testing.B) {
	l := 50
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulNaive(m)
	}
}
func BenchmarkMatMulNaive500(b *testing.B) {
	l := 500
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulNaive(m)
	}
}

func BenchmarkMatMulTr5(b *testing.B) {
	l := 5
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMulTr(m.T())
	}
}

func BenchmarkMatMulTr50(b *testing.B) {
	l := 50
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMulTr(m.T())
	}
}
func BenchmarkMatMulTr500(b *testing.B) {
	l := 500
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMulTr(m.T())
	}
}

func BenchmarkMatMulTr5000(b *testing.B) {
	l := 5000
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMulTr(m.T())
	}
}

// -- GAUSS short --

func BenchmarkGaussShort5(b *testing.B) {

	m := NewId(5)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_ = m.GaussShort()
	}
}

func BenchmarkGaussShort50(b *testing.B) {

	m := NewId(50)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_ = m.GaussShort()
	}
}

func BenchmarkGaussShort500(b *testing.B) {

	m := NewId(500)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_ = m.GaussShort()
	}
}

func BenchmarkGaussShort5000(b *testing.B) {

	m := NewId(5000)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_ = m.GaussShort()
	}
}

// ---- GAUSS FULL ---------

func BenchmarkGaussFull5(b *testing.B) {

	m := NewId(5)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _, _ = m.GaussFull()
	}
}
func BenchmarkGaussFull50(b *testing.B) {

	m := NewId(50)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _, _ = m.GaussFull()
	}
}

func BenchmarkGaussFull500(b *testing.B) {

	m := NewId(500)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _, _ = m.GaussFull()
	}
}

func BenchmarkGaussFull5000(b *testing.B) {

	m := NewId(5000)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _, _ = m.GaussFull()
	}
}

// Naive multiplication is always worse the the transposed implementation, including the cost of the transpose operation.
// Naive is 0(n^3), and TransposedMul, while asymptotically O(n^3),  is ~ 50x more efficient and looks more like O(n^2)
// for smaller dimensions.
// GaussShort is only slightly faster than GaussFull, but will be significantly faster on non invertible large matrices.

//
//
// goos: linux
// goarch: amd64
// pkg: github.com/xavier268/z2z
// cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
// BenchmarkMatMulNaive5-8          1000000              1148 ns/op              96 B/op          2 allocs/op
// BenchmarkMatMulNaive50-8             631           1902974 ns/op             464 B/op          2 allocs/op
// BenchmarkMatMulNaive500-8              1        1892414292 ns/op           65584 B/op          3 allocs/op

// BenchmarkMatMulTr5-8             3542947               340.8 ns/op           192 B/op          4 allocs/op
// BenchmarkMatMulTr50-8              29668             39742 ns/op             928 B/op          4 allocs/op
// BenchmarkMatMulTr500-8               207           5751970 ns/op           65790 B/op          4 allocs/op
// BenchmarkMatMulTr5000-8                1        2671063268 ns/op         9486432 B/op          5 allocs/op

// BenchmarkGaussShort5-8           7469222               208.8 ns/op           144 B/op          3 allocs/op
// BenchmarkGaussShort50-8            76323             15722 ns/op             880 B/op          3 allocs/op
// BenchmarkGaussShort500-8             319           3498960 ns/op           65686 B/op          3 allocs/op
// BenchmarkGaussShort5000-8              1        2124097429 ns/op         9486536 B/op          7 allocs/op

// BenchmarkGaussFull5-8            4821045               234.7 ns/op           192 B/op          4 allocs/op
// BenchmarkGaussFull50-8             74164             14936 ns/op             928 B/op          4 allocs/op
// BenchmarkGaussFull500-8              294           3893741 ns/op           65743 B/op          4 allocs/op
// BenchmarkGaussFull5000-8               1        2434553147 ns/op         9486480 B/op          6 allocs/op

// PASS
// ok      github.com/xavier268/z2z        30.459s
