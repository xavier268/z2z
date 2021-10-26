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
