package Repository

import (
	"context"
	"sync"
	"time"

	"github.com/Vikash08Mishra/reqhandlerapi/models/Entity"
)

type ActionRepository struct {
	lock    *sync.RWMutex
	actions []Entity.ResourceActionEntity
}

func NewActionRepository() *ActionRepository {

	return &ActionRepository{
		lock:    &sync.RWMutex{},
		actions: []Entity.ResourceActionEntity{},
	}
}

// Gets actions from database
func (a *ActionRepository) GetActions(deviceUuid, userUuid string, ctx context.Context) ([]Entity.ResourceActionEntity, error) {

	a.lock.RLock()
	defer a.lock.RUnlock()

	actions := []Entity.ResourceActionEntity{}
	for _, action := range a.actions {
		if action.DeviceUuid == deviceUuid && action.UserUuid == userUuid {
			actions = append(actions, action)
		}
	}

	return actions, nil
}

// Created/Update actions into database
func (a *ActionRepository) CreateActions(actions []Entity.ResourceActionEntity, ctx context.Context) error {

	a.lock.Lock()
	defer a.lock.Unlock()

	a.actions = append(a.actions, actions...)
	return nil
}

// Update actions status
func (a *ActionRepository) UpdateActionStatus(action Entity.ResourceActionEntity, ctx context.Context) error {

	a.lock.Lock()
	defer a.lock.Unlock()

	i := 0
	currentAction := Entity.ResourceActionEntity{}
	for ; i < len(a.actions)-1; i++ {
		currentAction = a.actions[i]
		if currentAction.DeviceUuid == action.DeviceUuid && currentAction.UserUuid == action.UserUuid && currentAction.ResourceUuid == action.ResourceUuid {
			break
		}
	}

	currentAction.ActionType = action.ActionType
	currentAction.Status = action.Status
	currentAction.StatusUpdatedOn = time.Now().UTC()
	if action.Status == "COMPLETED" || action.Status == "FAILED" {
		currentAction.SkipOverride = true
	}

	a.actions[i] = currentAction

	return nil
}

// Deletes actions from database
func (a *ActionRepository) DeleteAction(deviceUuid, userUuid, resourceUuid string, ctx context.Context) error {

	a.lock.Lock()
	defer a.lock.Unlock()

	i := 0
	for ; i < len(a.actions)-1; i++ {
		action := a.actions[i]
		if action.DeviceUuid == deviceUuid && action.UserUuid == userUuid && action.ResourceUuid == resourceUuid {
			break
		}
	}

	a.actions[i] = a.actions[len(a.actions)-1]

	a.actions = a.actions[:len(a.actions)-1]

	return nil
}
