package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Function to insert a user into MongoDB
func insertUser(client *mongo.Client, user User) error {
	collection := client.Database("vault").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}

// Function to delete a user from MongoDB
func deleteUser(client *mongo.Client, user User) error {
	collection := client.Database("vault").Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"email": user.Email})
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return err
}

func getUsersFromMongoDB(client *mongo.Client) ([]User, error) {
	collection := client.Database("vault").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

// Function to compare two lists of users and identify new and users to delete
func compareUsers(yamlUsers, mongoUsers []User) ([]User, []User) {
	var newUsers []User
	var usersToDelete []User

	// Create a map of existing users
	existingUserMap := make(map[string]bool)
	for _, user := range mongoUsers {
		existingUserMap[user.Email] = true
	}

	newUserMap := make(map[string]bool)
	for _, user := range yamlUsers {
		newUserMap[user.Email] = true
	}

	// Identify new users and users to delete
	for _, user := range yamlUsers {
		if _, exists := existingUserMap[user.Email]; !exists {
			newUsers = append(newUsers, user)
		}
	}

	for _, user := range mongoUsers {
		if _, exists := newUserMap[user.Email]; !exists {
			usersToDelete = append(usersToDelete, user)
		}
	}

	return newUsers, usersToDelete
}
