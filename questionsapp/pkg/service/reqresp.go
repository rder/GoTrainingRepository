package service

type (

	CreateQuestionResponse struct {
		ID        int32 `json:"id"`
	}

	CreateQuestionRequest struct {
		IDUser    int32 `json:"idUser"`
		Subject  string `json:"subject"`
		Description string `json:"description"`   
	}

	CreateAnswerRequest struct {
		IDUser    int32 `json:"idUser"`
		IDquestion  string `json:"idquestion"`
		Description string `json:"description"`   
	}


	GetAllQuestionResponse struct {
		ID        int32 `json:"id"`
		IDUser    int32 `json:"idUser"`
		Subject  string `json:"subject"`
		Description string `json:"description"`   
	}
	
	GetAllQuestionsByUserResponse struct {
		ID        int32 `json:"id"`
		IDUser    int32 `json:"idUser"`
		Subject  string `json:"subject"`
		Description string `json:"description"`   
	}

	GetAllQuestionsByUserRequest struct {
		IDUser    int32 `json:"idUser"`	
	}

	GetQuestionByIDRequest struct {
		ID  int32 `json:"id"`	
	}
	
	
)




