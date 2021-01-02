package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID"`
	ProductName  string `json:"ProductName" sql:"productName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum"`
	ProductImage string `json:"ProductImage" sql:"productImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl"`
}
