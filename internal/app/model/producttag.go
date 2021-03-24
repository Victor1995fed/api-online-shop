package model

//Product ...
type ProductTag struct {
	ID        int
	ProductId int
	TagId     int
}

//GetTableName ...
func (p *ProductTag) GetTableName() string {
	return "product_tag"
}
