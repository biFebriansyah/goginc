package repositories

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/models"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type RepoUserIF interface {
	CreateUser(data *models.User) (*config.Result, error)
	GetAllUser() (*config.Result, error)
	GetAuthData(user string) (*models.User, error)
}

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (*config.Result, error) {
	q := `INSERT INTO tiketz.user(
		username,
		password,
		role
	) VALUES(
		:username, 
		:password,
		:role
	)
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil
}

func (r *RepoUser) GetAllUser() (*config.Result, error) {
	var data models.Users
	q := `SELECT username, "role", created_at, updated_at FROM tiketz."user" ORDER BY created_at DESC`

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &config.Result{Data: data}, nil
}

func (r *RepoUser) GetAuthData(user string) (*models.User, error) {
	var result models.User
	q := `SELECT user_id, username, "role", "password" FROM tiketz."user" WHERE username = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}
