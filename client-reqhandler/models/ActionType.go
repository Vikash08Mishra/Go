package models

type ActionType int

const (
	UNKNOWN ActionType = iota
	INSTALL
	REMOVE
	RE_INSTALL
	UPGRADE
)
