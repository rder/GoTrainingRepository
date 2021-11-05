package service

import "context"

type Service interface {
	GetAllQuestions(ctx context.Context) ([]Question, error)
	GetAllQuestionsByUser(ctx context.Context, id int) ([]Question, error)
	CreateQuestion(ctx context.Context, q Question) (int, error)
	DeleteQuestion (ctx context.Context, q Question) (int, error)
	UpdateQuestion (ctx context.Context, a Answer) (int, error)
	
}