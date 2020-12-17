package Repository

import (
	"context"

	"github.com/Vikash08Mishra/reqhandlerapi/models/Entity"
)

type IActionRepository interface {
	GetActions(deviceUuid, userUuid string, ctx context.Context) ([]Entity.ResourceActionEntity, error)

	CreateActions(actions []Entity.ResourceActionEntity, ctx context.Context) error

	UpdateActionStatus(action Entity.ResourceActionEntity, ctx context.Context) error

	DeleteAction(deviceUuid, userUuid, resourceUuid string, ctx context.Context) error
}
