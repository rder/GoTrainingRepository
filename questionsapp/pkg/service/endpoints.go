package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	GetAllQuestions endpoint.Endpoint
    GetAllQuestionsByUser endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
	GetQuestionByID endpoint.Endpoint
	UpdateQuestion endpoint.Endpoint
	DeleteQuestion endpoint.Endpoint
}


func MakeEndpoints (s Service) Endpoints{
	return Endpoints{
		GetAllQuestions: makeGetAllQuestionsEndpoint(s),
		GetAllQuestionsByUser: makeGetAllQuestionsByUserEndpoint(s),
		CreateQuestion : makeCreateQuestionEndpoint(s),		
		GetQuestionByID : makeGetQuestionByIDEndpoint(s),	
    	UpdateQuestion : makeUpdateQuestionEndpoint(s),
    	DeleteQuestion : makeDeleteQuestionEndpoint(s),

	}
}

func makeGetAllQuestionsEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,request interface{})(interface{},error)	{
		
		question:= s.GetAllQuestions(ctx)
		return question,nil
	}
	
}

func makeGetAllQuestionsByUserEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,request interface{})(interface{},error)	{
		question:= s.GetAllQuestionsByUser(ctx,interface{})
		return question,nil
	}

	
}

func makeCreateQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error)	{
		req := request.(CreateQuestionRequest)
		q:= s.CreateQuestion(ctx,req.question)
		return q.id,nil
	}
}

func makeGetQuestionByIDEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(interface{},error){
		req := request.(CreateQuestionRequest)
		q:= s.GetQuestionByID(ctx,req.Question)
		return q,nil
	}
}

func makeUpdateQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(CreateAnswerRequest)
		q:= s.UpdateQuestion(ctx,req.Answer)
		return q,nil
	}
}

func makeDeleteQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(CreateQuestionRequest)
		q:= s.DeleteQuestion(ctx,question)
		return q,nil
	}
}

    	