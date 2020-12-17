package models

import (
	"time"
)

type ResourcesActionResponseV1 struct {
	Actions []ResourceActionV1Model
}

type ResourceActionV1Model struct {
	ResourceUuid    string
	ResourceType    ResourceType
	Source          string
	ActionType      ActionType
	Version         string
	Target          string
	Status          string // Initialized, INPROGRESS, COMPLETED, FAILED
	StatusUpdatedOn time.Time
}
