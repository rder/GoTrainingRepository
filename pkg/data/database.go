package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Connection to MongoDB*/
var collectionQ *mongo.Collection
var collectionA *mongo.Collection
var ctxM = context.TODO()

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctxM, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

	err = client.Ping(ctxM, nil)
	if err != nil {
		log.Fatal(err)
	}
	/*Create DataBase & Collection Questions*/
	databaseQA := client.Database("QuestionsAnswers")
  	collectionQ = databaseQA.Collection("Questions")
	collectionA = databaseQA.Collection("Answers")
	defer client.Disconnect(ctxM)  // Disconnect from DB
}



func CreateQuestion(question Question) int{	
	insertResult, err := collectionQ.InsertOne(context.TODO(), question)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Question Insert: ", insertResult.InsertID)
	return insertResult.InsertID
}

func GetAllQuestions() []interface{}{
	questions,err:= collectionQ.find(ctxM,bson.M{})

	if err!=nil {
		log.Fatal("Error in db",err)
	}

	var results []bson.M   
    if err = questions.All(ctxM,&results);err!=nil {
        log.Fatal(err)
		panic(err)
    } 
	return results
}


func GetAllQuestionsByUser(idUser int) []interface{}{
	questions,err:= collectionQ.find(ctxM,bson.M{"_id": &idUser})	
	if err!=nil {
		log.Fatal("Error in db",err)
	}

	var results []bson.M
   
    if err = questions.All(ctxM,&results);err!=nil {
        log.Fatal(err)
    } 
	return results
}
