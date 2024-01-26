package updateCustomerById

type Service interface {
	Execute(Request) error
}

type Request struct {
	Id   int    ` json:"id"`
	Name string ` json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
}
