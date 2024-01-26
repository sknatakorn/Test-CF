package handler

import (
	"Test-CF/infrastructure"
	"Test-CF/internal/cores/createCustomer"
	"github.com/gin-gonic/gin"
)

type createCustomerHdl struct {
	createCustomer createCustomer.Service
}

func NewCreateCustomerByIdHdl(createCustomer createCustomer.Service) *createCustomerHdl {
	return &createCustomerHdl{createCustomer: createCustomer}
}

func (hdl createCustomerHdl) HandlerFunc(c *gin.Context) {
	var request createCustomer.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})

		return
	}
	err := hdl.createCustomer.Execute(request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, "Create Success!")
}
