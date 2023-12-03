package logic

import (
	"context"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"testing"
)

func TestNewAccountCheckLogic(t *testing.T) {
	herosArr := make([]string, 0)

	//关关张
	ggz := "10010,10097,10013"
	herosArr = append(herosArr, ggz)
	//北伐
	//bfq := "10204,10101,10102"
	//herosArr = append(herosArr, bfq)
	//麒麟弓
	qlg := "10068,10064,10017"
	herosArr = append(herosArr, qlg)
	//SP群弓
	//spqg := "10203,10206,10115"
	//herosArr = append(herosArr, spqg)
	//太尉盾
	twd := "10033,10014,10122"
	herosArr = append(herosArr, twd)
	//周太凌
	ztl := "10031,10022,10088"
	herosArr = append(herosArr, ztl)

	ctx := context.Background()
	err := NewAccountCheckLogic(ctx, &vo.AccountCheckReq{
		CheckGoodsUrl: "https://m.jiaoyimao.com/jg1009207-3/1700647130393133.html?shareFrom=CopyUrl",
		//间隔区间
		PriceRange: util.ToJsonString(ctx, []string{"500", "1000"}),
		//指定英雄
		DefiniteHeros: herosArr,
		//红度
		DefiniteStage: "",
		//指定武将是否必须觉醒
		IsDefiniteHeroMustAwake: true,
		//指定武将是否必须开三兵书
		IsDefiniteHeroMustTalent3: true,
		//指定特技
		//MustSpecialTech: []string{"援助", "踩踏"},
		//指定战法
		MustTactic: []string{
			//事件战法
			"抚辑军民", "三势阵", "草船借箭",
			//太尉
			"士别三日", "锋矢阵", "刮骨疗毒", "藤甲兵", "魅惑",
			//父女
			"据水断桥", "箕形阵", "青州兵", "横扫千军", "威谋靡亢", "盛气凌敌",
			//麒麟
			"夺魂挟魄", "杯蛇鬼车", "太平道法", "无当飞军", "八门金锁阵",
			//突击骑
			"虎豹骑", "铁骑驱驰", "当锋摧决", "三势阵", "一骑当千",
			//SP群弓
			"白马义从", "折冲御侮",
		},
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}
