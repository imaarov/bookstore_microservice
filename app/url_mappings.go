package app

import (
	"github.com/imaarov/bookstore_microservice/controllers/ping"
	"github.com/imaarov/bookstore_microservice/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users:user_id", users.GetUser)
	router.GET("/users", users.CreateUser)
}
