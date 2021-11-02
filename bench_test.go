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
		m = m.matMulTr(m.T())
	}
}

func BenchmarkMatMulTr50(b *testing.B) {
	l := 50
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulTr(m.T())
	}
}
func BenchmarkMatMulTr500(b *testing.B) {
	l := 500
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulTr(m.T())
	}
}

func BenchmarkMatMulTr5000(b *testing.B) {
	l := 5000
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.matMulTr(m.T())
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

// Naive multiplication is always worse the the transposed implementation, including the cost of the transpose operation.
// Naive is 0(n^3), and TransposedMul, while asymptotically O(n^3),  is ~ 50x more efficient and looks more like O(n^2)
// for smaller dimensions.

//
//

// goos: linux
// goarch: amd64
// pkg: github.com/xavier268/z2z
// cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
// BenchmarkMatMulNaive5-8           983400              1131 ns/op              96 B/op          2 allocs/op
// BenchmarkMatMulNaive50-8             636           1876318 ns/op             464 B/op          2 allocs/op
// BenchmarkMatMulNaive500-8              1        1869947495 ns/op           65584 B/op          3 allocs/op
// BenchmarkMatMulTr5-8             1811682               660.1 ns/op           192 B/op          4 allocs/op
// BenchmarkMatMulTr50-8              15530             76447 ns/op             928 B/op          4 allocs/op
// BenchmarkMatMulTr500-8                42          28080307 ns/op           66412 B/op          4 allocs/op
// BenchmarkMatMulTr5000-8                1        26017520025 ns/op        9486432 B/op          5 allocs/op
// BenchmarkGaussShort5-8           8083322               130.2 ns/op           144 B/op          3 allocs/op
// BenchmarkGaussShort50-8            93700             15489 ns/op             880 B/op          3 allocs/op
// BenchmarkGaussShort500-8             322           3680521 ns/op           65685 B/op          3 allocs/op
// BenchmarkGaussShort5000-8              1        2343296750 ns/op         9486432 B/op          5 allocs/op
// PASS
// ok      github.com/xavier268/z2z        44.000s
