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

//指定武将
func TestAccountSearchLogicForHero(t *testing.T) {
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
	err := NewAccountSearchContext(ctx, &vo.AccountSearchReq{
		//间隔区间
		PriceRange: util.ToJsonString(ctx, []string{"500", "1000"}),
		//指定英雄
		DefiniteHeros: herosArr,
		//红度
		DefiniteStage: "",
		//当前查询页数，每页最多15个
		PageSize: 10,
		//指定武将是否必须觉醒
		IsDefiniteHeroMustAwake: true,
		//指定武将是否必须开三兵书
		IsDefiniteHeroMustTalent3: true,
		//指定特技
		MustSpecialTech: []string{"援助", "踩踏"},
		//指定战法
		MustTactic: []string{
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
			//"焰逐风飞", "白马义从", "折冲御侮",
		},
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}
