package square

type Square struct {
	side float64
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}
