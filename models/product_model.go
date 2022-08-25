package models

type ProductModel struct {
	ProductId    string
	ProductName  string
	ProductPrice string
	ProductUrl   string
}

var productVal *ProductModel

func InitProductModel(model ProductModel) *ProductModel {

	productVal = &ProductModel{
		ProductId:    model.ProductId,
		ProductName:  model.ProductName,
		ProductPrice: model.ProductPrice,
		ProductUrl:   model.ProductUrl,
	}

	return productVal
}

func ReturnProductModel() *ProductModel {
	return productVal
}
