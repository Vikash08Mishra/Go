package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Vikash08Mishra/reqhandlerapi/models"

	"github.com/Vikash08Mishra/reqhandlerapi/Business"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ActionController struct {
	actionService *Business.ActionService
}

func NewActionController(actionService *Business.ActionService) *ActionController {
	return &ActionController{
		actionService: actionService,
	}
}

// Gets existing Device Actions for resources
func (a *ActionController) GetActionsForResources(c *gin.Context) {

	userUuid, deviceUuid, _, valid := validateInputRequest(c)

	if !valid {
		fmt.Printf("Invalid request parameters %v, %v", userUuid, deviceUuid)
		return
	}

	ctx := context.Background()
	actions, err := a.actionService.GetActions(deviceUuid, userUuid, ctx)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(actions.Actions) < 1 {
		c.JSON(http.StatusNoContent, actions)
		return
	}

	c.JSON(http.StatusOK, actions)
	return
}

// Creates device actions for resources.
func (a *ActionController) CreateActionsForResources(c *gin.Context) {

	userUuid, deviceUuid, _, valid := validateInputRequest(c)

	if !valid {
		fmt.Printf("Invalid request parameters %v, %v", userUuid, deviceUuid)
		return
	}

	var request models.ResourceOverridesRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	err := a.actionService.CreateActions(deviceUuid, userUuid, request, ctx)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, request)

	return
}

// Update Action Status endpoint.
func (a *ActionController) UpdateActionStatus(c *gin.Context) {

	userUuid, deviceUuid, _, valid := validateInputRequest(c)

	if !valid {
		fmt.Printf("Invalid request parameters %v, %v", userUuid, deviceUuid)
		return
	}

	var request models.ResourcesActionStatusRequestV1

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	err := a.actionService.UpdateActionStatus(deviceUuid, userUuid, request.Actions[0], ctx)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)

	return
}

// Deletes Device actions endpoint.
func (a *ActionController) DeleteActionsForResources(c *gin.Context) {

	userUuid, deviceUuid, resourceUuid, valid := validateInputRequest(c)

	if !valid {
		fmt.Printf("Invalid request parameters %v, %v, %v", userUuid, deviceUuid, resourceUuid)
		return
	}

	ctx := context.Background()
	err := a.actionService.DeleteActions(deviceUuid, userUuid, resourceUuid, ctx)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)

	return
}

// Validates Input request and parameters
func validateInputRequest(c *gin.Context) (string, string, string, bool) {

	userUuid := c.Param("user-uuid")
	deviceUuid := c.Param("device-uuid")
	resourceUuid := c.Param("resource-uuid")

	_, err := uuid.Parse(userUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"UserUuid Must be passed and should be of form Guid": err.Error()})
		return userUuid, deviceUuid, resourceUuid, false
	}

	_, err = uuid.Parse(deviceUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"DeviceUuid Must be passed and should be of form Guid": err.Error()})
		return userUuid, deviceUuid, resourceUuid, false
	}

	if resourceUuid != "" {
		_, err = uuid.Parse(resourceUuid)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"resourceUuid Must be of form Guid": err.Error()})
			return userUuid, deviceUuid, resourceUuid, false
		}
	}

	return userUuid, deviceUuid, resourceUuid, true
}
