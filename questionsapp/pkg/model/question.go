package model

type Question struct {
    ID        int32 `json:"id"`
    IDUser    int32 `json:"idUser"`
    Subject  string `json:"subject"`
    Description string `json:"description"`   
}

type Answer struct {
    IDAnswer        int32 `json:"idAnswer"`
    IDQuestion   int32 `json:"idQuestion"`
    Description string `json:"description"`   
}


/*
type Repository interface{
    
	GetAllQuestions (ctx context.Context) (question []Question)
    GetAllQuestionsByUser (ctx context.Context, id int) (question []Question)
    CreateQuestion (ctx context.Context, question Question) (int)
    GetQuestionByID (ctx context.Context, id int) (Question,error)
    UpdateQuestion (ctx context.Context, question Question) (int,error)
    DeleteQuestion(ctx context.Context, id int) (int,error)

}*/
