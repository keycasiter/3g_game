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
	err := NewAccountSearchContext(ctx, &vo.GetSgzGameZoneItemListReq{
		PriceRange: util.ToJsonString(ctx, []string{"500", "1000"}),
	}).Process()
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
}
