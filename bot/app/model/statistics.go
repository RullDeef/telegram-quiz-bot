package model

// Модельная сущность статистики пользователя
type Statistics struct {
	// Идентификатор пользователя
	UserID                int64

	// Количество пройденных квизов (успешно и неуспешно)
	QuizzesCompleted      uint

	// Среднее время прохождения квиза в секундах
	MeanQuizCompleteTime  float64

	// Среднее время ответа в секундах
	MeanQuestionReplyTime float64

	// Общее количество ответов
	TotalReplies          uint

	// Количество верных ответов
	CorrectReplies        uint

	// Процент верных ответов (от 0.0 до 1.0)
	CorrectRepliesPercent float64
}
