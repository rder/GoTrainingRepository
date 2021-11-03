package service

import "context"

type Service interface {
	GetAllQuestions(ctx context.Context) ([]Question, error)
	GetAllQuestionsByUser(ctx context.Context, id int) ([]Question, error)
	CreateQuestion(ctx context.Context, q Question) (int, error)
	
}