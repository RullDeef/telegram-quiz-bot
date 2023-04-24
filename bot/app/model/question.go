package model

type Question struct {
	ID      int64
	Text    string
	Topic   string
	Answers []Answer
}

func (q Question) HasCorrectAnswer() bool {
	for _, answer := range q.Answers {
		if answer.IsСorrect {
			return true
		}
	}
	return false
}
