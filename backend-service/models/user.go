package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"webkodes.com/admin/utils"
)

// Tokens  model for storing jwt token
type Tokens struct {
	AccessToken string    `json:"access_token" bson:"access_token,omitempty"`
	Validity    time.Time `json:"validity" bson:"validity,omitempty"`
}

// User model for storing user data
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	Role      string             `json:"role" bson:"role,omitempty"`
	Tokens    Tokens             `json:"tokens" bson:"tokens"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
}

// InsertUser method for inserting a single user
func InsertUser(user User) {
	client := utils.GetClient()
	collection := client.Database("admin").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalln("Error on inserting new User", err)
	}
	fmt.Println(insertResult)
}

// GetUsers method for retrieving all users
func GetUsers(filter bson.M) []*User {
	var users []*User
	client := utils.GetClient()
	collection := client.Database("admin").Collection("users")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var user User
		err = cur.Decode(&user)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		users = append(users, &user)
	}
	return users
}

// GetSingleUser method for retrieving a single user
func GetSingleUser(filter bson.M) User {
	var user User
	client := utils.GetClient()
	collection := client.Database("admin").Collection("users")
	documentReturned := collection.FindOne(context.TODO(), filter)
	err := documentReturned.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

// UpdateUser method for updating a single user
func UpdateUser(updatedUser bson.M, filter bson.M) int64 {
	client := utils.GetClient()
	collection := client.Database("admin").Collection("users")
	newUserData := bson.D{{Key: "$set", Value: updatedUser}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, newUserData)
	if err != nil {
		log.Fatal("Error on updating one User", err)
	}
	return updatedResult.ModifiedCount
}

// DeleteUser method for deleting a single user
func DeleteUser(filter bson.M) int64 {
	client := utils.GetClient()
	collection := client.Database("admin").Collection("users")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one User", err)
	}
	return deleteResult.DeletedCount
}
