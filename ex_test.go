package z2z

import "fmt"

var _ = 0

func ExampleMat_Gauss() {

	m := NewId(4)
	m.Set(1, 3, 1)
	m.Set(2, 1, 1)

	fmt.Print("m\n", m)
	istar, inv, ok := m.Gauss()

	fmt.Print("istar\n", istar)
	fmt.Print("inv\n", inv)
	fmt.Println("inversible :", ok)
	fmt.Println("check :", NewId(4).Equal(inv.MatMul(m)))
	fmt.Println("check :", NewId(4).Equal(m.MatMul(inv)))

	// output:
	// m
	// 0001
	// 1010
	// 0110
	// 1000
	// istar
	// 0001
	// 0010
	// 0100
	// 1000
	// inv
	// 0001
	// 1010
	// 1110
	// 1000
	// inversible : true
	// check : true
	// check : true
}
