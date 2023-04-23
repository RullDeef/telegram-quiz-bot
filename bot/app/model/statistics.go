package model

type Statistics struct {
	UserID                int64
	QuizzesCompleted      uint
	MeanQuizCompleteTime  float64
	MeanQuestionReplyTime float64
	CorrectReplies        uint
	CorrectRepliesPercent float64
}
