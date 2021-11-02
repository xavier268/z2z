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

// -- GAUSS --

func BenchmarkGauss5(b *testing.B) {

	m := NewId(5)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Gauss()
	}
}

func BenchmarkGauss50(b *testing.B) {

	m := NewId(50)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Gauss()
	}
}

func BenchmarkGauss500(b *testing.B) {

	m := NewId(500)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Gauss()
	}
}

func BenchmarkGauss5000(b *testing.B) {

	m := NewId(5000)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		_, _, _ = m.Gauss()
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
// BenchmarkMatMulNaive5-8           912189              1131 ns/op              96 B/op          2 allocs/op
// BenchmarkMatMulNaive50-8             632           1888703 ns/op             464 B/op          2 allocs/op
// BenchmarkMatMulNaive500-8              1        1872507360 ns/op           65584 B/op          3 allocs/op
//
// BenchmarkMatMulTr5-8             1801214               652.2 ns/op           192 B/op          4 allocs/op
// BenchmarkMatMulTr50-8              15664             76284 ns/op             928 B/op          4 allocs/op
// BenchmarkMatMulTr500-8                39          28192895 ns/op           66472 B/op          4 allocs/op
// BenchmarkMatMulTr5000-8                1        25917038799 ns/op        9486432 B/op          5 allocs/op
//
// BenchmarkGauss5-8                5173652               230.4 ns/op           192 B/op          4 allocs/op
// BenchmarkGauss50-8                 72410             16368 ns/op             928 B/op          4 allocs/op
// BenchmarkGauss500-8                  298           3887356 ns/op           65742 B/op          4 allocs/op
// BenchmarkGauss5000-8                   1        2416653762 ns/op         9486480 B/op          6 allocs/op
// PASS
// ok      github.com/xavier268/z2z        41.864s
