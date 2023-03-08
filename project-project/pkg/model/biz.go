package model

const (
	Normal         = 1
	Personal int32 = 1
)

const AESKey = "sdfgyrhgbxcdgryfhgywertd"

const (
	NoDelete = iota
	NoDeleted
)

const (
	NoArchive = iota
	Archived
)

const (
	Open = iota
	Private
	Custom
)

const (
	Default = "default"
	Simple  = "simple"
)