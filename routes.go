package main

import (
	"github.com/julienschmidt/httprouter"
)

func CreateNewRouter(handlers *Handlers) *httprouter.Router {
	router := httprouter.New()

	router.POST("/set_relation", handlers.uController.setUserRelations)
	router.POST("/list_following", handlers.uController.listAllFollowingConnections)
	router.POST("/list_followed", handlers.uController.listAllFollowedConnections)
	router.POST("/list_connections_stat", handlers.uController.listConnectionsStats)
	router.POST("/add_favorite", handlers.uController.addToFavorites)
	router.POST("/remove_favorite", handlers.uController.removeFromFavorites)
	router.POST("/list_favorites", handlers.uController.listFavorites)

	return router
}
