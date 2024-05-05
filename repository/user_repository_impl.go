package repository

import (
	"database/sql"
	"errors"

	"github.com/aribroo/go-ecommerce/entity"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Insert(user *entity.User) error {

	q := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"

	_, err := r.db.Exec(q, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil

}

func (r *UserRepositoryImpl) FindById(id int) (*entity.User, error) {

	q := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id = ?"
	row := r.db.QueryRow(q, id)

	user := new(entity.User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return user, nil

}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {

	q := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?"
	row := r.db.QueryRow(q, email)

	user := new(entity.User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return user, nil

}

func (r *UserRepositoryImpl) FindAll() ([]*entity.User, error) {

	q := "SELECT id, first_name, last_name, email, created_at FROM users"
	rows, err := r.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		var user entity.User

		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil

}

func (r *UserRepositoryImpl) Update(id int, user *entity.UpdateUserPayload) error {

	q := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?"

	_, err := r.db.Exec(q, user.FirstName, user.LastName, id)
	if err != nil {
		return err
	}

	return nil

}

func (r *UserRepositoryImpl) Remove(id int) (int64, error) {

	q := "DELETE FROM users WHERE id = ?"

	result, err := r.db.Exec(q, id)
	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return row, nil

}
