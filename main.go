package main

import (
	"Todoapp/config"
	"fmt"
)

func main () {
	fmt.Println(config.Config.Port)
}