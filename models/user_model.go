package models

type UserModel struct {
	PageTitle    string
	UserId       string
	UserName     string
	UserEmail    string
	UserProduct1 string
	UserProduct2 string
	UserProduct3 string
	UserProduct4 string
	UserProduct5 string
}

var userVal *UserModel

func InitUserModel(model UserModel) *UserModel {

	userVal = &UserModel{
		PageTitle:    "User Dashboard",
		UserId:       model.UserId,
		UserName:     model.UserName,
		UserEmail:    model.UserEmail,
		UserProduct1: model.UserProduct1,
		UserProduct2: model.UserProduct2,
		UserProduct3: model.UserProduct3,
		UserProduct4: model.UserProduct4,
		UserProduct5: model.UserProduct5,
	}

	return userVal
}

func ReturnUserModel() *UserModel {
	return userVal
}
