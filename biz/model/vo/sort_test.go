package vo

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/util"
	"sort"
	"testing"
)

func TestSortSpeed(t *testing.T) {
	var vos BattleGeneralsOrderBySpeed
	for i := 0; i < 10; i++ {
		vos = append(vos, &BattleGeneral{
			BaseInfo: &po.MetadataGeneral{
				AbilityAttr: &po.AbilityAttr{
					SpeedBase: util.Random(0, 100),
				},
			},
		})
	}
	sort.Sort(vos)
	for _, vo := range vos {
		fmt.Printf("%f\n", vo.BaseInfo.AbilityAttr.SpeedBase)
	}
	//95.419952
	//77.509331
	//65.271926
	//53.067911
	//52.946387
	//48.101298
	//34.253297
	//23.188408
	//22.932015
	//14.362389
}
