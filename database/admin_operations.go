package database

import (
	"authentication_with_db/models"
	"authentication_with_db/utils"
	"database/sql"
	"log"
)

// LoginAdmin - For signing in or logging in
func LoginAdmin(givenEmail, givenPass string) (string, bool) {
	var name, email, pass string
	var flag bool
	rows, err := Db.Query(`SELECT admin_name, admin_email, admin_password FROM admins WHERE admin_email = $1`, givenEmail)
	if err != nil {
		flag = false
		log.Fatal("Error - ", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err1 := rows.Scan(&name, &email, &pass); err1 != nil {
			flag = false
			log.Fatal(err)
		}
	}

	if email == givenEmail {
		if utils.CheckPasswordMatch(givenPass, pass) {
			flag = true
			GetValuesForAdminModel(name, givenEmail)
		}
	}

	if flag {
		return name, true
	} else {
		return "", false
	}
}

func GetValuesForAdminModel(name, email string) {
	// Get elements for admin dashboard
	countMap := getEntityCount()
	productTb := getProductsTable()
	userTb := getUsersTable()
	adminTb := getAdminsTb()

	data := models.AdminModel{
		AdminEmail:   email,
		AdminName:    name,
		EntityCounts: countMap,
		ProductsTb:   productTb,
		UsersTb:      userTb,
		AdminsTb:     adminTb,
	}
	models.InitAdminModel(data)
}

// GetEntityCount - Take the counts of the entities like admin, user, and products
func getEntityCount() map[string]string {
	var (
		usersCount, adminsCount, productsCount string
		totalCounts                            = make(map[string]string)
	)

	row := Db.QueryRow(`SELECT COUNT(*) FROM users;`)
	if err := row.Scan(&usersCount); err != nil {
		log.Fatal(err)
	}

	row = Db.QueryRow(`SELECT COUNT(*) FROM admins;`)
	if err := row.Scan(&adminsCount); err != nil {
		log.Fatal(err)
	}

	row = Db.QueryRow(`SELECT COUNT(*) FROM products_tb;`)
	if err := row.Scan(&productsCount); err != nil {
		log.Fatal(err)
	}

	totalCounts["admins"] = adminsCount
	totalCounts["users"] = usersCount
	totalCounts["products"] = productsCount

	return totalCounts
}

func getProductsTable() []models.ProductModel {
	var id, name, price, url string
	row, err := Db.Query(`SELECT * FROM products_tb;`)
	if err != nil {
		log.Fatal(err.Error())
	}

	var res []models.ProductModel
	for row.Next() {
		if err := row.Scan(&id, &name, &price, &url); err != nil {
			log.Fatal(err.Error())
		}

		data := models.ProductModel{
			ProductId:    id,
			ProductName:  name,
			ProductPrice: price,
			ProductUrl:   url,
		}
		res = append(res, data)
	}

	return res
}

func getUsersTable() []models.UserModel {
	var id, name, email string
	var product1, product2, product3, product4, product5 *string

	rows, err := Db.Query(`SELECT user_id, user_name, user_email, product_id_1, product_id_2, product_id_3, product_id_4, product_id_5 FROM users;`)
	if err != nil {
		log.Fatal(err)
	}

	var res []models.UserModel
	for rows.Next() {
		if err := rows.Scan(&id, &name, &email, &product1, &product2, &product3, &product4, &product5); err != nil {
			log.Fatal(err)
		}

		var null string = ""
		// null checking
		if product1 == nil {
			product1 = &null
		}
		if product2 == nil {
			product2 = &null
		}
		if product3 == nil {
			product3 = &null
		}
		if product4 == nil {
			product4 = &null
		}
		if product5 == nil {
			product5 = &null
		}
		// null checking ends here

		data := models.UserModel{
			UserId:       id,
			UserName:     name,
			UserEmail:    email,
			UserProduct1: *product1,
			UserProduct2: *product2,
			UserProduct3: *product3,
			UserProduct4: *product4,
			UserProduct5: *product5,
		}
		res = append(res, data)
	}
	return res
}

func getAdminsTb() []models.AdminsTbContent {
	var id, name, email string

	rows, err := Db.Query(`SELECT admin_id, admin_name, admin_email FROM admins;`)
	if err != nil {
		log.Fatal(err)
	}

	var res []models.AdminsTbContent
	for rows.Next() {
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal(err)
		}
		data := models.AdminsTbContent{
			AdminId:    id,
			AdminName:  name,
			AdminEmail: email,
		}
		res = append(res, data)
	}
	return res
}

func CreateUser(model models.UserModel, password string) sql.Result {
	result, err := Db.Exec(`INSERT INTO users(user_name, user_email, user_password, product_id_1, product_id_2, product_id_3, product_id_4, product_id_5) 
VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, model.UserName, model.UserEmail, password, model.UserProduct1, model.UserProduct2, model.UserProduct3, model.UserProduct4, model.UserProduct5)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Get the user details to update the user data
func GetUserData(id string) models.UserModel {
	var name, email string
	var product1, product2, product3, product4, product5 *string

	row := Db.QueryRow(`SELECT user_name, user_email, product_id_1, product_id_2, product_id_3, product_id_4, product_id_5 
FROM users WHERE user_id = $1;`, id)

	if err := row.Scan(&name, &email, &product1, &product2, &product3, &product4, &product5); err != nil {
		log.Fatal(err.Error())
	}

	var null string = ""
	// null checking
	if product1 == nil {
		product1 = &null
	}
	if product2 == nil {
		product2 = &null
	}
	if product3 == nil {
		product3 = &null
	}
	if product4 == nil {
		product4 = &null
	}
	if product5 == nil {
		product5 = &null
	}
	// null checking ends here

	data := models.UserModel{
		UserId:       id,
		UserName:     name,
		UserEmail:    email,
		UserProduct1: *product1,
		UserProduct2: *product2,
		UserProduct3: *product3,
		UserProduct4: *product4,
		UserProduct5: *product5,
	}

	return data
}

func UpdateUserDatas(id, product1, product2, product3, product4, product5 string) bool {
	_, err := Db.Exec(`UPDATE users SET product_id_1 = $1, product_id_2 = $2, product_id_3 = $3, product_id_4 = $4, product_id_5 = $5 WHERE user_id = $6;`, product1, product2, product3, product4, product5, id)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}

func DeleteUser(id string) {
	_, err := Db.Exec(`DELETE FROM users WHERE user_id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return
}
