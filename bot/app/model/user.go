package model

// Ограничение базы данных для поля Роли
const (
  UserRoleUser = "USER";
  UserRoleAdmin = "ADMIN"
)

// Модельная сущность Пользователь
type User struct {
	// ID пользователя
	ID int64

	// Имя пользователя
	Nickname string

	// Telegram ID пользователя
	TelegramID string

	// Роль пользователя (ADMIN, USER)
	Role string
}

func NewUser(ID int64, Nickname, TelegramID, Role string) User {
	return User{ID, Nickname, TelegramID, Role}
}
