package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/questions/").Handler(httptransport.NewServer(
		endpoints.GetAllQuestions,
		encodeResponse,		
	))

	r.Methods("GET").Path("/questionsUser/{id}").Handler(httptransport.NewServer(
		endpoints.GetAllQuestionsByUser,
		decodeQuestionUserRequest,
		encodeResponse,		
	))

	r.Methods("POST").Path("/createquestions/{id}").Handler(httptransport.NewServer(
		endpoints.CreateQuestion,
		decodeQuestionCreateRequest,
		encodeResponse,		
	))

	r.Methods("GET").Path("/question/{id}").Handler(httptransport.NewServer(
		endpoints.GetQuestionByID,
		decodeGetQuestionByIDRequest,
		encodeResponse,		
	))

	r.Methods("PUT").Path("/updatequestion/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateQuestion,
		decodeGetQuestionByIDRequest,
		encodeResponse,		
	))
    
	r.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteQuestion,
		decodeGetQuestionByIDRequest,
		encodeResponse,		
	))
	

	fmt.Println("Start listening")
    fmt.Println(http.ListenAndServe(":8080", r))
	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

/*Encode & Decode*/



func decodeGetQuestionByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetQuestionByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeQuestionUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetAllQuestionsByUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeQuestionCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
