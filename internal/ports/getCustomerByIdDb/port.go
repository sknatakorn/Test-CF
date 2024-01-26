package getCustomerByIdDb

type Port interface {
	Execute(id int) (*Response, error)
}

type Response struct {
	Id   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;NOT NULL" json:"name"`
	Age  int    `gorm:"column:age;NOT NULL" json:"age"`
}

func (Response) TableName() string {
	return "tbl_customers"
}
