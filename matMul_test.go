package z2z

import (
	"fmt"
	"testing"
)

func TestVerifyMatMulOptim(t *testing.T) {

	verifyMatMulOptim(t, 6, 6, 6)

}

func verifyMatMulOptim(t *testing.T, l int, c int, cc int) {
	m := NewMat(l, c)
	p := NewMat(c, cc)
	m.Randomize()
	p.Randomize()

	r1 := m.MatMul(p)
	r2 := m.matMulOptim(p)

	if !r1.Equal(r2) {
		fmt.Println("Want :\n", r1)
		fmt.Println("Got  :\n", r2)
		t.Fatal("matrix optimization does not give correct result")

	}

}
