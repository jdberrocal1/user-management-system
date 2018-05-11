package main

import "fmt"

var currentId int

var users Users

// Give us some seed data
func init() {
	RepoCreateUser(User{FirstName: "Subject1"})
	RepoCreateUser(User{FirstName: "Subject2"})
}

func RepoFindUser(id int) User {
	for _, u := range users {
		if u.Id == id {
			return u
		}
	}
	// return empty Todo if not found
	return User{}
}

func RepoCreateUser(u User) User {
	currentId++
	u.Id = currentId
	users = append(users, u)
	return u
}

func RepoDestroyUser(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
