package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type User struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type UsersData struct {
	Users []User `yaml:"users"`
}

// func goDotEnvVariable(key string) string {
// 	err := godotenv.Load()

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

func main() {
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.9gtht.mongodb.net/?retryWrites=true&w=majority", mongoUser, mongoPassword)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	data, err := os.ReadFile("vault.yaml")
	if err != nil {
		log.Fatalf("Error reading vault.yaml: %v", err)
	}

	var userData UsersData

	if err := yaml.Unmarshal(data, &userData); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	// Retrieve existing users from MongoDB
	existingUsers, err := getUsersFromMongoDB(client)
	if err != nil {
		log.Fatalf("Error retrieving existing users from MongoDB: %v", err)
	}

	newUsers, usersToDelete := compareUsers(userData.Users, existingUsers)
	fmt.Println("New Users:")
	for _, user := range newUsers {
		fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)
	}

	fmt.Print("\n")
	fmt.Println("Users to Delete:")
	for _, user := range usersToDelete {
		fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)
	}

	// Insert new users into MongoDB
	for _, user := range newUsers {
		if err := insertUser(client, user); err != nil {
			log.Printf("Error inserting user into MongoDB: %v", err)
		}
	}

	// Delete users not in YAML file
	for _, user := range usersToDelete {
		if err := deleteUser(client, user); err != nil {
			log.Printf("Error deleting user from MongoDB: %v", err)
		}
	}

	return
}
