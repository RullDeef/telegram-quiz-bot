package mockinteractor

import (
	"testing"
	"time"

	"github.com/RullDeef/telegram-quiz-bot/model"
)

type responseMatch struct {
	expectation string
	response    model.Response
}

type MockInteractor struct {
	msgChan   chan model.Message
	responses chan model.Response

	expectations []string

	unsentMessages     []model.Message
	uncaughtResponses  []model.Response
	unmatchedResponses []responseMatch
}

func New() *MockInteractor {
	return &MockInteractor{
		msgChan:   make(chan model.Message, 10),
		responses: make(chan model.Response, 10),
	}
}

func (mi *MockInteractor) Dispose() {
	close(mi.msgChan)
	close(mi.responses)
	mi.msgChan = nil
	mi.responses = nil
}

func (mi *MockInteractor) SlipMessages(sender *model.User, messages ...string) {
	for _, msg := range messages {
		mi.msgChan <- model.Message{
			Sender:         sender,
			Text:           msg,
			ReceiveTime:    time.Now(),
			IsPrivate:      true,
			IsButtonAction: false,
		}
	}
}

func (mi *MockInteractor) SlipButtonAction(sender *model.User, actionIDs ...int64) {
	for _, id := range actionIDs {
		mi.msgChan <- model.Message{
			Sender:         sender,
			ReceiveTime:    time.Now(),
			IsPrivate:      true,
			IsButtonAction: true,
			ButtonActionID: id,
		}
	}
}

func (mi *MockInteractor) Expect(message string) {
	mi.expectations = append(mi.expectations, message)
}

func (mi *MockInteractor) MessageChan() chan model.Message {
	return mi.msgChan
}

func (mi *MockInteractor) SendResponse(r model.Response) {
	mi.responses <- r
}

func (mi *MockInteractor) AssertErrors(t *testing.T) {
	mi.collectUnsentMessages()
	mi.collectUncaughtResponses()
	for _, msg := range mi.unsentMessages {
		t.Errorf(`unsent message: "%s"`, msg.Text)
	}
	for _, r := range mi.uncaughtResponses {
		t.Errorf(`uncaught response: "%s"`, r.Text)
	}
	for _, r := range mi.unmatchedResponses {
		t.Errorf(`unmatched response: "%s", expect: "%s"`, r.response.Text, r.expectation)
	}
}

func (mi *MockInteractor) collectUnsentMessages() {
	for {
		select {
		case msg := <-mi.msgChan:
			mi.unsentMessages = append(mi.unsentMessages, msg)
		default:
			return
		}
	}
}

func (mi *MockInteractor) collectUncaughtResponses() {
	for {
		select {
		case r := <-mi.responses:
			if len(mi.expectations) > 0 {
				if r.Text != mi.expectations[0] {
					mi.unmatchedResponses = append(mi.unmatchedResponses,
						responseMatch{
							expectation: mi.expectations[0],
							response:    r,
						})
				}
				mi.expectations = mi.expectations[1:]
			} else {
				mi.uncaughtResponses = append(mi.uncaughtResponses, r)
			}
		default:
			return
		}
	}
}
