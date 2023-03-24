package model

type UserRepository interface {
	Create(User) bool
	FindByID(id uint64) User
	FindByTelegramID(id string) User
	Update(User)
	Delete(User)
}
