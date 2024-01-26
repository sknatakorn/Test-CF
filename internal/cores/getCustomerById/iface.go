package getCustomerById

type Service interface {
	Execute(Request) (*Response, error)
}

type Request struct {
	Id int ` json:"id"`
}

type Response struct {
	Id   int    ` json:"id"`
	Name string ` json:"name"`
	Age  int    `json:"age"`
}
