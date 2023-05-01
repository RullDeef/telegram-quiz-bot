package service

import (
	"fmt"
	"math/rand"
	"strings"

	model "github.com/RullDeef/telegram-quiz-bot/model"
)

const defaultQuestsInQuizCount = 15

type QuizService struct {
	nQuestsInQuiz int
	QuestionRepo  model.QuestionRepository
}

func NewQuizService(
	QuestionRepo model.QuestionRepository,
) *QuizService {
	return &QuizService{
		nQuestsInQuiz: defaultQuestsInQuizCount,
		QuestionRepo:  QuestionRepo,
	}
}

func (qs *QuizService) SetNumQuestionsInQuiz(number int) {
	qs.nQuestsInQuiz = number
}

// Формирование квиза со случайным выбором тематики
//
// Вызывает CreateQuiz()
//
// Выход: все вопросы по выбранной случайным образом тематике, ошибка
//
// Возможные ошибки:
//	 - в случае успеха возвращается nil.
//   - в случае ошибки генерации числа ищутся вопросы с несуществующей тематикой.
func (qs *QuizService) CreateRandomQuiz() (model.Quiz, error) {
	var rand_topic string

	topics := []string{"lisp", "prolog", "python", "Go"}
	n_topics := len(topics)

	index_rand_topic := rand.Intn(n_topics)

	if index_rand_topic >= 0 {
		rand_topic = topics[index_rand_topic]
	} else {
		rand_topic = "#not"
	}

	return qs.CreateQuiz(rand_topic)
}

// Формирование квиза
// Принимает: тематику квиза
// Возвращает: сформированный квиз из 15-ти вопросов по данной тематике и ошибку
func (qs *QuizService) CreateQuiz(topic string) (model.Quiz, error) {
	quiz := model.Quiz{
		Topic: topic,
	}
	var i_quest int
	questions, err := qs.QuestionRepo.FindByTopic(topic)

	if err == nil {
		for i := 0; i < qs.nQuestsInQuiz; i++ {
			i_quest = rand.Intn(qs.nQuestsInQuiz)
			quiz.Questions = append(quiz.Questions, questions[i_quest])
		}
	}

	return quiz, err
}

// Добавление вопроса
// Возвращает идентификатор созданного вопроса
func (qs *QuizService) AddQuestionToTopic(topic string, question string) (int64, error) {
	topic = strings.Trim(topic, " \n\t")
	question = strings.Trim(question, " \n\t")

	if len(topic) == 0 {
		return 0, fmt.Errorf("topic must not be empty")
	}
	if len(question) == 0 {
		return 0, fmt.Errorf("question must not be empty")
	}
	q, err := qs.QuestionRepo.Create(model.Question{
		Text:  question,
		Topic: topic,
	})
	if err != nil {
		return 0, err
	}
	return q.ID, nil
}

// Добавление ответа к вопросу
func (qs *QuizService) AddAnswer(questionID int64, answer string, isCorrect bool) error {
	q, err := qs.QuestionRepo.FindByID(questionID)
	if err != nil {
		return err
	}

	answer = strings.Trim(answer, " \n\t")
	if len(answer) == 0 {
		return fmt.Errorf("question answer must not be empty")
	}

	if isCorrect && q.HasCorrectAnswer() {
		return fmt.Errorf("failed to add question answer: correct answer already exists")
	}

	q.Answers = append(q.Answers, model.Answer{
		Text:      answer,
		IsСorrect: isCorrect,
	})
	return qs.QuestionRepo.Update(q)
}

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
