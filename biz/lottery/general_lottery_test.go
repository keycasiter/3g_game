package lottery

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"sort"
	"testing"
)

func TestNewGeneralLotteryContext(t *testing.T) {
	conf.InitConfig()
	mysql.InitMysql()
	ctx := context.Background()

	resp, err := NewGeneralLotteryContext(&GeneralLotteryRequest{
		Ctx:            ctx,
		GeneralLottery: consts.PK_DongRuLeiTing,
		RollTimes:      500,
		Uid:            "ssss",
	}).Run()
	if err != nil {
		fmt.Errorf("err:%v", err)
		t.Fail()
	}
	var list GeneralLotteryInfoSortByHitNum = resp.GeneralLotteryInfoList
	sort.Sort(sort.Reverse(list))

	fmt.Printf("保底次数:%d\n", resp.ProtectedMustHitNum)
	for _, general := range list {
		fmt.Printf("%s , 抽中次数:%d , 抽中占比:%.2f%% , 设置概率:%.2f%%\n", general.GeneralInfo.Name, general.HitNum, general.HitRate*100, general.LotteryRate*100)
	}
}
