package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"regexp"
	"strings"
	"testing"
)

func TestSearchAccount(t *testing.T) {
	ctx := context.Background()
	req := &vo.GetSgzGameZoneItemListReq{
		GameId:     consts.GameId,
		Fcid:       consts.FcId,
		OsId:       consts.OsId,
		Cid:        consts.CId,
		PlatformId: consts.PlatformId,
		Sort:       consts.Sort,
		ExtConditions: util.ToJsonString(ctx, &vo.ExtConditions{
			Stage: "",
			Hero:  "",
		}),
		StdCatId:         consts.StdCatId,
		JymCatId:         consts.JymCatId,
		FilterLowQuality: consts.FilterLowQuality,
		Keyword:          "",
		PriceRange:       util.ToJsonString(ctx, []string{"500", "1000"}),
		Page:             1,
	}
	httpResp, err := util.HttpGet(ctx, consts.Url_GetSgzGameZoneItemList, nil, util.StructToMap(req))
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	fmt.Printf("\n:%s", httpResp)
	resp := &vo.GetSgzGameZoneItemListResp{}
	err = json.Unmarshal([]byte(httpResp), &resp)
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	fmt.Printf("\n:%v", resp)
}

func TestParseHtml(t *testing.T) {
	httpRes, err := util.HttpGet(context.Background(), "https://m.jiaoyimao.com/jg1009207/1701412803136059.html", nil, nil)
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	//fmt.Printf("httpRes:\n%s", httpRes)
	reg, err := regexp.Compile("window.__INITIAL_STATE__ =.*\\n")
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	jsonStr := reg.FindString(httpRes)
	jsonStr = strings.ReplaceAll(jsonStr, "window.__INITIAL_STATE__ =", "")
	//fmt.Printf("jsonStr:\n%s", jsonStr)

	data := &vo.AccountItemInfo{}
	err = json.Unmarshal([]byte(jsonStr), data)
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	for _, hero := range data.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
		fmt.Printf("%d=%s\n", hero.HeroId, hero.Name)
	}
}

func TestAccountSearchLogic(t *testing.T) {
	ctx := context.Background()
	err := NewAccountSearchContext(ctx, &vo.AccountSearchReq{}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}

// 指定武将
func TestAccountSearchLogicForHero(t *testing.T) {
	herosArr := make([]string, 0)

	//关关张
	//ggz := "10010,10097,10013"
	//herosArr = append(herosArr, strings.Split(ggz, ",")...)
	//北伐
	//bfq := "10204,10101,10102"
	//herosArr = append(herosArr, strings.Split(bfq, ",")...)
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
	//ztl := "10031,10022,10088"
	//herosArr = append(herosArr, strings.Split(ztl, ",")...)
	//三仙
	//sxd := "10044,10043,10030"
	//herosArr = append(herosArr, strings.Split(sxd, ",")...)
	//渊骑
	//yq := "10056,10014,10024"
	//herosArr = append(herosArr, strings.Split(yq, ",")...)

	ctx := context.Background()
	err := NewAccountSearchContext(ctx, &vo.AccountSearchReq{
		//区段
		Keyword: "4000区段",
		//价格区间
		PriceRange: util.ToJsonString(ctx, []string{"1300", "2000"}),
		//指定英雄
		DefiniteHeros: herosArr,
		//红度
		DefiniteStage: "",
		//总红度
		DefiniteTotalStage: "6",
		//当前查询页数，每页最多15个
		PageSize: 3,
		//指定武将是否必须觉醒
		IsDefiniteHeroMustAwake: false,
		//指定武将是否必须开三兵书
		IsDefiniteHeroMustTalent3: false,
		//指定特技
		//MustSpecialTech: []string{"援助"},
		//指定战法
		MustTactic: []string{
			//常用必备战法
			"所向披靡",
			"破阵摧坚",
			"百骑劫营",
			"暂避其锋",
			"兵无常势",
			"陷阵营",
			"西凉铁骑",
			"伪书相间",
			"用武通神",
			"万箭齐发",
			"象兵",
			"锦帆军",
			"破军威胜",
			"速乘其利",
			"竭力佐谋",
			"火炽原燎",
			"裸衣血战",
			//进阶战法
			"抚辑军民", "三势阵", "草船借箭", "裸衣血战",
			//太尉
			"士别三日", "熯天炽地", "锋矢阵", "刮骨疗毒", "藤甲兵", "魅惑",
			//父女
			"据水断桥", "箕形阵", "青州兵", "横扫千军", "威谋靡亢", "盛气凌敌",
			//麒麟
			"夺魂挟魄", "杯蛇鬼车", "太平道法", "无当飞军", "八门金锁阵",
			"婴城自守",
			//突击骑
			"虎豹骑", "铁骑驱驰", "当锋摧决", "三势阵", "一骑当千", "百骑劫营",
			//SP群弓
			"焰逐风飞", "白马义从", "折冲御侮", "掣刀斫敌",
		},
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}