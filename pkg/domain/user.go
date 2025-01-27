package domain

import "context"

type Users struct {
	UserId    int64
	Username  string
	Password  string
	Role      int64
	UpdatedAt string
	CreatedAt string
}

type LoginUser struct {
	UserId   int64
	Username string
	Token    string
}

type UserRepository interface {
	RegisterUser(ctx context.Context, user Users) (userId int64, err error)
	LoginUser(ctx context.Context, username string) (Users, error)
}

type UserService interface {
	RegisterUser(context.Context, Users) (int64, error)
	LoginUser(ctx context.Context, username string) (user Users, token string, err error)
}
