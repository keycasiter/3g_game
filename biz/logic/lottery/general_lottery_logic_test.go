package lottery

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/cache"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/spf13/cast"
	"sort"
	"testing"
)

func TestNewGeneralLotteryContext(t *testing.T) {
	conf.InitConfig()
	mysql.InitMysql()
	cache.InitCache()
	ctx := context.Background()
	rollTimes := int64(5)

	resp, err := NewGeneralLotteryLogic(ctx, &vo.GeneralLotteryRequest{
		GeneralLottery: consts.PKSaiJiZuiZhongDaKaChi,
		RollTimes:      rollTimes,
		Uid:            "ssss",
	}).Run()
	if err != nil {
		fmt.Errorf("err:%v", err)
		t.Fail()
	}
	var list vo.GeneralLotteryInfoSortByHitNum = resp.GeneralLotteryInfoList
	sort.Sort(sort.Reverse(list))

	fmt.Printf("保底次数:%d\n", resp.ProtectedMustHitNum)
	fmt.Printf("出橙次数:%d\n", resp.Hit5LevGeneralNum)
	fmt.Printf("出橙率:%.2f%%\n", cast.ToFloat64(rollTimes)/cast.ToFloat64(resp.Hit5LevGeneralNum))
	for _, general := range list {
		fmt.Printf("%s , 抽中次数:%d , 抽中占比:%.2f%% , 设置概率:%.2f%%\n", general.GeneralInfo.Name, general.HitNum, general.HitRate*100, general.LotteryRate*100)
	}
}
