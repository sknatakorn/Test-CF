package httpserv

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	r := gin.Default()

	bindGetCustomerByIdRoute(r)
	bindCreateCustomerByIdRoute(r)
	bindDeleteCustomerByIdRoute(r)
	bindUpdateCustomerByIdRoute(r)

	appPort := viper.GetString("app.port")
	if appPort == "" {
		appPort = "8080"
	}

	r.Run(":" + appPort)

}
