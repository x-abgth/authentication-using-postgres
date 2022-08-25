package models

type AdminsTbContent struct {
	AdminId    string
	AdminName  string
	AdminEmail string
}

type AdminModel struct {
	PageTitle    string
	AdminName    string
	AdminEmail   string
	EntityCounts map[string]string
	ProductsTb   []ProductModel
	UsersTb      []UserModel
	AdminsTb     []AdminsTbContent
}

var adminVal *AdminModel

func InitAdminModel(model AdminModel) *AdminModel {

	adminVal = &AdminModel{
		PageTitle:    "Admin Dashboard",
		AdminName:    model.AdminName,
		AdminEmail:   model.AdminEmail,
		EntityCounts: model.EntityCounts,
		ProductsTb:   model.ProductsTb,
		UsersTb:      model.UsersTb,
		AdminsTb:     model.AdminsTb,
	}

	return adminVal
}

func ReturnAdminModel() *AdminModel {
	return adminVal
}
