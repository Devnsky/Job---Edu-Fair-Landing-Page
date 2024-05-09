package feedback

import (
	"context"
	"eduFair/domain"
)

type service struct {
	FeedbackRepository domain.FeedbackRepository
}

func NewService(feedbackRepository domain.FeedbackRepository) domain.FeedbackService {
	return &service{
		FeedbackRepository: feedbackRepository,
	}
}

func (s service) GetAll(ctx context.Context) domain.ApiResponse {
	data, err := s.FeedbackRepository.GetAll(ctx)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}
	var feedbacks []domain.FeedbackData
	for _, v := range data {
		feedbacks = append(feedbacks, domain.FeedbackData{
			Feedback: v.Feedback,
		})
	}
	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    feedbacks,
	}
}

func (s service) Register(ctx context.Context, data domain.FeedbackData) domain.ApiResponse {
	feeedbackModel := domain.Feedback{
		Feedback: data.Feedback,
	}
	err := s.FeedbackRepository.Register(ctx, &feeedbackModel)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    feeedbackModel,
	}
}
