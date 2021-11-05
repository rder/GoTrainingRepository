package service

import (
	"context"

	"questionsapp/pkg/data"
	"questionsapp/pkg/model"

	"github.com/go-kit/kit/log"
)

type srv struct{
	repository data.Repository
	logger log.Logger
}


func NewService (rep data.Repository, logger log.Logger) Service{

	return &srv{ 
		repository: rep,
		logger: logger,
	}
}

func (s srv) GetAllQuestions (ctx context.Context) (question []model.Question){
	logger:= log.With(s.logger, "method","GetAllQuestions")
	return data.GetAllQuestion(ctx)

}

func (s srv) GetAllQuestionsByUser (ctx context.Context, id int) (question []model.Question){

	//logger:= log.With(s.logger, "method","GetAllQuestionsByUser")
	return data.GetAllQuestionsByUser(ctx,id)

}

func (s srv) CreateQuestion (ctx context.Context, question model.Question) (int,error){

	//logger:= log.With(s.logger, "method","CreateQuestion")
	return data.CreateQuestion(ctx,question)

}

func (s srv) GetQuestionByID(ctx context.Context, id int)(question []model.Question){
	//logger:= log.With(s.logger, "method","GetQuestionByID")
	return data.GetQuestionByIDMongo(id)
}


func (s srv) UpdateQuestion (ctx context.Context, answer model.Answer) (int error){

	//logger:= log.With(s.logger, "method","UpdateQuestion")	
	return data.UpdateQuestion(ctx, answer)
}

func (s srv) DeleteQuestion (ctx context.Context, id int) (int error){

//	logger:= log.With(s.logger, "method","DeleteQuestion")
	return data.DeleteQuestionMongo(ctx,id)
}
