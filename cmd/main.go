package main

import (
	"questionplatform/boot"
)

func main() {
	boot.InitViper("./config.yaml")
	boot.Loggersetup()
	boot.DatabaseSetUp()
	boot.InitResource("./idiom.sql")
	boot.ServerSetUp()
}
