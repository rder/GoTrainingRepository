package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"questionsapp/pkg/model"

	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	mgo "gopkg.in/mgo.v2"
)
	var RepoErr = errors.New("Unable to handle Repo Request")
	const QuestionCollection = "Questions"
	const AnswerCollection = "Answers"


	type Repository interface {
		CreateQuestion(ctx context.Context, question model.Question) error
		GetAllQuestion(ctx context.Context) ([]model.Question, error)
		GetAllQuestionsByUser (ctx context.Context,id int) ([]model.Question , error)
		GetQuestionByID (ctx context.Context,id int) (model.Question , error)
		UpdateQuestion (ctx context.Context,id int) (int, error)
		DeleteQuestion (ctx context.Context,id int) (int, error)

	}
		
	// Repo structure for mongodb
	type mongorepo struct {
		db     *mgo.Database
		logger log.Logger
	}

	type sqlrepo struct {
		sqldb  *sql.DB
		logger log.Logger
	}

	// Repo for mongodb
	func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
		return &mongorepo{
				db:db,
				logger: log.With(logger, "repo", "mongodb"),}, nil
	}


	/*All Methods*/

	func (repo *mongorepo) CreateQuestion(ctx context.Context, question model.Question) error {
		fmt.Println("create question mongo repo", repo.db)
		err := repo.db.C(QuestionCollection).Insert(question)
		if err != nil {
			fmt.Println("Error occured inside CreateQuestion in repo")
		return err
		} else {
			fmt.Println("Question Created:", question.id)
		}
		return nil
	}

	func (repo *mongorepo) GetAllQuestion(ctx context.Context) ([]model.Question, error) {
		fmt.Println("All question mongo repo", repo.db)

		questions,err:= repo.db.C(QuestionCollection)(ctx,bson.M{}).Find()

		var results []bson.M   
		
		if err = questions.All(ctx,&results);err!=nil {
			log.Fatal(err)			
		} 
		
		return results, err
	}



	func (repo *mongorepo) GetAllQuestionsByUser(ctx context.Context, idUser int) ([]model.Question, error) {
		fmt.Println("All GetAllQuestionsByUser mongo repo", repo.db)

		questions,err:= repo.db.C(QuestionCollection).Find(ctx,bson.M{"_id": &idUser})

		var results []bson.M   
		
		if err = questions.All(ctx,&results);err!=nil {
			log.Fatal(err)			} 
		
		return results, err
	}

	


	func (repo *mongorepo) GetQuestionByID(ctx context.Context, id int) (model.Question, error) {
		fmt.Println("All GetAllQuestionsByUser mongo repo", repo.db)

		questions,err:= repo.db.C(QuestionCollection).findOne(ctx,bson.M{"_id": &id})
		var results bson.M   
		
		if err = questions.All(ctx,&results);err!=nil {
			log.Fatal(err)			
		} 		
		return results, err
	}


	func (repo *mongorepo) UpdateQuestion(ctx context.Context, answer model.Answer) (int, error) {
		fmt.Println("UpdateQuestion mongo repo", repo.db)
		
		update := bson.M{"$set": bson.M{"_id": &answer.IDQuestion}}
		result, err := repo.db.C(AnswerCollection).FindOne.UpdateOne(
			context.Background(),
			update,
		)
		if err != nil {
			log.Fatal("Update Question fail %v\n", err)
			
		}	
		return result.UpsertedID , err
	}


	func (repo *mongorepo) DeleteQuestion(ctx context.Context, answer model.Answer) (int, error) {
		fmt.Println("DeleteQuestion mongo repo", repo.db)
		
		err := repo.db.C(AnswerCollection).Remove(bson.M{"_id": &answer.idAnswer})
		if err != nil {
			log.Fatal("Remove Answer fail %v\n", err)
			
		}

		errQ := repo.db.C(QuestionCollection).Remove(bson.M{"_id": &answer.idQuestion})
		if errQ != nil {
			log.Fatal("Remove Question fail %v\n", errQ)
			
		}
		return errQ.DeletedCount ,err
	}

	/*OPTIONAL SQL*/
/*
	func (repo *sqlrepo) CreateQuestion(ctx context.Context, question service.Question) error {
			
			fmt.Println("create Question sql repo", sqldb),
			dbt:= sqlrepo.sqldb
			sqlresp, err := dbt.Prepare("INSERT INTO Question (userid,subject,description) VALUES (?,?,?)")
			sqlresp.Exec(question.idUser, question.subject, question.description)
			fmt.Println("Question SQL Created:", sqlresp, err)
			return nil
	}

	*/