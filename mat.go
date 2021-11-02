package z2z

import (
	"fmt"
	"math/bits"
	"math/rand"
	"strings"
)

// Mat is a matrix, its dimensions are l_ines x c_olumns.
// Neither l nor c should be 0.
type Mat struct {
	// column bits are packed in one or more uint64
	// each line uses nbOfWordsPerLine uint64
	// The last uint64 of each line may be partial. In such case, it must be padded with 0.
	l, c int      // nb of lines and columns in BITS
	d    []uint64 // actual data
}

const (
	uintmax = 0xFFFF_FFFF_FFFF_FFFF
	Version = "0.3.3"
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

// NewId constructs the identity matrix n x n, diagonal is set to 1.
func NewId(n int) *Mat {
	return NewProjector(n, n)
}

// NewProjector construct a Projector (ie, diagonal set to 1) of dimension l x c
func NewProjector(l, c int) *Mat {
	r := NewMat(l, c)
	for i := 0; i < l && i < c; i++ {
		r.Set(i, i, 1)
	}
	return r
}

// Clone returns a deep copy of m.
func (m *Mat) Clone() *Mat {
	mm := NewMat(m.l, m.c)
	copy(mm.d, m.d)
	return mm
}

// Dimensions of m : (lines , columns).
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
// Note : Vect are defined as "horizontal" matrix, for performance.
func NewVect(c int) *Mat {
	return NewMat(1, c)
}

// NewVectUint64 constructs a Vect from the bits of a uint64.
// The size of the Vect is adjusted to fit all the non zero bits.
func NewVectUint64(value uint64) *Mat {
	v := NewVect(bits.Len(uint(value)))
	v.d[0] = value
	return v
}

// NewVectInt constructs a Vect from the bits of a int.
// The size of the Vect is adjusted to fit all the non zero bits.
func NewVectInt(value int) *Mat {
	v := NewVect(bits.Len(uint(value)))
	v.d[0] = uint64(value)
	return v
}

// Set a single coefficient.
// val should be 0 or 1. Attempting to set another value will panic.
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

// Get the corresponding bit (0 or 1).
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

// String display bits with the most-significant-bit first, as if printed with %b, line by line.
// Note that doing so, the main diagonal will be displayed from top right to bottom left.
func (m *Mat) String() string {
	if m.c*m.l == 0 {
		return "\n"
	}
	var b strings.Builder
	for l := 0; l < m.l; l++ {
		fmt.Fprintf(&b, "%s\n", m.stringL(l))
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

// Randomize all coefficients using pseudo-random default generator.
func (m *Mat) Randomize() {
	for i := range m.d {
		m.d[i] = rand.Uint64()
	}
	m.Normalize()
}

// Ones set all coeff to 1
// Optimized for efficiency.
func (m *Mat) Ones() {
	for i := range m.d {
		m.d[i] = uintmax
	}
	m.Normalize()
}

// Zeros set all coeff to 0
// Optimized for efficiency.
func (m *Mat) Zeros() {
	for i := range m.d {
		m.d[i] = 0
	}
	// m.normalize()
}

// Normalize set the internal padding bits to 0.
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

// NewLine adds a new, empty line to the BOTTOM of the matrix, at the end.
func (m *Mat) NewLine() {
	m.d = append(m.d, make([]uint64, m.nbOfWordsPerLine())...)
	m.l++
}

// NewCol adds a new empty column to the LEFT of the matrix ( higher columns index)
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

// MatMul multiply both matrixes, returning a m x n in a new one.
// Naive implementation.
// m and n are unchanged.
func (m *Mat) matMulNaive(n *Mat) *Mat {
	if n == nil || m.c != n.l {
		panic("dimensions are mismatched")
	}
	r := NewMat(m.l, n.c)

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

// MatMul multiply both matrixes, returning a m x n in a new one.
// m and n are unchanged.
func (m *Mat) MatMul(n *Mat) *Mat {
	return m.MatMulTr(n.T())
	//return m.matMulNaive(n)
}

// Multiply m by the transposed of p.
// m and p are unchanged.
// More efficient because word by word operations process 64 elementary operations at once.
func (m *Mat) MatMulTr(p *Mat) *Mat {
	if p == nil || m.c != p.c {
		panic("dimensions are mismatched")
	}
	r := NewMat(m.l, p.l)

	w := m.nbOfWordsPerLine()
	for i := 0; i < m.l; i++ {
		for j := 0; j < p.l; j++ {
			var v uint64 = 0
			for c := 0; c < w; c++ {
				v ^= m.d[i*w+c] & p.d[j*w+c]
			}
			r.Set(i, j, bits.OnesCount64(v)&1)
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

// Equal test for equality, ie same dimensions and same content.
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

// NewFromInt constructs a matrix from the provided uint64_s.
// Each value represent a line.
// There are always 64 columns.
// Use CloneDims to adjust dimensions afterwards if needed.
func NewFromInt(ss ...uint64) *Mat {
	m := NewMat(len(ss), 64)
	for d := range m.d {
		m.d[d] = ss[d]
	}
	return m
}

// CloneDims clone m into a new matrix with different dimensions.
// m is unchanged, n is 0-padded if necessary.
// Specify a 0 or negative value to keep the original dimension.
func (m *Mat) CloneDims(l, c int) (n *Mat) {
	if l <= 0 {
		l = m.l
	}
	if c <= 0 {
		c = m.c
	}
	n = NewMat(l, c)
	for i := 0; i < m.l && i < n.l; i++ {
		for j := 0; j < m.c && j < n.c; j++ {
			n.Set(i, j, m.Get(i, j))
		}
	}
	return n
}

// Apply a function to the matrix that takes coordinates as input,
// generating the value to Set (0 or 1)
func (m *Mat) Apply(f func(i, j int) (value int)) {
	for i := 0; i < m.l; i++ {
		for j := 0; j < m.c; j++ {
			m.Set(i, j, f(i, j))
		}
	}
}
