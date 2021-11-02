package z2z

import (
	"fmt"
	"testing"
)

func TestMatMulTr(t *testing.T) {

	verifyMatMulTr(t, 6, 8, 5)

}

func verifyMatMulTr(t *testing.T, l int, c int, cc int) {
	for i := 0; i < 20; i++ {
		m := NewMat(l, c)
		p := NewMat(c, cc)
		m.Randomize()
		p.Randomize()

		r1 := m.matMulNaive(p)
		r2 := m.matMulTr(p.T())

		if !r1.Equal(r2) {
			fmt.Println("Want :\n", r1)
			fmt.Println("Got  :\n", r2)
			t.Fatal("matrix optimization does not give correct result")

		}
	}

}
