package main

import (
	"Todoapp/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	/*
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.PassWord = "testetst2"
		fmt.Println(u)

		u.CretateUser()
	*/

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u,_ = models.GetUser(1)
		fmt.Println(u) */
	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/
	/*
		user, _ := models.GetUser(4)
		user.CreateTodo("forth Todo")

		/*
			todos,_ := models.GetTodos()
			for _,v := range todos {
				fmt.Println(v)
			} */
	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser(2)
		for _, v := range todos {
			fmt.Println(v)
		} */

	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
