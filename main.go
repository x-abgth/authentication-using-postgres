package main

import (
	"authentication_with_db/database"
	"authentication_with_db/handlers"
	"authentication_with_db/template"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Served static html contents
func init() {
	var err error

	template.Tpl, err = template.Tpl.ParseGlob("./views/*.gohtml")
	template.Tpl.New("partials").ParseGlob("./views/partials/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gorillaMux := mux.NewRouter()

	// database connections
	database.ConnectDb()
	defer database.Db.Close()

	// serving other files like css, and images using only http package
	fileServe := http.FileServer(http.Dir("./views/assets/"))
	gorillaMux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))

	// user's
	gorillaMux.HandleFunc("/", handlers.UserLoginPageHandler)
	gorillaMux.HandleFunc("/user-authenticate", handlers.AuthenticateUserHandler)
	gorillaMux.HandleFunc("/user-register", handlers.UserRegisterPageHandler)
	gorillaMux.HandleFunc("/user-dashboard", handlers.UserDashboardPageHandler)

	// Admin's
	gorillaMux.HandleFunc("/admin-login", handlers.AdminLoginPageHandler)
	gorillaMux.HandleFunc("/admin-authenticate", handlers.AdminAuthenticateHandler)
	gorillaMux.HandleFunc("/admin-dashboard", handlers.AdminDashboardPageHandler)

	// Admin operations handlers
	//gorillaMux.HandleFunc("/update-user/{userId}", handlers.UpdateUserHandler).Methods("PUT")
	gorillaMux.HandleFunc("/delete-user/{userId}", handlers.DeleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", gorillaMux))
}
