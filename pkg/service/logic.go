package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/qanda/pkg/data"
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

func (s srv) GetAllQuestions () (question []Question){
	logger:= log.With(s.logger, "method","GetAllQuestions")
	return data.GetAllQuestionsMongo(question)

}

func (s srv) GetAllQuestionsByUser (ctx context.Context, id int) (question []Question){

	logger:= log.With(s.logger, "method","GetAllQuestionsByUser")
	return data.GetAllQuestionsByUserMongo(id)

}

func (s srv) CreateQuestion (ctx context.Context, question Question) (int){

	logger:= log.With(s.logger, "method","CreateQuestion")
	return data.CreateQuestionMongo()

}

func (s srv) GetQuestionByID(ctx context.Context, id int)(question []Question){
	logger:= log.With(s.logger, "method","GetQuestionByID")
	return data.GetQuestionByIDMongo(id)
}


func (s srv) UpdateQuestion (ctx context.Context, id int) (int error){

	logger:= log.With(s.logger, "method","UpdateQuestion")	
	return data.UpdateQuestionMongo(id)
}

func (s srv) DeleteQuestion (ctx context.Context, id int) (int error){

	logger:= log.With(s.logger, "method","DeleteQuestion")
	return data.DeleteQuestionMongo(id)
}
