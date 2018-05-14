package main

func main() {
	a := App{}

	a.Initialize("root", "", "user_management")

	a.Run(":8080")
}
