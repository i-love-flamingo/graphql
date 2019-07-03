package infrastructure

import (
	"context"
	"flamingo.me/graphql/example/todo/domain"
	"math/big"
)

type TodoService struct{}

func (ts *TodoService) Todos(ctx context.Context, userid string) ([]*domain.Todo, error) {
	return []*domain.Todo{
		{Task: "a" + userid, Points: big.NewFloat(12.34), Points2: 22.34},
		{Task: "b" + userid},
		{Task: "c" + userid},
	}, nil
}
