package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// ##### Server #####
	controllers.StartMainServer()

	// ##### Session #####
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

	// ##### Database #####

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// -- CreateUser
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// -- CreateUser(2)
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// -- GetUser
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// -- UpdateUser
	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// -- DeleteUser
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// -- CreateTodo
	// user, _ := models.GetUser(2) // IDが増分されてるので...
	// user.CreateTodo("First Todo")

	// -- CreateTodo(2)
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// -- GetTodo
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// -- GetTodos
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// -- GetTodosByUser
	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	// t, _ := models.GetTodo(4)
	// t.DeleteTodo()
}
