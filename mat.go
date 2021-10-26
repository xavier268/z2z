package z2z

import (
	"fmt"
	"math/bits"
	"math/rand"
	"strings"
)

// Mat is a matrix, its dimesions are l x c.
// Neither l nor c should be 0.
type Mat struct {
	// column bits are packed in a uint64
	// each line uses nbOfWordsPerLine uint64
	// The last uint64 of each line may be partial. In such case, it must be padded with 0.
	// To extract the line l, slice [l * (1 + ((m.c - 1) / 64)) : i < (l+1)*(1+((m.c-1)/64))]
	l, c int      // nb of lines and columns in BITS
	d    []uint64 // actual data
}

const (
	uintmax = 0xFFFF_FFFF_FFFF_FFFF
)

// NewMat constructs a l x c matrix.
func NewMat(l, c int) *Mat {
	m := new(Mat)
	m.l, m.c = l, c
	if l*c > 0 {
		m.d = make([]uint64, l*m.nbOfWordsPerLine())
	}
	return m
}

// NewId constructs Id matrix n x n
func NewId(n int) *Mat {
	r := NewMat(n, n)
	for i := 0; i < n; i++ {
		r.Set(i, i, 1)
	}
	return r
}

// Clone retuns a deep copy of m.
func (m *Mat) Clone() *Mat {
	mm := NewMat(m.l, m.c)
	copy(mm.d, m.d)
	return mm
}

// Dimensions returns (lines , volumns).
func (m *Mat) Dimensions() (int, int) {
	return m.l, m.c
}

// nbOfWordsPerLine provides the nb of uint64 words per line.
func (m *Mat) nbOfWordsPerLine() int {
	return (1 + ((m.c - 1) / 64))
}

// bitCoordinates transform the l,c coordinates of the bit,
// into a word and shift coordinate.
func (m *Mat) bitCoordinates(l, c int) (word, shift int) {
	word = l*m.nbOfWordsPerLine() + c/64
	shift = c % 64
	return word, shift
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

// Set a single coefficient.
// val should be à or 1.
// Exceeding Dimensions is ok, and will be silently ignored.
func (m *Mat) Set(l, c, val int) {
	w, s := m.bitCoordinates(l, c)
	if l >= m.l || c >= m.c {
		return
	}
	switch val {
	case 0:
		m.d[w] &= (uintmax ^ (1 << s))
	case 1:
		m.d[w] |= (1 << s)
	default:
		panic("invalid value input for Set")
	}
}

// Get the corresponding bit.
// Exceeding Dimensions is ok, and will return 0.
func (m *Mat) Get(l, c int) int {
	w, s := m.bitCoordinates(l, c)
	if l >= m.l || c >= m.c {
		return 0
	}
	if m.d[w]&(1<<s) == 0 {
		return 0
	}
	return 1
}

// stringL provides a string representation of the selected line.
func (m *Mat) stringL(l int) string {
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

// String display bits with msb first, as if printed with %b, line by line.
func (m *Mat) String() string {
	if m.c*m.l == 0 {
		return "\n"
	}
	var b strings.Builder
	for l := 0; l < m.l; l++ {
		fmt.Fprintf(&b, "\t%s\n", m.stringL(l))
	}
	return b.String()
}

// Xor is the term to term addition of the Z/2Z group
// m is modified to contain m ^ n
// No checks are made on dimension compatibility.
func (m *Mat) Xor(n *Mat) {
	for i := 0; i < len(m.d) && i < len(n.d); i++ {
		m.d[i] ^= n.d[i]
	}
}

// And is the term to term multiplication in Z/2Z
// m is modified to contain m & n
// No checks are made on dimension compatibility.
func (m *Mat) And(n *Mat) {
	for i := 0; i < len(m.d) && i < len(n.d); i++ {
		m.d[i] &= n.d[i]
	}
}

// Perfom the scalar product in Z/2Z.
// m and n are unchanged.
// Result is 0 or 1.
func (m *Mat) ScalarProduct(n *Mat) int {
	if n == nil {
		return 0
	}
	r := 0
	for i := 0; i < len(m.d) && i < len(n.d); i++ {
		r += bits.OnesCount64(m.d[i] & n.d[i])
	}
	return r & 1
}

// OnesCount return number of non nul coefficients.
func (m *Mat) OnesCount() int {
	r := 0
	for _, d := range m.d {
		r += bits.OnesCount64(d)
	}
	return r
}

// ZerosCount return number of zero coefficients.
func (m *Mat) ZerosCount() int {
	return m.c*m.l - m.OnesCount()
}

// Randomize all coeffficients.
func (m *Mat) Randomize() {
	for i := range m.d {
		m.d[i] = rand.Uint64()
	}
	m.Normalize()
}

// Ones set all coeff to 1
func (m *Mat) Ones() {
	for i := range m.d {
		m.d[i] = uintmax
	}
	m.Normalize()
}

// Zeros set all coeff to 0
func (m *Mat) Zeros() {
	for i := range m.d {
		m.d[i] = 0
	}
	// m.normalize()
}

// Normalize set the padding bits to 0.
func (m *Mat) Normalize() {
	if m.c%64 == 0 {
		return
	}
	wc := m.nbOfWordsPerLine()
	var pad uint64 = uintmax >> (64 - m.c%64)
	for l := 0; l < m.l; l++ {
		m.d[(l+1)*wc-1] &= pad
	}
}

// NewLine adds a new, empty line to the BOTTOM matrix, at the end.
func (m *Mat) NewLine() {
	m.d = append(m.d, make([]uint64, m.nbOfWordsPerLine())...)
	m.l++
}

// NewCol adds a new empty column to the LEFT of the matrix.
func (m *Mat) NewCol() {
	if m.c%64 != 0 {
		m.c++
		return
	}
	// Here,  m.c % 64 = 0
	wc := m.nbOfWordsPerLine()                // before increase
	m.c++                                     // update dims
	wcc := wc + 1                             // after increase
	m.d = append(m.d, make([]uint64, m.l)...) // set the correct global size
	for l := m.l - 1; l > 0; l-- {            // shift  all lines excpt the first one to new position
		copy(m.d[l*wcc:l*wcc+wcc], m.d[l*wc:l*wc+wc])
	}
	for l := 1; l < m.l; l++ { // clear added line content
		m.d[wcc*l-1] = 0
	}

	m.Normalize()
}

// MatMul multiply both matrixes, returning a new one.
// m and n are unchanged.
func (m *Mat) MatMul(n *Mat) *Mat {
	if n == nil || m.c != n.l {
		panic("dimensions are mismatched")
	}
	r := NewMat(m.l, n.c)

	// TODO
	for l := 0; l < m.l; l++ {
		for lc := 0; lc < m.c; lc++ {
			for c := 0; c < n.c; c++ {
				v := r.Get(l, c) ^ (m.Get(l, lc) & n.Get(lc, c))
				r.Set(l, c, v)
			}
		}
	}
	return r
}

// T construct a new transposed matrix from m.
// m is unchanged.
func (m *Mat) T() *Mat {
	r := NewMat(m.c, m.l)
	for l := 0; l < m.l; l++ {
		for c := 0; c < m.c; c++ {
			r.Set(c, l, m.Get(l, c))
		}
	}
	return r
}

func (m *Mat) Equal(n *Mat) bool {
	if n == nil || m.c != n.c || m.l != n.l {
		return false
	}
	m.Normalize() // making sure not to compare non significant bits
	n.Normalize()
	for i := range m.d {
		if m.d[i] != n.d[i] {
			return false
		}
	}
	return true
}
