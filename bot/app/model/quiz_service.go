package model

type QuizService interface {
	SetNumQuestionsInQuiz(number int)
	CreateQuiz(topic string) (Quiz, error)
	CreateRandomQuiz() (Quiz, error)
	AddQuestionToTopic(topic string, question string) (int64, error)
	AddAnswer(questionID int64, answer string, isCorrect bool) error
	ViewQuestionsByTopic(topic string) ([]string, error)
}
