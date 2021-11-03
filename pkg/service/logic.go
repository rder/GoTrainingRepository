package service

import (
	"context"
	"log"

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

func (s srv) GetAllQuestions () (question []Question){
	logger:= log.With(s.logger, "method","GetAllQuestions")
	return CreateQuestionMongo(question)

}

func (s srv) GetAllQuestionsByUser (ctx context.Context, id int) (question []Question){

	logger:= log.With(s.logger, "method","GetAllQuestionsByUser")
	return GetAllQuestionsByUserMongo(id)

}

func (s srv) CreateQuestion (ctx context.Context, question Question) (int){

	logger:= log.With(s.logger, "method","CreateQuestion")
	return GetAllQuestionsMongo()

}
