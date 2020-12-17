package models

type ResourceType int

const (
	// UNKNOWN ResourceType = iota -- check later ho to use same enum values in a package
	APPS ResourceType = iota
	WORKFLOWS
	PROFILES
	TIME_WINDOW
	SCRIPTS
	SENSORS
)
