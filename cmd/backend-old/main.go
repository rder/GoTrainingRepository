package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type Questions struct {
    ID        int32 `json:"id"`
    IDUser    int32 `json:"idUser"`
    subject  string `json:"subject"`
    description string `json:"description"`   
}

type Answers struct {
    ID        int32 `json:"id"`
    IDquestion  int32 `json:"IDquestion"`
    description string `json:"description"`   
   
}

/*

Get one question by its ID
Get a list of all questions
Get all the questions created by a given user
Create a new question
Update an existing question (the statement and/or the answer)
Delete an existing question

*/


var questions = []Questions{}
var answers = []Answers{}

var idCounter int
func main() {
    r := mux.NewRouter()
    usersR := r.PathPrefix("/questions").Subrouter()
    usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllQuestions)
    usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getAllQuestionsByUser)
    usersR.Path("").Methods(http.MethodPost).HandlerFunc(createQuestion)
    usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getQuestionByID)
    usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(updateQuestion)
    usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteQuestion)
    fmt.Println("Start listening")
    fmt.Println(http.ListenAndServe(":8080", r))
}

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(questions); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}

func getAllQuestionsByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(questions); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}

func getQuestionByID(w http.ResponseWriter, r *http.Request) {
	id ,_:= strconv.Atoi(mux.Vars(r)["id"])
    index := indexByID(questions, id)
    if index < 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(questions[index]); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}

func updateQuestion(w http.ResponseWriter, r *http.Request) {
    id ,_:= strconv.Atoi(mux.Vars(r)["id"])
    index := indexByID(questions, id)
    if index < 0 {
        http.Error(w, "Questions not found", http.StatusNotFound)
        return
    }
    u := Questions{}
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        fmt.Println(err)
        http.Error(w, "Error decoidng response object", http.StatusBadRequest)
        return
    }
    questions[index] = u
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(&u); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}
func deleteQuestion(w http.ResponseWriter, r *http.Request) {
    id,_ := strconv.Atoi(mux.Vars(r)["id"])
    index := indexByID(questions, id)
    if index < 0 {
        http.Error(w, "Question not found", http.StatusNotFound)
        return
    }
    questions = append(questions[:index], questions[index+1:]...)
    w.WriteHeader(http.StatusOK)
}

func createQuestion(w http.ResponseWriter, r *http.Request) {
    u := Questions{}
    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        fmt.Println(err)
        http.Error(w, "Error decoidng response object", http.StatusBadRequest)
        return
    }
    questions = append(questions, u)
    response, err := json.Marshal(&u)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(response)
}

func indexByID(question []Questions, id int) int {
    for i := 0; i < len(question); i++ {
        if int(question[i].ID) == id {
            return i
        }
    }
    return -1
}