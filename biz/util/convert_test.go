package util

import (
	"encoding/json"
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"testing"
)

func TestStructToMap(t *testing.T) {
	m := StructToMap(&vo.GetSgzGameZoneItemListReq{
		GameId:           consts.GameId,
		Fcid:             0,
		OsId:             0,
		Cid:              0,
		PlatformId:       0,
		Sort:             "",
		ExtConditions:    "",
		StdCatId:         0,
		JymCatId:         0,
		FilterLowQuality: "",
		Keyword:          "",
		PriceRange:       "",
		Page:             0,
	})
	jsonStr, _ := json.Marshal(m)

	fmt.Printf("%s\n", jsonStr)
}
