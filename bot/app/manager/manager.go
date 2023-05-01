package manager

import (
	"sync"

	"github.com/RullDeef/telegram-quiz-bot/controller"
	"github.com/RullDeef/telegram-quiz-bot/model"
	"github.com/RullDeef/telegram-quiz-bot/service"
	log "github.com/sirupsen/logrus"
)

type InteractorFactory func(
	*BotManager,
	int64, // chatID
	chan model.Message,
) model.Interactor

type BotManager struct {
	interactorFactory InteractorFactory
	userService       *service.UserService
	statService       *service.StatisticsService
	quizService       *service.QuizService
	logger            *log.Logger
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
	userService *service.UserService,
	statService *service.StatisticsService,
	quizService *service.QuizService,
	logger *log.Logger,
) *BotManager {
	return &BotManager{
		interactorFactory: interactorFactory,
		userService:       userService,
		statService:       statService,
		quizService:       quizService,
		logger:            logger,
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
		if msg.Text == commandRegister {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				bm.newUserController(interactor).Register(*msg.Sender)
			})
		} else if msg.Text == commandChangeNickname {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				bm.newUserController(interactor).ChangeNickname()
			})
		} else if msg.Text == commandHelp {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				bm.newUserController(interactor).ShowHelp()
			})
		} else if msg.Text == commandCreateQuiz {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.userService,
					interactor,
				).CreateQuiz()
			})
		} else if msg.Text == commandViewQuizzes {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.userService,
					interactor,
				).ViewMyQuizzes()
			})
		} else if msg.Text == commandEditQuiz {
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				// TODO: check sender role here
				controller.NewAdminController(
					bm.userService,
					interactor,
				).EditQuiz()
			})
		}
	} else { // message came from group chat
		if msg.Text == commandStartQuiz {
			// start new quiz
			bm.runJob(msg.ChatID, func(interactor model.Interactor) {
				controller.NewSessionController(
					bm.userService,
					bm.statService,
					bm.quizService,
					interactor,
					bm.logger,
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

func (bm *BotManager) newUserController(interactor model.Interactor) *controller.UserController {
	return controller.NewUserController(
		bm.userService,
		bm.statService,
		interactor,
		bm.logger,
	)
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
