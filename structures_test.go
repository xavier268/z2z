package z2z

import (
	"fmt"
	"testing"
)

type td struct {
	int
	string
}

func TestString(t *testing.T) {
	s8 := "00000000"
	s16 := s8 + s8
	s32 := s16 + s16
	s64 := s32 + s32
	tdata := []td{
		{0, ""},
		{1, "0"},
		{2, "00"},
		{32, s32},
		{64, s64},
		{63, s64[1:]},
		{65, s64 + "0"},
	}

	for _, d := range tdata {
		v := NewMat(1, d.int)
		ss := v.String()
		if d.string != ss {
			fmt.Printf("Length   : %d\nExpected : %s\nGot      : %s\n", d.int, d.string, ss)
			t.Fatal()
		}
	}

}
