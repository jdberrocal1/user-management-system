package userHandler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implemented!")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(users); err != nil {
	// 	panic(err)
	// }
	fmt.Fprintln(w, "Not Implemented!")
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// var userId int
	// var err error
	// if userId, err = strconv.Atoi(vars["userId"]); err != nil {
	// 	panic(err)
	// }
	// user := RepoFindUser(userId)
	// if user.ID > 0 {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	if err := json.NewEncoder(w).Encode(user); err != nil {
	// 		panic(err)
	// 	}
	// 	return
	// }

	// // If we didn't find it, 404
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusNotFound)
	// if err := json.NewEncoder(w).Encode(ResponseHandler.JSONErr{Code: http.StatusNotFound, Message: "Not Found"}); err != nil {
	// 	panic(err)
	// }
	fmt.Fprintln(w, "Not Implemented!")
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	// var user UserModel.User
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &user); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }

	// u := RepoCreateUser(user)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(u); err != nil {
	// 	panic(err)
	// }
	fmt.Fprintln(w, "Not Implemented!")
}
