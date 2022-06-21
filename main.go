package main

import (
	"Todoapp/config"
	"fmt"
	"log"
)

func main () {
	fmt.Println(config.Config.Port)
	log.Println("test")
}