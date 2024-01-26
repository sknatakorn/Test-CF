package updateCustomerByIdDb

type Port interface {
	Execute(*Request) error
}
type Request struct {
	Id   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;NOT NULL" json:"name"`
	Age  int    `gorm:"column:age;NOT NULL" json:"age"`
}

type Response struct {
}

func (Request) TableName() string {
	return "tbl_customers"
}
