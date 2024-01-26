package httpserv

import (
	"Test-CF/handler"
	"Test-CF/infrastructure"
	"Test-CF/internal/cores/createCustomer"
	"Test-CF/internal/cores/deleteCustomerById"
	"Test-CF/internal/cores/getCustomerById"
	"Test-CF/internal/cores/updateCustomerById"
	"Test-CF/internal/ports/createCustomerDb"
	"Test-CF/internal/ports/deleteCustomerByIdDb"
	"Test-CF/internal/ports/getCustomerByIdDb"
	"Test-CF/internal/ports/updateCustomerByIdDb"
	"github.com/gin-gonic/gin"
)

func bindGetCustomerByIdRoute(r *gin.Engine) {
	db := getCustomerByIdDb.NewAdapterDb(infrastructure.Db)
	service := getCustomerById.NewService(db)
	hdl := handler.NewGetCustomerByIdHdl(service)

	r.GET("/customers/:id", hdl.HandlerFunc)
}

func bindDeleteCustomerByIdRoute(r *gin.Engine) {
	db := deleteCustomerByIdDb.NewAdapterDb(infrastructure.Db)
	service := deleteCustomerById.NewService(db)
	hdl := handler.NewDeleteCustomerByIdHdl(service)

	r.DELETE("/customers/:id", hdl.HandlerFunc)
}

func bindCreateCustomerByIdRoute(r *gin.Engine) {
	db := createCustomerDb.NewAdapterDb(infrastructure.Db)
	service := createCustomer.NewService(db)
	hdl := handler.NewCreateCustomerByIdHdl(service)

	r.POST("/customers", hdl.HandlerFunc)
}

func bindUpdateCustomerByIdRoute(r *gin.Engine) {
	db := updateCustomerByIdDb.NewAdapterDb(infrastructure.Db)
	service := updateCustomerById.NewService(db)
	hdl := handler.NewUpdateCustomerByIdHdl(service)

	r.PUT("/customers", hdl.HandlerFunc)
}
