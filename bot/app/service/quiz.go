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
//   - ошибка при получении списка тематик из репозитория
func (qs *QuizService) CreateRandomQuiz() (model.Quiz, error) {
	topics, err := qs.QuestionRepo.GetAllTopics()
	if err != nil {
		return model.Quiz{}, err
	}

	n_topics := len(topics)
	index_rand_topic := rand.Intn(n_topics)
	rand_topic := topics[index_rand_topic]

	return qs.CreateQuiz(rand_topic)
}

// Формирование квиза
// Принимает: тематику квиза
// Возвращает: сформированный квиз из 15-ти вопросов по данной тематике и ошибку
func (qs *QuizService) CreateQuiz(topic string) (model.Quiz, error) {
	questions, err := qs.QuestionRepo.FindByTopic(topic)
	if err != nil {
		return model.Quiz{}, err
	}

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	if len(questions) > qs.nQuestsInQuiz {
		questions = questions[:qs.nQuestsInQuiz]
	}

	for _, q := range questions {
		rand.Shuffle(len(q.Answers), func(i, j int) {
			q.Answers[i], q.Answers[j] = q.Answers[j], q.Answers[i]
		})
	}

	return model.Quiz{
		Topic:     topic,
		Questions: questions,
	}, nil
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

func (qs *QuizService) ViewQuestionByID(id int64) (model.Question, error) {
	question, err := qs.QuestionRepo.FindByID(id)
	return question, err
}

func (qs *QuizService) UpdateQuestion(question model.Question) error {
	return qs.QuestionRepo.Update(question)
}
