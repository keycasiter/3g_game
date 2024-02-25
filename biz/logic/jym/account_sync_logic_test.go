package jym

import (
	"context"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"testing"
)

func TestNewAccountSyncContext(t *testing.T) {
	//userName, pwd, ips := util.UseDps()
	//util.UseProxy = true
	//util.ProxyUserName = userName
	//util.ProxyPassword = pwd
	//util.ProxyIpPool = ips
	conf.InitConfig()
	mysql.InitMysql()

	ctx := context.Background()
	err := NewAccountSyncContext(ctx, &vo.AccountSearchReq{
		//区段
		//Keyword: "3000区段",
		//价格区间
		//PriceRange: util.ToJsonString(ctx, []string{"1000", "7000"}),
		//指定英雄
		//DefiniteHeros: herosArr,
		//红度
		DefiniteStage: "",
		//总红度
		DefiniteTotalStage: "",
		//指定武将是否必须觉醒
		IsDefiniteHeroMustAwake: false,
		//指定武将是否必须开三兵书
		IsDefiniteHeroMustTalent3: false,
		//可以跨服，公开
		CrossServerAndPublic: false,
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}
