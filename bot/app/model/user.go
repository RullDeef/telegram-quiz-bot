package model

type User struct {
	ID         int64
	Nickname   string
	TelegramID string
	Role       string
}

func NewUser(ID int64, Nickname, TelegramID, Role string) *User {
	return &User{ID, Nickname, TelegramID, Role}
}
