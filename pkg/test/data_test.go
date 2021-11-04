package test
import ("testing"
		"github.com/qanda/pkg/data"
)

func TestGetAllQuestions(t *testing.T) {
	results := data.GetAllQuestionsMongo()	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},
		Question {
			IdUser: 2, 
			Subject:"Maths 2", 
			Description: "Algorithms 2", 
		},
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}

func TestGetAllQuestionsByUser(t *testing.T) {
	results := data.GetAllQuestionsByUserMongo(1)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}


func TestGetAllQuestionsByUser(t *testing.T) {
	results := data.CreateQuestionMongo()	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}


func TestDeleteQuestion(t *testing.T) {
	results := data.DeleteQuestionMongo(2)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}

func TestGetUpdateQuestion(t *testing.T) {

	answer:= Answer{			
			IDQuestion : 1 ,
			Description : "Algorithm 3", 		
	}

   
	results := data.UpdateQuestionMongo(answer)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithm 3",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}