package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

type Users []User

// func (u *User) getUser(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }
// func (u *User) updateUser(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }
// func (u *User) deleteUser(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (u *User) createUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("INSERT INTO users(firstName, lastName) VALUES('%s', %s)", u.FirstName, u.Lastname)
// 	_, err := db.Exec(statement)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetUsers(db *sql.DB, start, count int) ([]User, error) {
// 	return nil, errors.New("Not implemented")
// }
