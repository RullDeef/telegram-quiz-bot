package model

type UserRepository interface {
	Create(User) User
	FindByID(id int64) (User, error)
	FindByTelegramID(id string) (User, error)
	Update(User) error
	Delete(User)
}
