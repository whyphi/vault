package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type UsersData struct {
	Users []User `yaml:"users"`
}

func main() {
	data, err := os.ReadFile("vault.yaml")
	if err != nil {
		log.Fatalf("Error reading vault.yaml: %v", err)
	}

	var userData UsersData

	if err := yaml.Unmarshal(data, &userData); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	for _, user := range userData.Users {
		fmt.Printf("Name: %s, Email: %s \n", user.Name, user.Email)
	}
}
