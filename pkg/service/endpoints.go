package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	GetAllQuestions endpoint.Endpoint
    GetAllQuestionsByUser endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
}


func MakeEndpoints (s Service) Endpoints{
	return Endpoints{
		GetAllQuestions: makeGetAllQuestionsEndpoint(s),
		GetAllQuestionsByUser: makeGetAllQuestionsByUserEndpoint(s),
		CreateQuestion : makeCreateQuestionEndpoint(s),
	}
}

func makeGetAllQuestionsEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,request interface{})(interface{},error)	{
		
		question:= s.GetAllQuestions(ctx)
		return question,nil
	}
	
}

func makeGetAllQuestionsByUserEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,id int)([]Question,error)	{
		question:= s.GetAllQuestionsByUser(ctx,id)
		return question,nil
	}

	
}

func makeCreateQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,question Question)(id int)	{
		q:= s.CreateQuestion(ctx,question)
		return q.id,nil
	}
}