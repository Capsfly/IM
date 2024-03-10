package main

import (
	"awesomeProject/router"
	"awesomeProject/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
