package z2z

import (
	"fmt"
	"strings"
)

// Mat is a matrix, its dimesions are l x c.
// Neither l nor c should be 0.
type Mat struct {
	l, c int      // nb of lines, columns
	data []uint64 // actual data
}

// NewMat constructs a l x c matrix.
func NewMat(l, c int) *Mat {
	v := new(Mat)
	v.l, v.c = l, c
	if l*c > 0 {
		v.data = make([]uint64, l*(1+((c-1)/64)))
	}
	return v
}

// Dimensions returns (lines , volumns).
func (m *Mat) Dimensions() (int, int) {
	return m.l, m.c
}

// NewVect gives a 1 x c matrix.
func NewVect(c int) *Mat {
	return NewMat(1, c)
}

// StringL provides a string representation of the selected line.
func (m *Mat) StringL(l int) string {
	var b strings.Builder
	for i, d := range m.data[l:] {
		if (i+1)*64 <= m.l {
			fmt.Fprintf(&b, "%064b", d)
		} else {
			s := fmt.Sprintf("%064b", d)
			b.WriteString(string(s[64-(m.l%64):]))
		}
	}
	return b.String()
}
