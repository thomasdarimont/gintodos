package api

import (
	"gintodos/data"
	"gintodos/model"
	"gintodos/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoApi struct {
	service *model.TodoService
}

func RegisterTodoEndpoints(router *gin.Engine) {

	api := &TodoApi{
		service: model.NewTodoService(store.NewMemoryTodoStore(), data.NewExampleDataInitializer()),
	}
	api.registerRoutes(router)
}

func (a *TodoApi) registerRoutes(router *gin.Engine) {
	router.GET("/todos", a.findAllTodos)
	router.GET("/todos/:id", a.getTodoByID)
	router.PATCH("/todos/:id", a.updateTodoByID)
	router.POST("/todos", a.newTodo)
}

//
// curl -v http://localhost:9090/todos
func (a *TodoApi) findAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, a.service.FindAll())
}

//
// curl -v http://localhost:9090/todos/1
func (a *TodoApi) getTodoByID(context *gin.Context) {
	id := context.Param("id")
	todo, err := a.service.FindById(id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

//
// curl -v -H "Content-Type: application/json" -d '{"item":"Learn go"}' http://localhost:9090/todos
func (a *TodoApi) newTodo(context *gin.Context) {

	var newTodo model.TodoNew
	if err := context.BindJSON(&newTodo); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(newTodo.Item) == 0 {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todo, err := a.service.CreateTodo(newTodo)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.IndentedJSON(http.StatusCreated, todo)
}

//
// curl -v -X PATCH -H "Content-Type: application/json" -d '{"item":"Item 1 Hack some go", "Completed": true}' http://localhost:9090/todos/1
// curl -v -X PATCH -H "Content-Type: application/json" -d '{"item":"Item 1 Additional hacks"}' http://localhost:9090/todos/1
// curl -v -X PATCH -H "Content-Type: application/json" -d '{"Completed": true}' http://localhost:9090/todos/2
func (a *TodoApi) updateTodoByID(context *gin.Context) {

	var update *model.TodoUpdate
	if err := context.BindJSON(&update); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id := context.Param("id")

	found, updated, err := a.service.UpdateTodo(id, update)

	if !found {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	if updated == nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.IndentedJSON(http.StatusOK, updated)
}
