package repository

import (
	"github.com/aribroo/go-ecommerce/entity"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindById(id int) (*entity.User, error)
	Insert(user *entity.User) error
	FindAll() ([]*entity.User, error)
	Update(id int, user *entity.UpdateUserPayload) error
	Remove(id int) (int64, error)
}
