package main

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	Lastname  bool   `json:"lastName"`
}

type Users []User
