package util

import (
	"github.com/keycasiter/3g_game/biz/model/common"
	"github.com/keycasiter/3g_game/biz/model/enum"
)

func BuildSuccMeta() *common.Meta {
	return &common.Meta{
		StatusCode: enum.ResponseCode_Success,
		StatusMsg:  "成功",
	}
}

func BuildFailMeta() *common.Meta {
	return &common.Meta{
		StatusCode: enum.ResponseCode_UnknownError,
		StatusMsg:  "失败",
	}
}
