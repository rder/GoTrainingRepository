package service

import (
	"context"

	"questionsapp/pkg/data"

	"github.com/go-kit/kit/log"
)

type srv struct{
	repository Repository
	logger log.Logger
}


func NewService (rep Repository, logger log.Logger) Service{

	return &srv{ 
		repository: rep,
		logger: logger,
	}
}

func (s srv) GetAllQuestions (ctx context.Context) (question []Question){
	logger:= log.With(s.logger, "method","GetAllQuestions")
	return data.GetAllQuestion(ctx)

}

func (s srv) GetAllQuestionsByUser (ctx context.Context, id int) (question []Question){

	//logger:= log.With(s.logger, "method","GetAllQuestionsByUser")
	return data.GetAllQuestionsByUser(ctx,id)

}

func (s srv) CreateQuestion (ctx context.Context, question Question) (int,error){

	//logger:= log.With(s.logger, "method","CreateQuestion")
	return data.CreateQuestion(ctx,question)

}

func (s srv) GetQuestionByID(ctx context.Context, id int)(question []Question){
	//logger:= log.With(s.logger, "method","GetQuestionByID")
	return data.GetQuestionByIDMongo(id)
}


func (s srv) UpdateQuestion (ctx context.Context, answer Answer) (int error){

	//logger:= log.With(s.logger, "method","UpdateQuestion")	
	return data.UpdateQuestion(ctx, answer)
}

func (s srv) DeleteQuestion (ctx context.Context, id int) (int error){

//	logger:= log.With(s.logger, "method","DeleteQuestion")
	return data.DeleteQuestionMongo(ctx,id)
}
