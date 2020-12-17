package main

import (
	"github.com/Vikash08Mishra/reqhandlerapi/Business"
	"github.com/Vikash08Mishra/reqhandlerapi/Repository"
	"github.com/Vikash08Mishra/reqhandlerapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAndStartServer() {

	// Initialize Default gin Web server Engine
	r := gin.Default()

	actionService := Business.NewActionService(Repository.NewActionRepository())
	actionController := controllers.NewActionController(actionService)

	//Setup actions route
	r.GET("/dsm/users/:user-uuid/devices/:device-uuid/resources/actions", actionController.GetActionsForResources)
	r.POST("/dsm/users/:user-uuid/devices/:device-uuid/resources/overrides", actionController.CreateActionsForResources)
	r.POST("/dsm/users/:user-uuid/devices/:device-uuid/resources/action-status", actionController.UpdateActionStatus)
	r.DELETE("/dsm/users/:user-uuid/devices/:device-uuid/resource/:resource-uuid/actions", actionController.DeleteActionsForResources)

	// Start server. Panic if error in starting server
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
