package models

type ResourcesActionStatusRequestV1 struct {
	Actions []ResourcesActionStatusV1Model
}

type ResourcesActionStatusV1Model struct {
	ResourceUuid string
	ActionType   ActionType
	Version      string
	Target       string
	Status       string // Initialized, INPROGRESS, COMPLETED, FAILED
}
