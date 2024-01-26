package handler

import (
	"Test-CF/infrastructure"
	"Test-CF/internal/cores/deleteCustomerById"
	"github.com/gin-gonic/gin"
	"strconv"
)

type deleteCustomerByIdHdl struct {
	deleteCustomerById deleteCustomerById.Service
}

func NewDeleteCustomerByIdHdl(deleteCustomerById deleteCustomerById.Service) *deleteCustomerByIdHdl {
	return &deleteCustomerByIdHdl{deleteCustomerById: deleteCustomerById}
}

func (hdl deleteCustomerByIdHdl) HandlerFunc(c *gin.Context) {
	var request deleteCustomerById.Request
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}
	request.Id = id

	err = hdl.deleteCustomerById.Execute(request)
	if err != nil {
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, "Delete Success")
}
