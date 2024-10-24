package z2z

import (
	"fmt"
	"testing"
)

func TestGaussLines(t *testing.T) {

	m := NewMat(5, 63*3)
	m.Randomize()
	ref := m.Clone()
	// fmt.Println(m)

	m.swapLines(3, 1)
	for c := 0; c < m.c; c++ {
		v, w := m.Get(3, c), m.Get(1, c)
		if v != ref.Get(1, c) || w != ref.Get(3, c) {
			t.Fatal("swapLines failed")
		}
	}
	m.swapLines(3, 1)
	if !m.Equal(ref) {
		t.Fatal("double swap failed")
	}

	// add 3 to 2
	ref = m.Clone()
	m.addLines(2, 3)
	for c := 0; c < m.c; c++ {
		v, w := ref.Get(2, c), ref.Get(3, c)
		if m.Get(2, c) != v^w {
			t.Fatal("addLines failed")
		}
	}
	m.addLines(2, 3) // since add is xor, adding twice will restore the initial value !
	if !m.Equal(ref) {
		t.Fatal("double add failed")
	}

	m = NewMat(127, 63)
	m.Randomize()
	ref = m.Clone()
	// fmt.Println(m)

	m.swapLines(126, 62)
	for c := 0; c < m.c; c++ {
		v, w := m.Get(126, c), m.Get(62, c)
		if v != ref.Get(62, c) || w != ref.Get(126, c) {
			t.Fatal("swapLines failed")
		}
	}
	m.swapLines(126, 62)
	if !m.Equal(ref) {
		t.Fatal("double swap failed")
	}

}

func TestGaussCols(t *testing.T) {

	m := NewMat(50, 9)
	m.Randomize()
	ref := m.Clone()
	m.swapCols(2, 6)

	if m.Equal(ref) {
		t.Fatal("suspicious equals")
	}
	for l := 0; l < m.l; l++ {
		if m.Get(l, 2) != ref.Get(l, 6) || m.Get(l, 6) != ref.Get(l, 2) {
			t.Fatal("swapCols failed")
		}
	}
	m.swapCols(2, 6)
	if !m.Equal(ref) {
		t.Fatal("double sawpCols failed")
	}

	// adding
	ref = m.Clone()
	m.addCols(2, 3)
	if m.Equal(ref) {
		t.Fatal("suspicious equals")
	}
	for l := 0; l < m.l; l++ {
		if m.Get(l, 2) != ref.Get(l, 3)^ref.Get(l, 2) {
			t.Fatal("addCols failed")
		}
	}

	m.addCols(2, 3)
	if !m.Equal(ref) {
		t.Fatal("double addCols failed")
	}
}

// --- GAUSS SHORT -------------

func TestGaussInvert(t *testing.T) {
	var m *Mat

	m = NewId(65)
	if !testInvert(t, m) {
		t.Fatal("should be inversible")
	}

	m = NewId(64)
	if !testInvert(t, m) {
		t.Fatal("should be inversible")
	}

	m = NewId(11)
	m.swapCols(2, 7)
	m.addCols(1, 8)
	if !testInvert(t, m) {
		t.Fatal("should be inversible")
	}

	m = NewId(11)
	m.swapCols(2, 7)
	m.addCols(1, 8)
	if !testInvert(t, m) {
		t.Fatal("should be inversible")
	}

	m = NewMat(3, 3)
	m.Randomize()
	testInvert(t, m)

	m = NewMat(6, 3)
	m.Randomize()
	if testInvert(t, m) {
		t.Fatal("should NOT be inversible")
	}

	m = NewMat(8, 8)
	m.Randomize()
	testInvert(t, m)

	m = NewMat(127, 63)
	m.Randomize()
	if testInvert(t, m) {
		t.Fatal("should NOT be inversible")
	}

	m = NewMat(63, 63)
	m.Randomize()
	testInvert(t, m)

	m = NewMat(63, 63)
	m.Randomize()
	testInvert(t, m)

}

func testInvert(_ *testing.T, m *Mat) bool {
	var r *Mat
	iv := m.Inverse()
	ok := iv != nil
	id := NewId(m.c)
	if ok {
		fmt.Println("m is inversible", m.l, m.c)
		r = m.MatMul(iv)
		if !r.Equal(id) {
			fmt.Println(m)
			fmt.Println(id)
			fmt.Println(iv)
			fmt.Println(ok)
			fmt.Println(r)
			fmt.Println("inverse did not verify", m.l, m.c)
		} else {
			fmt.Println("inverse verified", m.l, m.c)
		}
	} else {
		fmt.Println("m is not inversible", m.l, m.c)
	}
	return ok
}

// ----- GAUSS FULL ------------

func TestGaussFull(t *testing.T) {
	testGaussFull(t, 1, 1)
	testGaussFull(t, 5, 5)
	testGaussFull(t, 7, 9)
	testGaussFull(t, 64, 64)
	testGaussFull(t, 128, 128)
	testGaussFull(t, 127, 127)
	testGaussFull(t, 129, 129)
	testGaussFull(t, 129, 65)
}

func TestRank(t *testing.T) {
	var (
		ok bool
		rk int
	)
	m := NewId(4)
	_, _, ok, rk = m.Gauss()
	if !ok || rk != 4 {
		t.Fatal("unexpected rank 4x4")
	}
	m = m.CloneDims(6, 0)
	_, _, ok, rk = m.Gauss()
	if ok || rk != 4 {
		t.Fatal("unexpected rank 6x4")
	}
}

func testGaussFull(t *testing.T, l int, c int) {
	for i := 0; i < 20; i++ {
		l, c = c, l // swap l and c

		m := NewMat(l, c)
		m.Randomize()

		re, iv, ok, rk := m.Gauss()
		ivxm := iv.MatMul(m)
		id := NewId(m.l)

		fmt.Printf("testGaussFull : round = %d, dims = %dx%d\t invertible : %v\t rank = %d\n", i, l, c, ok, rk)

		// iv itself should ALWAYS be invertible
		if iv.Inverse() == nil {
			t.Fatal("iv should have been invertible")
		}

		if rk > m.c || rk > m.l {
			t.Fatalf("rank %d is inconsistant with dimensions %d x %d ", rk, m.l, m.c)
		}

		if !ok && rk == m.c && rk == m.l {
			t.Fatalf("ok = %v while rank % d ==  dimensions (%d x %d)", ok, rk, m.l, m.c)
		}

		if !ivxm.Equal(re) {
			fmt.Println("m")
			fmt.Println(m)
			fmt.Println("re")
			fmt.Println(re)
			fmt.Println("iv")
			fmt.Println(iv)
			fmt.Println("ok : ", ok)
			fmt.Println()
			fmt.Println("iv * m")
			fmt.Println(ivxm)
			fmt.Println("iv * m == re ? ", ivxm.Equal(re))
			t.Fatal("error : iv * m != re ")
		}

		if ok && !id.Equal(re) {
			fmt.Println("re")
			fmt.Println(re)
			fmt.Println("iv")
			fmt.Println(iv)
			fmt.Println("ok : ", ok)

			fmt.Println("\nm")
			fmt.Println(m)
			fmt.Println("iv * m")
			fmt.Println(ivxm)
			fmt.Println("id == re ? ", id.Equal(re))
			t.Fatal("error : ok is true, but re != id")
		}

	}
}
