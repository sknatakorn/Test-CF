package createCustomerDb

type Port interface {
	Execute(req *Request) error
}
type Request struct {
	Id   int    ` json:"id"`
	Name string ` json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
}

func (Request) TableName() string {
	return "tbl_customers"
}
