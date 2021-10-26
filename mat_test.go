package z2z

import (
	"fmt"
	"testing"
)

var (
	s8  = "00000000"
	s16 = s8 + s8
	s32 = s16 + s16
	s64 = s32 + s32
	s   = "\t"
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
func TestVectString(t *testing.T) {
	type td struct {
		c int    // nb columns
		z string // expected string result
	}
	tdata := []td{
		{0, n},
		{1, s + "0" + n},
		{2, s + "00" + n},
		{32, s + s32 + n},
		{64, s + s64 + n},
		{63, s + s64[1:] + n},
		{65, s + s64 + "0" + n},
	}

	tdata1 := []td{
		{1, s + "1" + n},
		{2, s + "01" + n},
		{64, s + s32 + s16 + s8 + "00000001" + n},
		{63, s + s32 + s16 + s8 + "0000001" + n},
		{65, s + s32 + s16 + s8 + "000000001" + n},
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
	t.Skip()
	for i := 0; i < 64; i++ {
		fmt.Print(NewVectUint64(1 << i).String())
		fmt.Print(NewVectInt(1 << i).String())
	}
}
