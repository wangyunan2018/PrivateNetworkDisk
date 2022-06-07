package main

import (
	"my-network-disk/models"
	"my-network-disk/routers"
)

func main() {
	models.InitMysql()
	routers.InitRouter()
}
