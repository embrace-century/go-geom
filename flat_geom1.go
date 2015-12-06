package geom

type geom1 struct {
	geom0
}

func (g *geom1) Coord(i int) []float64 {
	return g.flatCoords[i*g.stride : (i+1)*g.stride]
}

func (g *geom1) Coords() [][]float64 {
	return inflate1(g.flatCoords, 0, len(g.flatCoords), g.stride)
}

func (g *geom1) LastCoord() []float64 {
	if len(g.flatCoords) == 0 {
		return nil
	} else {
		return g.flatCoords[len(g.flatCoords)-g.stride:]
	}
}

func (g *geom1) NumCoords() int {
	return len(g.flatCoords) / g.stride
}

func (g *geom1) PushCoord(coord0 []float64) error {
	if len(coord0) != g.stride {
		return ErrStrideMismatch{Got: len(coord0), Want: g.stride}
	}
	g.flatCoords = append(g.flatCoords, coord0...)
	return nil
}

func (g *geom1) SetCoords(coords1 [][]float64) error {
	var err error
	if g.flatCoords, err = deflate1(nil, coords1, g.stride); err != nil {
		return err
	}
	return nil
}

func (g *geom1) verify() error {
	if g.stride != g.layout.Stride() {
		return errStrideLayoutMismatch
	}
	if len(g.flatCoords)%g.stride != 0 {
		return errLengthStrideMismatch
	}
	return nil
}
