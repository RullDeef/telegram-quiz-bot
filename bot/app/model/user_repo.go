package model

type UserRepository interface {
	Create(User) User
	FindByID(id int64) (User, error)
	FindByTelegramID(id string) (User, error)
	Update(User) error
	Delete(User)
}

//New version
type UserRepositoryNew interface {
	Create(User) error
	FindByID(id int64) (User, error)
	FindByTelegramID(id string) (User, error)
	Update(User) error
	Delete(User) error
}
