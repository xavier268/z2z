package z2z

import "testing"

// --- MatMul ---

func BenchmarkMatMul5(b *testing.B) {
	l := 5
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMul(m)
	}
}

func BenchmarkMatMul50(b *testing.B) {
	l := 50
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMul(m)
	}
}

func BenchmarkMatMul500(b *testing.B) {
	l := 500
	m := NewMat(l, l)
	m.Randomize()

	for i := 0; i < b.N; i++ {
		m = m.MatMul(m)
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

// Naive multiplication confirms a O(n^3) complexity.
// Gauss-Jordan is surprinsingly efficient for random matrixes.
//
//
//
// goos: linux
// goarch: amd64
// pkg: github.com/xavier268/z2z
// cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
// BenchmarkMatMul5-8       1071608              1110 ns/op              96 B/op          2 allocs/op
// BenchmarkMatMul50-8          633           1891460 ns/op             464 B/op          2 allocs/op
// BenchmarkMatMul500-8           1        1869713857 ns/op           65592 B/op          4 allocs/op
// BenchmarkGauss5-8        6070524               194.4 ns/op           192 B/op          4 allocs/op
// BenchmarkGauss50-8         75078             15316 ns/op             928 B/op          4 allocs/op
// BenchmarkGauss500-8          308           3882637 ns/op           65738 B/op          4 allocs/op
// BenchmarkGauss5000-8           1        2368967139 ns/op         9486480 B/op          6 allocs/op
// PASS
// ok      github.com/xavier268/z2z        11.647s
