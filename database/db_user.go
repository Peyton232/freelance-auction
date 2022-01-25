package database

import (
	"context"
	"log"
	"time"

	"github.com/Peyton232/freelance-auction/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ----------------------------------------------Mutation-----------------------------------------------
//Creation
func (db *DB) CreateUser(input *model.NewUser) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newID := primitive.NewObjectID().Hex()
	user := &model.User{
		UserID: newID,
		Email:  *input.Email,
		Name:   *input.Name,
		Phone:  *input.Phone,
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil || res == nil {
		log.Print(err)
		log.Print("\nunable to insert user into DB in database package\n")
		return nil
	}
	return user
}

// Deletion
func (db *DB) RemoveUser(userid string) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := db.FindUserByID(userid)
	if user == nil {
		return nil
	}

	collection.FindOneAndDelete(ctx, bson.M{"userid": userid})
	return user
}

// ----------------------------------------------QUERIES-----------------------------------------------
func (db *DB) FindUserByID(ID string) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	collection.FindOne(ctx, bson.M{"userid": ID}).Decode(&user)
	return &user
}
