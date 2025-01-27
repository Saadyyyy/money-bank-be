package userservice

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
	"github.com/Saadyyyy/money-bank-be/pkg/user/constant"
	"github.com/Saadyyyy/money-bank-be/utils/healper"
	"github.com/Saadyyyy/money-bank-be/utils/middleware"
	"github.com/jmoiron/sqlx"
)

type UserServiceImpl struct {
	repo domain.UserRepository
	db   *sqlx.DB
}

// RegisterUser implements domain.UserService.
func (u *UserServiceImpl) RegisterUser(ctx context.Context, users domain.Users) (userId int64, err error) {

	if users.Username == "" {
		return 0, fmt.Errorf("Username tidak boleh kosong harus diisi")
	}
	if users.Password == "" {
		return 0, fmt.Errorf(" Password tidak boleh kosong harus diisi")
	}
	users.Username = strings.ToLower(users.Username)

	fmt.Println(users.Username)

	// Cek username yang sudah ada
	if users.Username != "" {
		var existingUsername string
		queryCekUsername := `SELECT username FROM users WHERE username = $1 LIMIT 1`
		err = u.db.QueryRowContext(ctx, queryCekUsername, users.Username).Scan(&existingUsername)
		if err == nil {
			return 0, fmt.Errorf("username sudah digunakan")
		} else if err != sql.ErrNoRows {
			return 0, fmt.Errorf("gagal memeriksa username: %v", err)
		}
	}

	users.Role = constant.StatusUserBiasa

	pass, err := healper.HashPassword(users.Password)
	if err != nil {
		err = fmt.Errorf("failed to hash password")

	}
	users.Password = pass

	userId, err = u.repo.RegisterUser(ctx, users)
	if err != nil {
		return
	}

	return userId, nil
}

func (u *UserServiceImpl) LoginUser(ctx context.Context, username string) (users domain.Users, token string, err error) {
	pass, err := healper.HashPassword(users.Password)
	if err != nil {
		err = fmt.Errorf("gagal hash password")
	}

	users.Password = pass

	users, err = u.repo.LoginUser(ctx, username)
	if err != nil {
		return
	}
	token, err = middleware.CreateToken(users.UserId, users.Role)
	if err != nil {
		return domain.Users{}, "", fmt.Errorf("gagal create token: %+v", err)
	}

	return users, token, nil
}

func NewUserService(repo domain.UserRepository, db *sqlx.DB) domain.UserService {
	return &UserServiceImpl{repo: repo, db: db}
}
