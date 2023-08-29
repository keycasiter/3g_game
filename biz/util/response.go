package util

import (
	"github.com/keycasiter/3g_game/biz/model/common"
	"github.com/keycasiter/3g_game/biz/model/enum"
)

func BuildFailedMeta(errMsg string) *common.Meta {
	return &common.Meta{
		StatusCode: enum.ResponseCode_UnknownError,
		StatusMsg:  errMsg,
	}
}
