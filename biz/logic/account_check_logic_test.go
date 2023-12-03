package logic

import (
	"context"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"strings"
	"testing"
)

func TestNewAccountCheckLogic(t *testing.T) {
	herosArr := make([]string, 0)

	//关关张
	ggz := "10010,10097,10013"
	herosArr = append(herosArr, strings.Split(ggz, ",")...)
	//北伐
	bfq := "10204,10101,10102"
	herosArr = append(herosArr, strings.Split(bfq, ",")...)
	//麒麟弓
	qlg := "10068,10064,10017"
	herosArr = append(herosArr, strings.Split(qlg, ",")...)
	//SP群弓
	spqg := "10203,10206,10115"
	herosArr = append(herosArr, strings.Split(spqg, ",")...)
	//太尉盾
	twd := "10033,10014,10122"
	herosArr = append(herosArr, strings.Split(twd, ",")...)
	//周太凌
	ztl := "10031,10022,10088"
	herosArr = append(herosArr, strings.Split(ztl, ",")...)

	ctx := context.Background()
	err := NewAccountCheckLogic(ctx, &vo.AccountCheckReq{
		CheckGoodsUrl: "https://m.jiaoyimao.com/jg1009207/1701602891730957.html",
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
			//常用必备战法
			"所向披靡", "破阵摧坚", "百骑劫营", "暂避其锋", "兵无常势", "陷阵营", "西凉铁骑", "伪书相间", "用武通神", "万箭齐发",
			"象兵", "锦帆军", "破军威胜", "速乘其利", "竭力佐谋", "火炽原燎", "裸衣血战",
			//进阶战法
			"击其惰归", "掣刀斫敌", "刚勇无前", "婴城自守",
			//事件战法
			"抚辑军民", "三势阵", "草船借箭",
			//太尉
			"士别三日", "熯天炽地", "锋矢阵", "刮骨疗毒", "藤甲兵", "魅惑",
			//父女
			"据水断桥", "箕形阵", "青州兵", "横扫千军", "威谋靡亢", "盛气凌敌",
			//麒麟
			"夺魂挟魄", "杯蛇鬼车", "太平道法", "无当飞军", "八门金锁阵",
			//突击骑
			"虎豹骑", "铁骑驱驰", "当锋摧决", "三势阵", "一骑当千",
			//SP群弓
			"焰逐风飞", "白马义从", "折冲御侮",
		},
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}
