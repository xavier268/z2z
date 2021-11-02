package z2z

import "fmt"

var _ = 0

func ExampleMat_GaussShort() {

	m := NewId(4)
	m.Set(1, 3, 1)
	m.Set(2, 1, 1)

	fmt.Print("m\n", m)
	inv := m.GaussShort()
	ok := inv != nil

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
	// inv
	// 0001
	// 1010
	// 1110
	// 1000
	// inversible : true
	// check : true
	// check : true
}

func ExampleNewFromInt() {

	m := NewFromInt(
		0b111,
		0b001100110011,
		0b111111110000,
	)

	fmt.Println(m)
	fmt.Println(m.CloneDims(5, 5))              // adjust to 5x5
	fmt.Println(m.CloneDims(0, 6).Dimensions()) // adjust to 3x6

	// Output:
	// 0000000000000000000000000000000000000000000000000000000000000111
	// 0000000000000000000000000000000000000000000000000000001100110011
	// 0000000000000000000000000000000000000000000000000000111111110000
	//
	// 00111
	// 10011
	// 10000
	// 00000
	// 00000
	//
	// 3 6

}

func ExampleMat_Apply() {
	m := NewMat(3, 4)
	m.Apply(func(i, j int) int {
		return (i + j) % 2
	})
	fmt.Println(m)
	// Output:
	// 1010
	// 0101
	// 1010
}
func ExampleMat_Apply_usePreExistingValues() {
	m := NewId(4)
	m.Apply(func(i, j int) int {
		return 1 - m.Get(i, j)
	})
	fmt.Println(m)
	// Output:
	// 1110
	// 1101
	// 1011
	// 0111
}
