package user

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	DataBase *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DataBase: db,
	}
}
