package model

type SessionState struct {
	Quiz            Quiz
	Users           []User
	CurrentQuestion *Question
	IsPaused        bool
}

func NewSessionState(q Quiz, users []User) *SessionState {
	return &SessionState{
		Quiz:            q,
		Users:           users,
		CurrentQuestion: &q.Questions[0],
		IsPaused:        false,
	}
}
