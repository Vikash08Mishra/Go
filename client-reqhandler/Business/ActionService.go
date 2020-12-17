package Business

import (
	"context"
	"time"

	"github.com/Vikash08Mishra/reqhandlerapi/Repository"
	"github.com/Vikash08Mishra/reqhandlerapi/models"
	"github.com/Vikash08Mishra/reqhandlerapi/models/Entity"
)

type ActionService struct {
	actionRepository Repository.IActionRepository
}

func NewActionService(actionRepo Repository.IActionRepository) *ActionService {
	return &ActionService{
		actionRepository: actionRepo,
	}
}

// Gets device actions from database
func (actionService *ActionService) GetActions(
	deviceUuid,
	UserUuid string,
	ctx context.Context) (models.ResourcesActionResponseV1, error) {

	actions, err := actionService.actionRepository.GetActions(deviceUuid, UserUuid, ctx)

	if err != nil {
		return models.ResourcesActionResponseV1{}, err
	}

	resourceActionModels := actionEntityToModel(actions)

	resourceActions := models.ResourcesActionResponseV1{
		Actions: resourceActionModels,
	}

	return resourceActions, nil
}

// Create device actions
func (actionService *ActionService) CreateActions(
	deviceUuid,
	userUuid string,
	overrides models.ResourceOverridesRequestV1,
	ctx context.Context) error {

	actionEntities := []Entity.ResourceActionEntity{}

	for _, override := range overrides.Resources {
		actionEntity := Entity.ResourceActionEntity{
			DeviceUuid:      deviceUuid,
			UserUuid:        userUuid,
			ResourceUuid:    override.ResourceUuid,
			ResourceType:    override.ResourceType,
			ActionType:      override.ActionType,
			OverrideType:    override.OverrideType,
			Source:          override.Source,
			Version:         override.Version,
			Target:          override.Target,
			Status:          "INITIALIZED",
			SkipOverride:    false,
			StatusUpdatedOn: time.Now().UTC(),
		}

		actionEntities = append(actionEntities, actionEntity)
	}

	err := actionService.actionRepository.CreateActions(actionEntities, ctx)

	if err != nil {
		return err
	}

	return nil
}

// Update actions status
func (actionService *ActionService) UpdateActionStatus(
	deviceUuid,
	UserUuid string,
	status models.ResourcesActionStatusV1Model,
	ctx context.Context) error {

	actionEntity := Entity.ResourceActionEntity{
		DeviceUuid:   deviceUuid,
		UserUuid:     UserUuid,
		ResourceUuid: status.ResourceUuid,
		Status:       status.Status,
	}

	err := actionService.actionRepository.UpdateActionStatus(actionEntity, ctx)

	if err != nil {
		return err
	}

	return nil
}

// Delete resource actions
func (actionService *ActionService) DeleteActions(
	deviceUuid,
	userUuid,
	resourceUuid string,
	ctx context.Context) error {

	err := actionService.actionRepository.DeleteAction(deviceUuid, userUuid, resourceUuid, ctx)

	if err != nil {
		return err
	}

	return nil
}

// Parse ResourceActionEntity To ActionModels
func actionEntityToModel(actions []Entity.ResourceActionEntity) []models.ResourceActionV1Model {

	resourceActionModels := []models.ResourceActionV1Model{}

	for _, action := range actions {
		resourceActionModel := models.ResourceActionV1Model{
			ResourceUuid:    action.ResourceUuid,
			ResourceType:    action.ResourceType,
			Source:          action.Source,
			ActionType:      action.ActionType,
			Version:         action.Version,
			Target:          action.Target,
			Status:          action.Status,
			StatusUpdatedOn: action.StatusUpdatedOn,
		}

		resourceActionModels = append(resourceActionModels, resourceActionModel)
	}

	return resourceActionModels
}
