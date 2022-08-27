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
	userRoutes := gorillaMux.PathPrefix("/user").Subrouter()
	gorillaMux.HandleFunc("/", handlers.UserLoginPageHandler)
	userRoutes.HandleFunc("/authenticate", handlers.AuthenticateUserHandler)
	userRoutes.HandleFunc("/register", handlers.UserRegisterPageHandler)
	userRoutes.HandleFunc("/dashboard", handlers.UserDashboardPageHandler)
	userRoutes.HandleFunc("/logout", handlers.UserLogoutHandler)

	// Admin's
	adminRoutes := gorillaMux.PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/login", handlers.AdminLoginPageHandler)
	adminRoutes.HandleFunc("/authenticate", handlers.AdminAuthenticateHandler)
	adminRoutes.HandleFunc("/dashboard", handlers.AdminDashboardPageHandler)
	adminRoutes.HandleFunc("/logout", handlers.AdminLogoutHandler)

	// Admin operations handlers
	adminRoutes.HandleFunc("/new-user-form", handlers.AdminNewUserHandler)
	adminRoutes.HandleFunc("/create-new-user", handlers.AdminCreateUserHandler)
	adminRoutes.HandleFunc("/update-user-form/{userId}", handlers.UpdateUserPageHandler)
	adminRoutes.HandleFunc("/update-user/{userId}", handlers.UpdateUserHandler)
	adminRoutes.HandleFunc("/delete-user/{userId}", handlers.DeleteUserHandler)

	log.Fatal(http.ListenAndServe(":8080", gorillaMux))
}
