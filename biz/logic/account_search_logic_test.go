package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
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
	doc, err := goquery.NewDocument("https://m.jiaoyimao.com/jg1009207/1700896292454925.html?spm=gcmall.goods_listpage.detail.0")
	if err != nil {
		t.Errorf("err:%v", err)
		t.Failed()
	}
	for _, node := range doc.Nodes {
		fmt.Printf(node.Data)
	}
}
