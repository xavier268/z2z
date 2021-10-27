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

func testInvert(t *testing.T, m *Mat) bool {
	var r *Mat
	id, iv, ok := m.Gauss()

	if m.c == m.l {
		idd := NewId(m.c)
		if (idd.Equal(id)) != ok {
			fmt.Println(m)
			fmt.Println(id)
			fmt.Println(iv)
			fmt.Println(ok)
			t.Fatal("inconsitant id and ok")
		}
	}

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
