package domain

import "context"

type Feedback struct {
	ID       int    `db:"id"`
	Feedback string `db:"feedback"`
}

type FeedbackRepository interface {
	GetAll(ctx context.Context) ([]Feedback, error)
	Register(ctx context.Context, data *Feedback) error
}

type FeedbackService interface {
	GetAll(ctx context.Context) ApiResponse
	Register(ctx context.Context, data FeedbackData) ApiResponse
}

type FeedbackData struct {
	Feedback string `json:"feedback"`
}
