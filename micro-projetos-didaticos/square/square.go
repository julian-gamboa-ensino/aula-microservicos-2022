package square

type Square struct {
	Side float64
}

func (sq Square) Area() float64 {
	return sq.Side * sq.Side
}
