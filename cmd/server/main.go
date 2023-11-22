package main

import (
	"github.com/renatocantarino/go/APIS/configs"
)

func main() {
	config, _ := configs.Load(".")
	println(config.DbDriver)

}
