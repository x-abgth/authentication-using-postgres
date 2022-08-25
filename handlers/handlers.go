package handlers

import (
	"authentication_with_db/database"
	"authentication_with_db/models"
	"authentication_with_db/template"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UserLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	data := models.UserModel{PageTitle: "User Login"}
	err := template.Tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

func AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")

	isValid := database.LoginUser(email, pass)

	if isValid {
		http.Redirect(w, r, "/user-dashboard", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UserRegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.PostFormValue("name")
	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")

	isValid := database.RegisterUser(name, email, pass)

	if isValid {
		data := models.UserModel{
			PageTitle: "User Dashboard",
			UserName:  name,
			UserEmail: email,
		}
		models.InitUserModel(data)

		http.Redirect(w, r, "/user-dashboard", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UserDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	user := models.ReturnUserModel()

	if user != nil {
		err := template.Tpl.ExecuteTemplate(w, "user_dashboard.gohtml", *user)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func AdminLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	data := models.AdminModel{
		PageTitle: "Admin Login",
	}
	// TODO: Implement session and cookies or jwt
	err := template.Tpl.ExecuteTemplate(w, "admin_login.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

func AdminAuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")

	adminName, isValid := database.LoginAdmin(email, pass)

	if adminName != "" && isValid {
		http.Redirect(w, r, "/admin-dashboard", http.StatusFound)
	} else {
		http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
	}
}

func AdminDashboardPageHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Implement session and cookies or jwt
	data := models.ReturnAdminModel()
	if data != nil {
		err := template.Tpl.ExecuteTemplate(w, "admin_dashboard.gohtml", *data)
		if err != nil {
			log.Fatal("Error is here ", err)
		}
	} else {
		http.Redirect(w, r, "/admin-login", http.StatusSeeOther)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	database.DeleteUser(userId)

}
