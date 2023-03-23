package model

type UserRepository interface {
	Create(User)
	FindByID(id uint64)
	FindByTelegramID(id string)
	Update(User)
	Delete(User)
}
