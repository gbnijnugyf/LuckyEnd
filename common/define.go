package common

import (
	"time"
)

const CodeExpries = -2
const CodeError = -1
const CodeSuccess = 0

const Login = true
const UnLogin = false

const WishHaveDelete = 3
const WishHaveRealize = 2
const WishHaveLight = 1
const WishNotLight = 0

const LightWish = 0
const CancelLight = 1
const DeleteWish = 2
const HaveAchieve = 3

const MaxWishCount = 5
const MaxLightCount = 5
const MaxLightSameCount = 2

var ChinaTime *time.Location