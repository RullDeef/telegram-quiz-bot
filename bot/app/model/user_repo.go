package model

// Репозиторий для работы с пользователями
type UserRepository interface {

	// Создает пользователя
	Create(User) (User, error)

	// Производит поиск пользователя по идентифиатору
	FindByID(id int64) (User, error)

	// Производит поиск пользователя по тэгу в Телеграме
	FindByTelegramID(id string) (User, error)

	// Обновляет пользователя
	Update(User) error

	// Удаляет пользователя
	Delete(User) error
}
