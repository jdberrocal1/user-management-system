package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	routesConfig "github.com/user-management-system/cmd/routes"
	router "github.com/user-management-system/pkg/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App define the app structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize create the database connection
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = initializeRoutes()
}

func initializeRoutes() *mux.Router {
	return router.NewRouter(routesConfig.GetRoutes())
}

// Run start the application
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// func RespondWithError(w http.ResponseWriter, error JsonErr) {
// 	respondWithJSON(w, error.Code, map[string]string{"error": error.Message})
// }

// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

// func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
// var u UserModel.User
// decoder := json.NewDecoder(r.Body)
// if decodeError := decoder.Decode(&u); decodeError != nil {
// 	var err = ResponseHandler.JSONErr{http.StatusBadRequest, "Invalid request payload"}
// 	ResponseHandler.WithError(w, err)
// 	return
// }
// defer r.Body.Close()
// if createError := u.createUser(a.DB); createError != nil {
// 	var err = ResponseHandler.JSONErr{http.StatusInternalServerError, createError.Error()}
// 	ResponseHandler.WithError(w, err)
// 	return
// }
// ResponseHandler.WithJSON(w, http.StatusCreated, u)
// }
