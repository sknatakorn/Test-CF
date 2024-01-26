package main

import (
	"Test-CF/httpserv"
	"Test-CF/infrastructure"
)

func main() {
	infrastructure.InitConfig()
	infrastructure.InitDb()
	httpserv.Run()
}
