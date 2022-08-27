package database

import (
	"authentication_with_db/models"
	"authentication_with_db/utils"
	"log"
)

func RegisterUser(formName, formEmail, formPassword string) bool {

	encryptedFormPassword, err := utils.HashEncrypt(formPassword)
	if err != nil {
		log.Fatal("Encryption error : ", err)
		return false
	}

	_, err1 := Db.Exec(`INSERT INTO users(user_name, user_email, user_password) VALUES($1, $2, $3);`, formName, formEmail, encryptedFormPassword)
	if err1 != nil {
		log.Fatal(err1)
		return false
	}

	return true
}

func LoginUser(formEmail, formPassword string) bool {
	var name, email, pass string
	var product1, product2, product3, product4, product5 *string
	var flag bool
	rows, err := Db.Query(
		`SELECT 
    	user_name, 
    	user_email, 
    	user_password,
    	product_id_1, 
    	product_id_2, 
    	product_id_3, 
    	product_id_4, 
    	product_id_5 
	FROM users 
	WHERE user_email = $1;`, formEmail)
	if err != nil {
		flag = false
		log.Fatal("Error - ", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err1 := rows.Scan(
			&name,
			&email,
			&pass,
			&product1,
			&product2,
			&product3,
			&product4,
			&product5); err1 != nil {
			flag = false
			log.Fatal("Error - ", err1)
		}
	}

	if email == formEmail {
		if utils.CheckPasswordMatch(formPassword, pass) == true {
			flag = true
			var user models.UserModel
			if product1 == nil {
				user = models.UserModel{
					UserName:  name,
					UserEmail: email,
				}
			} else {
				user = models.UserModel{
					UserName:     name,
					UserEmail:    email,
					UserProduct1: *product1,
					UserProduct2: *product2,
					UserProduct3: *product3,
					UserProduct4: *product4,
					UserProduct5: *product5,
				}
			}

			models.InitUserModel(user)
		}
	}

	if flag {
		return true
	} else {
		return false
	}
}
