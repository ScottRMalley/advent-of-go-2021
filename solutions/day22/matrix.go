package day22

type SparseMat struct {
	xMax, yMax, zMax uint64
	data             map[uint64]uint8
}

func NewSparseMat(xMax, yMax, zMax int) *SparseMat {
	return &SparseMat{
		xMax: uint64(xMax),
		yMax: uint64(yMax),
		zMax: uint64(zMax),
		data: make(map[uint64]uint8),
	}
}

func (m *SparseMat) to1D(x, y, z int) uint64 {
	return (uint64(z) * m.xMax * m.yMax) + (uint64(y) * m.xMax) + uint64(x)
}

func (m *SparseMat) On(x, y, z int) {
	m.data[m.to1D(x, y, z)] = uint8(1)
}

func (m *SparseMat) Off(x, y, z int) {
	delete(m.data, m.to1D(x, y, z))
}

func (m *SparseMat) Len() int {
	return len(m.data)
}
