package handler

import (
	"context"
	"fmt"

	"github.com/0daryo/falcon/pb/server"
)

type User struct {
}

func (s *User) Get(ctx context.Context, r *server.GetUserRequest) (*server.User, error) {
	fmt.Println(r.GetId())
	return &server.User{
		Id:   r.GetId(),
		Name: "dummy",
		Age:  18,
	}, nil
}
