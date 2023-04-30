package model

import "sync"

type SessionState struct {
	Quiz            Quiz
	Users           []*User
	CurrentQuestion *Question
	isPaused        bool
	mutex           *sync.RWMutex

	pauseChan  chan struct{}
	resumeChan chan struct{}
}

func NewSessionState(q Quiz, users []*User) *SessionState {
	return &SessionState{
		Quiz:            q,
		Users:           users,
		CurrentQuestion: &q.Questions[0],
		isPaused:        false,
		mutex:           &sync.RWMutex{},

		pauseChan:  make(chan struct{}, 1),
		resumeChan: make(chan struct{}, 1),
	}
}

func (s *SessionState) WaitForPause() chan struct{} {
	return s.pauseChan
}

func (s *SessionState) WaitForResume() chan struct{} {
	return s.resumeChan
}

func (s *SessionState) Pause() {
	s.pauseChan <- struct{}{}
	s.mutex.Lock()
	s.isPaused = true
	s.mutex.Unlock()
}

func (s *SessionState) Resume() {
	s.resumeChan <- struct{}{}
	s.mutex.Lock()
	s.isPaused = false
	s.mutex.Unlock()
}

func (s *SessionState) IsPaused() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.isPaused
}
