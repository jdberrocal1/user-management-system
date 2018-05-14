package main

import (
	"fmt"

	UserModel "github.com/user-management-system/cmd/models/user"
)

var currentID int

var users UserModel.Users

// Give us some seed data
func init() {
	RepoCreateUser(UserModel.User{FirstName: "Subject1", Lastname: "Subject1 LastName"})
	RepoCreateUser(UserModel.User{FirstName: "Subject2", Lastname: "Subject2 LastName"})
}

func RepoFindUser(id int) UserModel.User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}
	// return empty Todo if not found
	return UserModel.User{}
}

func RepoCreateUser(u UserModel.User) UserModel.User {
	currentID++
	u.ID = currentID
	users = append(users, u)
	return u
}

func RepoDestroyUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
