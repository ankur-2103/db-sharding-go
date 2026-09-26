package main

import (
	"db-sharding/internal/handler"
	"db-sharding/internal/router"
	"db-sharding/internal/services"
	"db-sharding/internal/shard"
	"db-sharding/internal/store"
	"log"
	"net/http"
)

func main() {
	shards := map[int]shard.Shard{
		0: {
			ID:       1,
			Name:     "shard 1",
			Database: "db1",
			Host:     "192.9.9.1",
			Port:     22,
			Store:    store.NewStore(),
		},
		1: {
			ID:       2,
			Name:     "shard 2",
			Database: "db2",
			Host:     "192.9.9.2",
			Port:     22,
			Store:    store.NewStore(),
		},
		2: {
			ID:       3,
			Name:     "shard 3",
			Database: "db3",
			Host:     "192.9.9.3",
			Port:     22,
			Store:    store.NewStore(),
		},
		3: {
			ID:       4,
			Name:     "shard 4",
			Database: "db4",
			Host:     "192.9.9.4",
			Port:     22,
			Store:    store.NewStore(),
		},
	}

	strategy := router.ModuloStrategy{}
	shardRouter := router.NewShardRouter(shards,strategy)
	userService := services.NewUserService(shardRouter)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{userID}", userHandler.GetUser)
	mux.HandleFunc("POST /user", userHandler.Create)
	mux.HandleFunc("PUT /users/{userID}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{userID}", userHandler.DeleteUser)

	log.Println("server running on :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}