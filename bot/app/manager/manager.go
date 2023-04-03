package manager

import (
	"log"
	"sync"

	"github.com/RullDeef/telegram-quiz-bot/controller"
	"github.com/RullDeef/telegram-quiz-bot/model"
)

type InteractorFactory func(
	*BotManager,
	int64, // chatID
	chan model.Message,
) model.Interactor

type BotManager struct {
	userRepo          model.UserRepository
	quizRepo          model.QuizRepository
	interactorFactory InteractorFactory
	subscriptions     []subscription
	mutex             *sync.RWMutex
}

type subscription struct {
	chatID     int64
	msgChan    chan model.Message
	interactor model.Interactor
}

func NewBotManager(
	interactorFactory InteractorFactory,
	userRepo model.UserRepository,
	quizRepo model.QuizRepository,
) *BotManager {
	return &BotManager{
		userRepo:          userRepo,
		quizRepo:          quizRepo,
		interactorFactory: interactorFactory,
		subscriptions:     nil,
		mutex:             &sync.RWMutex{},
	}
}

func (bm *BotManager) DispatchMessage(msg model.Message) {
	if msg.IsButtonAction {
		log.Printf("[%s] [button %s]", msg.Sender.Nickname, msg.Text)
	} else {
		log.Printf("[%s] %s", msg.Sender.Nickname, msg.Text)
	}

	// check if needs to be broadcasted to interactors
	if bm.tryBroadcast(msg) {
		return // success
	}

	// per-user interaction
	if msg.IsPrivate {
		// TODO: wrap this in a map (maybe?)
		if msg.Text == "/ник" {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				controller.NewUserController(
					bm.userRepo,
					interactor,
				).ChangeNickname()
			})
		} else if msg.Text == "/помощь" {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				controller.NewUserController(
					bm.userRepo,
					interactor,
				).ShowHelp()
			})
		} else if msg.Text == "/создать" {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.quizRepo,
					bm.userRepo,
					interactor,
				).CreateQuiz()
			})
		} else if msg.Text == "/просмотр" {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.quizRepo,
					bm.userRepo,
					interactor,
				).ViewMyQuizzes()
			})
		} else if msg.Text == "/редактировать" {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.quizRepo,
					bm.userRepo,
					interactor,
				).EditQuiz()
			})
		}
	} else { // message came from group chat
		if msg.Text == "/квиз" {
			// start new quiz
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				controller.NewSessionController(
					bm.userRepo,
					bm.quizRepo,
					interactor,
				).Run()
			})
		}
	}
}

func (bm *BotManager) tryBroadcast(msg model.Message) bool {
	bm.mutex.RLock()
	defer bm.mutex.RUnlock()
	for _, sub := range bm.subscriptions {
		if sub.chatID == msg.ChatID {
			sub.msgChan <- msg
			return true
		}
	}
	return false
}

func (bm *BotManager) runJob(chatID int64, job func(model.Interactor)) {
	msgChan := make(chan model.Message)
	interactor := bm.interactorFactory(
		bm,
		chatID,
		msgChan,
	)

	bm.addSubscription(chatID, msgChan, interactor)
	go func(interactor model.Interactor, msgChan chan model.Message) {
		defer close(msgChan)
		defer bm.removeSubscription(chatID)
		job(interactor)
	}(interactor, msgChan)
}

func (bm *BotManager) addSubscription(
	chatID int64,
	msgChan chan model.Message,
	interactor model.Interactor,
) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	bm.subscriptions = append(bm.subscriptions, subscription{
		chatID:     chatID,
		msgChan:    msgChan,
		interactor: interactor,
	})
}

func (bm *BotManager) removeSubscription(chatID int64) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	for i, sub := range bm.subscriptions {
		if sub.chatID == chatID {
			bm.subscriptions = append(bm.subscriptions[:i], bm.subscriptions[i+1:]...)
			break
		}
	}
}
