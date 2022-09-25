package handlers

import (
	"authentication_with_db/database"
	"authentication_with_db/models"
	"authentication_with_db/template"
	"authentication_with_db/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// -------------------------------------------------------------------------------------------------------
// -------------------------------- USER SECTION HANDLERS STARTS FROM HERE ------------------------------
// -------------------------------------------------------------------------------------------------------

func UserLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	session, err := utils.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if session.Values["userAuthenticated"] == true {
		http.Redirect(w, r, "/user/dashboard", http.StatusFound)
	} else {
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		data := models.UserModel{PageTitle: "User Login"}
		err = template.Tpl.ExecuteTemplate(w, "index.gohtml", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")

	session, _ := utils.Store.Get(r, "session")

	isValid := database.LoginUser(email, pass)

	if isValid {
		session.Values["userAuthenticated"] = true
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		http.Redirect(w, r, "/user/dashboard", http.StatusFound)
	} else {
		session.Values["userAuthenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UserRegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.PostFormValue("name")
	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")

	session, _ := utils.Store.Get(r, "session")

	isValid := database.RegisterUser(name, email, pass)

	if isValid {
		data := models.UserModel{
			PageTitle: "User Dashboard",
			UserName:  name,
			UserEmail: email,
		}
		models.InitUserModel(data)
		session.Values["userAuthenticated"] = true
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		//
		http.Redirect(w, r, "/user/dashboard", http.StatusFound)
	} else {
		session.Values["userAuthenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UserDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	user := models.ReturnUserModel()

	session, _ := utils.Store.Get(r, "session")

	if user != nil {
		session.Values["userAuthenticated"] = true
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		err := template.Tpl.ExecuteTemplate(w, "user_dashboard.gohtml", *user)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		session.Values["userAuthenticated"] = false
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")

	// Clearing cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	session.Values["userAuthenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// -------------------------------------------------------------------------------------------------------
// -------------------------------- ADMIN SECTION HANDLERS STARTS FROM HERE ------------------------------
// -------------------------------------------------------------------------------------------------------

func AdminLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")

	if session.Values["adminAuthenticated"] != true {
		data := models.AdminModel{
			PageTitle: "Admin Login",
		}

		err := template.Tpl.ExecuteTemplate(w, "admin_login.gohtml", data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
	}
}

func AdminAuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("emailId")
	pass := r.PostFormValue("passwordVal")
	session, _ := utils.Store.Get(r, "session")

	adminName, isValid := database.LoginAdmin(email, pass)

	if adminName != "" && isValid {

		session.Values["adminAuthenticated"] = true
		session.Values["adminName"] = adminName
		session.Values["adminEmail"] = email
		session.Save(r, w)

		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
	} else {
		session.Values["adminAuthenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func AdminDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	// TODO: Implement session and cookies or jwt
	data := models.ReturnAdminModel()
	if data != nil {
		session.Values["adminAuthenticated"] = true
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		err := template.Tpl.ExecuteTemplate(w, "admin_dashboard.gohtml", *data)
		if err != nil {
			log.Fatal("Error is here ", err)
		}
	} else {
		session.Values["adminAuthenticated"] = false
		session.Save(r, w)
		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func AdminLogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")

	session.Values["adminAuthenticated"] = false
	session.Values["adminName"] = ""
	session.Values["adminEmail"] = ""
	session.Save(r, w)

	// Clearing cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func AdminNewUserHandler(w http.ResponseWriter, r *http.Request) {
	type title struct {
		PageTitle string
	}

	data := title{
		PageTitle: "Create User",
	}

	session, _ := utils.Store.Get(r, "session")

	if session.Values["adminAuthenticated"] == true {
		err := template.Tpl.ExecuteTemplate(w, "create-user-page.gohtml", data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func AdminCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.PostFormValue("userName")
	email := r.PostFormValue("userEmail")
	pass := r.PostFormValue("userPassword")
	product1 := r.PostFormValue("userProduct1")
	product2 := r.PostFormValue("userProduct2")
	product3 := r.PostFormValue("userProduct3")
	product4 := r.PostFormValue("userProduct4")
	product5 := r.PostFormValue("userProduct5")

	// Encrypting password
	encPass, err := utils.HashEncrypt(pass)
	if err != nil {
		log.Fatal(err.Error())
	}

	session, _ := utils.Store.Get(r, "session")

	data := models.UserModel{
		UserName:     name,
		UserEmail:    email,
		UserProduct1: product1,
		UserProduct2: product2,
		UserProduct3: product3,
		UserProduct4: product4,
		UserProduct5: product5,
	}

	adminName := session.Values["adminName"]
	adminEmail := session.Values["adminEmail"]

	database.CreateUser(data, encPass)
	// variable_name.(type) is known as type assertion
	database.GetValuesForAdminModel(adminName.(string), adminEmail.(string))

	// Clearing cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
}

func UpdateUserPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	data := database.GetUserData(userId)
	data.PageTitle = "Update User"

	session, _ := utils.Store.Get(r, "session")

	if session.Values["adminAuthenticated"] == true {
		err := template.Tpl.ExecuteTemplate(w, "update_user_page.gohtml", data)
		if err != nil {
			log.Fatal("Error is here ", err.Error())
		}
	} else {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	r.ParseForm()

	product1 := r.PostFormValue("userProduct1")
	product2 := r.PostFormValue("userProduct2")
	product3 := r.PostFormValue("userProduct3")
	product4 := r.PostFormValue("userProduct4")
	product5 := r.PostFormValue("userProduct5")

	if database.UpdateUserDatas(userId, product1, product2, product3, product4, product5) {

		session, _ := utils.Store.Get(r, "session")

		adminName := session.Values["adminName"]
		adminEmail := session.Values["adminEmail"]

		database.GetValuesForAdminModel(adminName.(string), adminEmail.(string))

		// Clearing cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
	} else {
		route := fmt.Sprintf("/admin/update-user-form/%s", userId)
		http.Redirect(w, r, route, http.StatusFound)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	database.DeleteUser(userId)

	session, _ := utils.Store.Get(r, "session")

	adminName := session.Values["adminName"]
	adminEmail := session.Values["adminEmail"]

	// variable_name.(type) is known as type assertion
	database.GetValuesForAdminModel(adminName.(string), adminEmail.(string))

	// Clearing cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
}
