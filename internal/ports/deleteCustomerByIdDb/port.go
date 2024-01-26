package deleteCustomerByIdDb

type Port interface {
	Execute(*Request) error
}
type Request struct {
	Id int `gorm:"column:id;primaryKey" json:"id"`
}

type Response struct {
}

func (Request) TableName() string {
	return "tbl_customers"
}
