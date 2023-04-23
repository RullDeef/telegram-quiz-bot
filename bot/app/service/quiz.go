package service

import (
	"fmt"
	"math/rand"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

const defaultQuestsInQuizCount = 15

type QuizService struct {
	nQuestsInQuiz int
	QuestionRepo  model.QuestionRepository
}

func NewQuizService(
	//QuizRepo model.QuizRepository,
	QuestionRepo model.QuestionRepository,
) *QuizService {
	return &QuizService{
		nQuestsInQuiz: defaultQuestsInQuizCount,
		QuestionRepo:  QuestionRepo,
	}
}

func (qs *QuizService) setNumQuestionsInQuiz(number int) {
	qs.nQuestsInQuiz = number
}

// Формирование квиза
// Принимает: тематику квиза
// Возвращает: сформированный квиз из 15-ти вопросов по данной тематике и ошибку
func (qs *QuizService) CreateQuiz(topic string) (model.Quiz, error) {
	var quiz model.Quiz
	var i_quest int
	questions, err := qs.QuestionRepo.FindByTopic(topic)

	//n_questions := len(questions)
	//qs.nQuestsInQuiz = 15 // ??

	for i := 0; i < qs.nQuestsInQuiz; i++ {
		i_quest = rand.Intn(qs.nQuestsInQuiz)
		quiz.Questions = append(quiz.Questions, questions[i_quest])
	}

	return quiz, err
}

// Добавление вопроса
// Возвращает идентификатор созданного вопроса
func (qs *QuizService) AddQuestionToTopic(topic string, question string) (int64, error) {
	q := model.Question{
		Text:  question,
		Topic: topic,
	}
	q, err := qs.QuestionRepo.Create(q)
	if err != nil {
		return 0, err
	}
	return q.ID, nil
}

// Добавление ответа к вопросу
/*func (qs *QuizService) AddAnswer(questionID int64, answer string, isCorrect bool) error {
	q, err := qs.QuestionRepo.FindByID(questionID)
	if err != nil {
		return err
	}

	// валидировать строку answer (если она пустая, то ошибка)

	// проверить, если правильный ответ уже есть и isCorrect = true -> ошибка

	// если никаких ошибок - добавить ответ в квесчон и апдейтнуть его в репозитории
	// q.Answers = append...
	return qs.QuestionRepo.Update(q)
}*/

// Просмотр вопросов одной тематики
func (qs *QuizService) ViewQuestionsByTopic(topic string) ([]string, error) {
	questions, err := qs.QuestionRepo.FindByTopic(topic)
	if err != nil {
		return nil, err
	}
	var descriptions []string
	for _, q := range questions {
		description := fmt.Sprintf("%d. %s", q.ID, q.Text)
		descriptions = append(descriptions, description)
	}
	return descriptions, nil
}
