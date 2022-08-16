package common

import (
	"time"
)

const CodeExpires = -2
const CodeError = -1
const CodeSuccess = 0

const WishNotLight = 0
const WishHaveLight = 1
const WishHaveRealize = 2
const WishHaveDelete = 3

const LightWish = 0
const CancelLight = 1
const DeleteWish = 2
const HaveAchieve = 3

const MaxWishCount = 5
const MaxLightCount = 5
const MaxLightSameCount = 2

const GetCountError = -1

const NoExist = "400 Bad Request"
const PasError = "401 Unauthorized"
const LoginSuccess = "200 OK"

var Login bool
var ChinaTime *time.Location
