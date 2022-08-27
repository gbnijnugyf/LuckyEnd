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
	WishNotLight    = 0
	WishHaveLight   = 1
	WishHaveRealize = 2
	WishHaveDelete  = 3
)

const (
	LightWish int = iota
	CancelLight
	DeleteWish
	HaveAchieve
)

const (
	MaxWishCount      = 5
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
