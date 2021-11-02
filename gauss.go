package z2z

// swap lines i and j
func (m *Mat) swapLines(i, j int) {
	if i == j {
		return
	}
	wc := m.nbOfWordsPerLine()
	for w := 0; w < wc; w++ {
		m.d[i*wc+w], m.d[j*wc+w] = m.d[j*wc+w], m.d[i*wc+w]
	}
}

// swap columns i and j
// less effient that swapping lines.
func (m *Mat) swapCols(i, j int) {
	if i == j {
		return
	}
	for l := 0; l < m.l; l++ {
		v, w := m.Get(l, i), m.Get(l, j)
		m.Set(l, i, w)
		m.Set(l, j, v)
	}
}

// add line j to line i
func (m *Mat) addLines(i, j int) {
	wc := m.nbOfWordsPerLine()
	for w := 0; w < wc; w++ {
		m.d[i*wc+w] ^= m.d[j*wc+w]
	}
}

// add col j to i
func (m *Mat) addCols(i, j int) {
	for l := 0; l < m.l; l++ {
		m.Set(l, i, m.Get(l, j)^m.Get(l, i))
	}
}

// Gauss apply the Gauss-Jordan pivot to compute inverse of m.
// If m is NOT invertible,  iv is nil
// m is unchanged.

// gaussShort returns as soon as it is clear that m is not invertible.
func (m *Mat) GaussShort() (iv *Mat) {

	if m.c != m.l {
		return nil
	}

	iv = NewMat(m.l, m.l) // l x l
	for i := 0; i < m.l && i < m.c; i++ {
		iv.Set(i, i, 1)
	}
	id := m.Clone() // l x c , same dim as m

	for r := 0; r < m.l && r < m.c; r++ {
		// ensure (r,r) is 1
		if id.Get(r, r) == 0 {
			// look for a line under to swap ?
			for l := r + 1; l < id.l; l++ {
				if id.Get(l, r) == 1 {
					id.swapLines(r, l)
					iv.swapLines(r, l)
					break
				}
			}
		}

		// no way to get a 1, lets continue, not invertible
		if id.Get(r, r) == 0 {
			return nil
		} else {
			// we have a 1 at (r,r),
			// lets clean up the rest of the lines
			for l := 0; l < id.l; l++ {
				if l != r && id.Get(l, r) == 1 {
					id.addLines(l, r)
					iv.addLines(l, r)
				}
			}
		}
	}
	return iv
}

// GaussFull computes a full row echelon format (re) and the quasi inverse (iv).
// ok is true if m was inversible, and in such case, re is identity.
// iv(l,l) * m (l,c) = re (l,c)
// If ok is true, determinant is 1, it is 0 otherwise.
// The rk is the rank of the matrix.
func (m *Mat) GaussFull() (re *Mat, iv *Mat, ok bool, rk int) {
	ok = (m.c == m.l)
	iv = NewId(m.l) // l x l, Id or projector in a sub-space.
	re = m.Clone()  // l x c , same dim as m

	rl, rc := 0, 0 // pivots indexes
	for rl < m.l && rc < m.c {
		// ensure (rl,rc) is 1
		if re.Get(rl, rc) == 0 {
			// look for a line under to swap ?
			for l := rl + 1; l < re.l; l++ {
				if re.Get(l, rc) == 1 {
					re.swapLines(rl, l)
					iv.swapLines(rl, l)
					break
				}
			}
		}

		// no way to get a 1, lets continue, not invertible
		if re.Get(rl, rc) == 0 {
			ok = false
			rc++
			continue
		} else {
			// we have a 1 at (rl,rc),
			// lets clean up the rest of the lines
			for l := 0; l < re.l; l++ {
				if l != rl && re.Get(l, rc) == 1 {
					re.addLines(l, rl)
					iv.addLines(l, rl)
				}
			}
			rc++
			rl++
			rk++
		}
	}
	return re, iv, ok, rk
}
