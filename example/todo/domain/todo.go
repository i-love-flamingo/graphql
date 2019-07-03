package domain

import "math/big"

type Todo struct {
	ID      string
	Task    string
	Points  *big.Float
	Points2 float64
}

type Task interface {
	A() string
}

func (*Todo) A() string {
	return "a"
}
