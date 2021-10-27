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
// If m is invertible, then id is identity and iv is inverse and ok is true
// If m is NOT invertible, id contains identity as a topleft submatrix, and ok is false
// and 0 elsewhere, and iv * m = id
// m is unchanged.
func (m *Mat) Gauss() (id *Mat, iv *Mat, ok bool) {
	ok = (m.c == m.l)
	iv = NewMat(m.l, m.l) // l x l
	for i := 0; i < m.l && i < m.c; i++ {
		iv.Set(i, i, 1)
	}
	id = m.Clone() // l x c , same dim as m

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
			ok = false
			continue
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
	return id, iv, ok
}
