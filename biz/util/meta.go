package util

import "github.com/keycasiter/3g_game/biz/model/common"

func BuildSuccMeta() *common.Meta {
	return &common.Meta{
		StatusCode: 0,
		StatusMsg:  "成功",
	}
}
