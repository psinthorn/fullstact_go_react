package app

import (
	"github.com/psinthorn/fullstack_go_react/backend/api/users/controllers/index"
	"github.com/psinthorn/fullstack_go_react/backend/api/users/controllers/users"
)

func urlsMapping() {

	// Section: Index
	router.GET("/", index.Welcome)
	router.GET("/about", index.About)
	// router.GET("/users", users.GetAll)
	// router.GET("/contents", contents.GetAll)

	// Index: API
	// router.GET("/ping", controllers.Ping.Pong)
	router.POST("/users", users.Create)
	router.GET("/allusers", users.GetAll)
	router.GET("/users/:id", users.Get)
	router.PATCH("/users/:id", users.Update)
	router.PUT("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
