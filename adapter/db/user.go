package db

import (
	"database/sql"

	"github.com/TechLoCo/sample-go-test/model"
	"github.com/TechLoCo/sample-go-test/usecase/repository"

	"github.com/jmoiron/sqlx"
)

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repository.User {
	return &user{
		db: db,
	}
}

func (u *user) Get(id int) (*model.User, error) {
	const query = "SELECT * FROM user WHERE id=?"
	user := model.User{}
	if err := u.db.Select(&user, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
