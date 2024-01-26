package handler

import (
	"Test-CF/infrastructure"
	"Test-CF/internal/cores/getCustomerById"
	"github.com/gin-gonic/gin"
	"strconv"
)

type getCustomerByIdHdl struct {
	getCustomerById getCustomerById.Service
}

func NewGetCustomerByIdHdl(getCustomerById getCustomerById.Service) *getCustomerByIdHdl {
	return &getCustomerByIdHdl{getCustomerById: getCustomerById}
}

func (hdl getCustomerByIdHdl) HandlerFunc(c *gin.Context) {
	var request getCustomerById.Request
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}
	request.Id = id

	res, err := hdl.getCustomerById.Execute(request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(infrastructure.GetCodeError(err), gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(200, res)
}
