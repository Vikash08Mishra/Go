package Entity

import (
	"time"

	"github.com/Vikash08Mishra/reqhandlerapi/models"
)

type ResourceActionEntity struct {
	DeviceUuid      string
	UserUuid        string
	ResourceUuid    string
	ResourceType    models.ResourceType
	ActionType      models.ActionType
	OverrideType    models.OverrideType
	Source          string
	Version         string
	Target          string
	Status          string // Initialized, INPROGRESS, COMPLETED, FAILED
	SkipOverride    bool   // For ONE_TIME override when status is completed this will be true as false.
	StatusUpdatedOn time.Time
}
