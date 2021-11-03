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
	
)

/*Encode & Decode encodeAllAnswersResponse([]Questions)*/
/*
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

*/




