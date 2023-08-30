package lottery

import (
	"context"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
)

type GeneralLotteryRequest struct {
	Ctx context.Context
	//抽取卡池枚举
	GeneralLottery consts.GeneralLotteryPool
	//抽取次数
	RollTimes int64
	//用户uid
	Uid string
}

type GeneralLotteryResponse struct {
	//抽取的武将信息
	GeneralLotteryInfoList []*GeneralLotteryInfo
	//保底统计
	ProtectedMustHitNum int64
	//五星武将出现率
	Hit5LevGeneralNum int64
}

type GeneralLotteryInfo struct {
	//武将信息
	GeneralInfo *po.General
	//抽中次数
	HitNum int64
	//本次抽中占比
	HitRate float64
	//游戏设置概率
	LotteryRate float64
}

//按出现次数排序
type GeneralLotteryInfoSortByHitNum []*GeneralLotteryInfo

func (arr GeneralLotteryInfoSortByHitNum) Len() int {
	return len(arr)
}
func (arr GeneralLotteryInfoSortByHitNum) Less(i, j int) bool {
	return arr[i].HitNum < arr[j].HitNum
}

func (arr GeneralLotteryInfoSortByHitNum) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
