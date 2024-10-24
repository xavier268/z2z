package z2z

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	s8  = "00000000"
	s16 = s8 + s8
	s32 = s16 + s16
	s64 = s32 + s32
	n   = "\n"
)

func TestNbWordsPerLine(t *testing.T) {

	type td struct {
		l  int
		c  int
		nb int
	}
	tdata := []td{
		{0, 0, 1},
		{1, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
		{0, 0, 1},
		{1, 63, 1},
		{1, 64, 1},
		{1, 65, 2},
		{2, 63, 1},
		{2, 64, 1},
		{2, 65, 2},
		{0, 63, 1},
		{0, 64, 1},
		{0, 65, 2},
		{2, 127, 2},
		{2, 128, 2},
		{2, 129, 3},
	}

	for i, tt := range tdata {
		m := NewMat(tt.l, tt.c)
		if m.nbOfWordsPerLine() != tt.nb {
			t.Fatalf("invalid nbOfWords, tdata index %d, want : %d, got %d", i, tt.nb, m.nbOfWordsPerLine())
		}
	}

}

func TestBitCoord(t *testing.T) {
	type td struct {
		l     int
		c     int
		word  int
		shift int
	}
	tdata := []td{
		{0, 0, 0, 0},
		{3, 65, 7, 1},
		{3, 64, 7, 0},
		{3, 63, 6, 63},
		{0, 7, 0, 7},
	}

	m := NewMat(5, 80)

	for i, td := range tdata {
		w, s := m.bitCoordinates(td.l, td.c)
		if w != td.word || s != td.shift {
			fmt.Printf("tdata : %d, want : %d, %d\t got : %d, %d", i, td.word, td.shift, w, s)
			t.Fatal("bit coordinates do not match")
		}
	}
}
func TestVectString(t *testing.T) {
	type td struct {
		c int    // nb columns
		z string // expected string result
	}
	tdata := []td{
		{0, n},
		{1, "0" + n},
		{2, "00" + n},
		{32, s32 + n},
		{64, s64 + n},
		{63, s64[1:] + n},
		{65, s64 + "0" + n},
	}

	tdata1 := []td{
		{1, "1" + n},
		{2, "01" + n},
		{64, s32 + s16 + s8 + "00000001" + n},
		{63, s32 + s16 + s8 + "0000001" + n},
		{65, s32 + s16 + s8 + "000000001" + n},
		{66, s32 + s16 + s8 + "0000000001" + n},
	}

	// test zeros
	for _, d := range tdata {
		v := NewVect(d.c)
		ss := v.String()
		if d.z != ss {
			fmt.Printf("test zeros - length   : %d\nExpected : %s\nGot      : %s\n", d.c, d.z, ss)
			t.Fatal()
		}
	}

	// test with one
	for _, d := range tdata1 {
		v := NewVect(d.c)
		v.d[0] = 1
		ss := v.String()
		if d.z != ss {
			fmt.Printf("test ones - length   : %d\nExpected : %s\nGot      : %s\n", d.c, d.z, ss)
			t.Fatal()
		}
	}
}

func TestString(t *testing.T) {
	//t.Skip()
	for i := 0; i < 64; i++ {
		//fmt.Print(NewVectUint64(1 << i).String())
		fmt.Print(NewVectInt(1 << i).String())
	}
}

func TestNormalize(t *testing.T) {
	m := NewMat(3, 5)
	m.c = 64
	m.Ones()
	fmt.Println(m)
	m.c = 5
	m.Normalize()
	m.c = 64
	fmt.Println(m)
}

func TestSetGet(t *testing.T) {
	m := NewMat(3, 63*3)
	m.Randomize()
	for i := 0; i < 100; i++ {
		l, c := rand.Intn(m.l), rand.Intn(m.c)
		m.Set(l, c, 0)
		v := m.Get(l, c)
		if v != 0 {
			t.Fatal("should read 0")
		}
		m.Set(l, c, 1)
		v = m.Get(l, c)
		if v != 1 {
			t.Fatal("should read 1")
		}
		v = rand.Intn(2)
		m.Set(l, c, v)
		if m.Get(l, c) != v {
			t.Fatalf("Random write/read failed")
		}
	}
	m.Ones()
	for i := 0; i < 100; i++ {
		l, c := rand.Intn(m.l), rand.Intn(m.c)
		if m.Get(l, c) != 1 {
			t.Fatal("should read 1")
		}
	}
	m.Zeros()
	for i := 0; i < 100; i++ {
		l, c := rand.Intn(m.l), rand.Intn(m.c)
		if m.Get(l, c) != 0 {
			t.Fatal("should read 0")
		}
	}
}

func TestNewLine(t *testing.T) {
	m := NewMat(4, 5)
	m.Ones()
	fmt.Println(m)
	for l := 0; l < 4; l++ {
		m.NewLine()
		fmt.Println(m)
	}
}

func TestNewCol1(t *testing.T) {
	m := NewMat(4, 5)
	m.Ones()
	fmt.Println(m)
	for l := 0; l < 4; l++ {
		m.NewCol()
		fmt.Println(m)
	}
}

func TestNewCol2(t *testing.T) {
	checkAddedCols(t, 10, 3, 5)
	checkAddedCols(t, 10, 63, 5)
	checkAddedCols(t, 10, 2*63, 6)
	checkAddedCols(t, 10, 3*63, 7)
	checkAddedCols(t, 10, 4*63, 8)
	checkAddedCols(t, 1, 4*63, 8)
	checkAddedCols(t, 10, 2, 300)

}

func checkAddedCols(t *testing.T, l int, c int, adds int) {
	_ = adds // compiler happy !
	m := NewMat(l, c)
	m.Randomize()
	ref := m.Clone()
	//fmt.Println(m)
	for l := 0; l < 6; l++ {
		m.NewCol()
		//fmt.Println(m)
	}
	//Check result
	for l := 0; l < ref.l; l++ {
		for c := 0; c < ref.c; c++ {
			if ref.Get(l, c) != m.Get(l, c) {
				fmt.Println("dims", ref.l, ref.c, "-->", m.l, m.c)
				fmt.Println(ref)
				fmt.Println(m)
				t.Fatalf("Sub matrix does not match after adding cols at l=%d, c=%d", l, c)
			}
		}
		for c := ref.c; c < m.c; c++ {
			if m.Get(l, c) != 0 {
				fmt.Println("dims", ref.l, ref.c, "-->", m.l, m.c)
				fmt.Println(ref)
				fmt.Println(m)
				t.Fatal("Added columns should be zero")
			}
		}
	}
}

func TestTranspose(t *testing.T) {
	var m, tm, ttm *Mat

	m = NewId(200)
	tm = m.T()
	if !m.Equal(tm) {
		t.Fatal("ID should self transpose")
	}

	m = NewMat(3, 7)
	tm = m.T()
	if m.Equal(tm) {
		t.Fatal("unexpected equal transpose")
	}
	ttm = tm.T()
	if !ttm.Equal(m) {
		t.Fatal("Double transpose should be idempotent")
	}

	m = NewMat(80, 7)
	tm = m.T()
	ttm = tm.T()
	if !ttm.Equal(m) {
		t.Fatal("Double transpose should be idempotent")
	}

	m = NewMat(230, 7)
	tm = m.T()
	ttm = tm.T()
	if !ttm.Equal(m) {
		t.Fatal("Double transpose should be idempotent")
	}

	m = NewMat(3, 277)
	tm = m.T()
	ttm = tm.T()
	if !ttm.Equal(m) {
		t.Fatal("Double transpose should be idempotent")
	}

}

func TestMatMul(t *testing.T) {

	if NewId(127).OnesCount() != 127 {
		t.Fatal("error constructing Id, wrong nb of bits set")
	}

	m := NewMat(5, 10)
	m.Randomize()
	n := NewMat(10, 3)
	n.Randomize()

	p := m.MatMul(n)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)

	l, c := p.Dimensions()
	if l != m.l || c != n.c {
		t.Fatal("mismatched dimensions")
	}

	tm := m.T()
	tn := n.T()
	tp := p.T()

	if !tp.Equal(tn.MatMul(tm)) {
		t.Fatal("Transposed multiplicatio does not work")
	}

	if !m.MatMul(NewId(10)).Equal(m) {
		t.Fatal("multiplication by ID failed")
	}
	if !NewId(5).MatMul(m).Equal(m) {
		t.Fatal("multiplication by ID failed")
	}

}

func TestNewFromBytes(t *testing.T) {
	type td struct {
		bs []byte
		st string
	}
	data := []td{
		{[]byte{0}, "00000000"},
		{[]byte{1}, "00000001"},
		{[]byte{128}, "10000000"},
		{[]byte{0, 0}, "0000000000000000"},
		{[]byte{0, 1}, "0000000100000000"},
		{[]byte{1, 0}, "0000000000000001"},
		{[]byte{1, 2}, "0000001000000001"},
		{[]byte{0, 3, 0}, "000000000000001100000000"},
		{[]byte{128, 3, 0}, "000000000000001110000000"},
		{[]byte{0, 3, 128}, "100000000000001100000000"},
	}

	for i, d := range data {
		v := NewFromBytes(d.bs)
		if v.l != 1 || v.c != len(d.bs)*8 {
			t.Fatal("dims are mismatched")
		}
		if v.stringL(0) != d.st {
			fmt.Printf("%d)  : bytes %b\ngot : %s\nwant: %s\n", i, d.bs, v, d.st)
			t.Fatal("incorrect result")
		}
	}
}
