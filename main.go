package main

import (
	"Todoapp/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
/*
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testetst"
	fmt.Println(u)

	u.CretateUser()
*/

u,_ := models.GetUser(1)
fmt.Println(u)

}
