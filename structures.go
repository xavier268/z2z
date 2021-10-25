package z2z

import (
	"fmt"
	"strings"
)

// bits are packed in uint64 words.

type Vect struct {
	len  int      // length in bits
	data []uint64 // actual data
}

func NewVect(len int) *Vect {
	v := new(Vect)
	v.len = len
	if len > 0 {
		v.data = make([]uint64, 1+((len-1)/64))
	}
	return v
}

func (v *Vect) String() string {
	var b strings.Builder
	for i, d := range v.data {
		if (i+1)*64 <= v.len {
			fmt.Fprintf(&b, "%064b", d)
		} else {
			s := fmt.Sprintf("%064b", d)
			b.WriteString(string(s[64-(v.len%64):]))
		}

	}
	return b.String()
}
