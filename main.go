package main

import (
	"Todoapp/app/controllers"
	"Todoapp/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()

}
