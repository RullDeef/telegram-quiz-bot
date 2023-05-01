package model

type UserService interface {
	CreateUser(username string, telegramId string) (User, error)
	SetUserRole(role string, telegramId string) bool
	GetUserByTelegramId(id string) (User, error)
	ChangeUsername(username string, telegramId string) bool
}
