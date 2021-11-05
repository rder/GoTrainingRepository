package service

import (
	"context"
	"questionsapp/pkg/model"
)

type Service interface {
	GetAllQuestions(ctx context.Context) ([]model.Question, error)
	GetAllQuestionsByUser(ctx context.Context, id int) ([]model.Question, error)
	CreateQuestion(ctx context.Context, q model.Question) (int, error)
	DeleteQuestion (ctx context.Context, q model.Question) (int, error)
	UpdateQuestion (ctx context.Context, a model.Answer) (int, error)
	GetQuestionByID(ctx context.Context, id int) (model.Question, error) 
}