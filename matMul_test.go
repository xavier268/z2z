package z2z

import (
	"fmt"
	"testing"
)

func TestMatMulTr(t *testing.T) {

	verifyMatMulTr(t, 6, 8, 5)
	verifyMatMulTr(t, 8, 6, 5)
	verifyMatMulTr(t, 6, 8, 10)
	verifyMatMulTr(t, 15, 8, 5)
	verifyMatMulTr(t, 64, 64, 64)
	verifyMatMulTr(t, 63, 127, 63)
	verifyMatMulTr(t, 64, 127, 64)
	verifyMatMulTr(t, 64, 128, 64)
	verifyMatMulTr(t, 256, 250, 256)

}

func verifyMatMulTr(t *testing.T, l int, c int, cc int) {
	for i := 0; i < 10; i++ {
		m := NewMat(l, c)
		p := NewMat(c, cc)
		m.Randomize()
		p.Randomize()

		r1 := m.matMulNaive(p)
		r2 := m.matMulTr(p.T())

		if !r1.Equal(r2) {
			fmt.Println("Want :\n", r1)
			fmt.Println("Got  :\n", r2)
			t.Fatalf("matrix optimization does not give correct result (%d x %d x %d)", l, c, cc)

		}
	}

}
