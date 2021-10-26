package z2z

import (
	"fmt"
	"math/bits"
	"strings"
)

// Mat is a matrix, its dimesions are l x c.
// Neither l nor c should be 0.
type Mat struct {
	// bits ares packed in uint64
	// each line uses (1+((c-1)/64))) uint64
	// The last uint64 of each line may be partial.
	// To extract the line l, slice [l * (1 + ((m.c - 1) / 64)) : i < (l+1)*(1+((m.c-1)/64))]
	l, c int      // nb of lines, columns in bits
	d    []uint64 // actual data
}

// NewMat constructs a l x c matrix.
func NewMat(l, c int) *Mat {
	m := new(Mat)
	m.l, m.c = l, c
	if l*c > 0 {
		m.d = make([]uint64, l*m.nbOfWordsPerLine())
	}
	return m
}

// Dimensions returns (lines , volumns).
func (m *Mat) Dimensions() (int, int) {
	return m.l, m.c
}

// nbOfWordsPerLine provides the nb of uint64 words per line.
func (m *Mat) nbOfWordsPerLine() int {
	return (1 + ((m.c - 1) / 64))
}

// NewVect retuns a 1 x c matrix.
func NewVect(c int) *Mat {
	return NewMat(1, c)
}

func NewVectUint64(value uint64) *Mat {
	v := NewVect(bits.Len(uint(value)))
	v.d[0] = value
	return v
}

func NewVectInt(value int) *Mat {
	v := NewVect(bits.Len(uint(value)))
	v.d[0] = uint64(value)
	return v
}

// StringL provides a string representation of the selected line.
func (m *Mat) StringL(l int) string {
	var b strings.Builder
	first := true
	for i := (l+1)*m.nbOfWordsPerLine() - 1; i >= l*m.nbOfWordsPerLine(); i-- {
		d := m.d[i]
		if first && m.c%64 != 0 {
			// for the first word, keep only the m.c % 64 least significant bits.
			s := fmt.Sprintf("%064b", d)
			b.WriteString(string(s[64-(m.c%64):]))
		} else {
			fmt.Fprintf(&b, "%064b", d)
		}
		first = false
	}
	return b.String()
}

// String display bits from left to right, line by line.
func (m *Mat) String() string {
	if m.c*m.l == 0 {
		return "\n"
	}
	var b strings.Builder
	for l := 0; l < m.l; l++ {
		fmt.Fprintf(&b, "\t%s\n", m.StringL(l))
	}
	return b.String()
}
