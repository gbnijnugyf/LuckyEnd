package common

import (
	"time"
)

const (
	CodeExpires = -2
	CodeError   = -1
	CodeSuccess = 0
)

const (
	DesireNotLight    = 0
	DesireHaveLight   = 1
	DesireHaveRealize = 2
	DesireHaveDelete  = 3
)

const (
	LightDesire int = iota
	CancelLight
	DeleteDesire
	HaveAchieve
)

const (
	MaxDesireCount    = 5
	MaxLightCount     = 5
	MaxLightSameCount = 2
)

const GetCountError = -1

var ChinaTime *time.Location

type Gender int

const (
	Male Gender = iota + 1
	FeMale
)
