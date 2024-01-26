package handler

import (
	"Test-CF/infrastructure"
	"Test-CF/internal/cores/updateCustomerById"
	"github.com/gin-gonic/gin"
)

type updateCustomerByIdHdl struct {
	updateCustomerById updateCustomerById.Service
}

func NewUpdateCustomerByIdHdl(updateCustomerById updateCustomerById.Service) *updateCustomerByIdHdl {
	return &updateCustomerByIdHdl{updateCustomerById: updateCustomerById}
}

func (hdl updateCustomerByIdHdl) HandlerFunc(c *gin.Context) {
	var request updateCustomerById.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}
	err := hdl.updateCustomerById.Execute(request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(200, "Update Success!")
}
