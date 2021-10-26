package z2z

import (
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
