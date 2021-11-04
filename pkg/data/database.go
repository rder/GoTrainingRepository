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
	//defer client.Disconnect(ctxM)  // Disconnect from DB
}

type Database interface{
	CreateQuestionDB (interface{}) interface{};
	GetAllQuestionsDB() interface{};
	DeleteQuestionDB(id int) interface{}
} 



func CreateQuestionMongo(question Question) int{	
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

func GetQuestionByIDMongo (id int) (interface{}){

	question,err:= collectionQ.findOne(ctxM,bson.M{"_id": &id})

	if err!=nil {
		log.Fatal("Error in db",err)
	}

	 
	return question
}


func GetAllQuestionsByUserMongo(idUser int) []interface{}{
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

func DeleteQuestionMongo(idQ int) []interface{}{

	// Delete record
	err := collectionA.Remove(bson.M{"_id": &idQ})
	if err != nil {
		log.Fatal("Remove Question fail %v\n", err)
		
	}

	errQ := collectionQ.Remove(bson.M{"_id": &idQ})
	if errQ != nil {
		log.Fatal("Remove Question fail %v\n", errQ)
		
	}
	return errQ.DeletedCount
}


func UpdateQuestionMongo(id int) []interface{}{
	
    update := bson.M{"$set": bson.M{"_id": &id}}

    result, err := collectionA.UpdateOne(
        context.Background(),
        update,
    )
	if err != nil {
		log.Fatal("Update Question fail %v\n", err)
		
	}	
	return result.UpsertedID , err
}

