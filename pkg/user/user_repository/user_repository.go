package userrepository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
	"github.com/jmoiron/sqlx"
)

const (
	queryCreateUser = `
		INSERT INTO users(username,password,role,created_at)VALUES($1,$2,$3,$4) returning user_id;
	`

	queryLoginUser = `
		select user_id,username,password,role,created_at,updated_at from users where username = $1 and deleted_at is null;
	`
)

type UserRepositoryInterfaceImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepositoryInterfaceImpl{db: db}
}

// Register for users
func (u *UserRepositoryInterfaceImpl) RegisterUser(ctx context.Context, user domain.Users) (userId int64, err error) {
	createdAt := time.Now()
	err = u.db.QueryRowContext(ctx, queryCreateUser, user.Username, user.Password, user.Role, createdAt).Scan(&userId)
	if err != nil {
		err = fmt.Errorf("query InsertUser error: %+v", err)
		return
	}
	return userId, nil
}

func (u *UserRepositoryInterfaceImpl) LoginUser(ctx context.Context, username string) (domain.Users, error) {
	row := u.db.QueryRowContext(ctx, queryLoginUser, username)
	user, err := ScanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no user found with username: %s", username)
		}
		return user, fmt.Errorf("error while scanning user: %v", err)
	}

	return user, nil
}

func ScanUser(row *sql.Row) (domain.Users, error) {
	var user domain.Users
	var createdAt, updatedAt sql.NullString

	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Role,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return user, fmt.Errorf("failed to scan user: %v", err)
	}

	return user, nil
}
