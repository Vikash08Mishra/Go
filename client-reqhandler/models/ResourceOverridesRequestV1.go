package models

type ResourceOverridesRequestV1 struct {
	Resources []ResourceOverrideV1Model `json:"resources"`
}

type ResourceOverrideV1Model struct {
	ResourceUuid string
	ResourceType ResourceType
	Source       string
	ActionType   ActionType
	OverrideType OverrideType
	Version      string
	Target       string
}
