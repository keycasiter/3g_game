package util

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestResumeSoldierNum(t *testing.T) {
	//兵力为0
	fmt.Println(ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 0,
	}, 100))
	a, b, c := ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 0,
	}, 100)
	assert.Equal(t, a, int64(0))
	assert.Equal(t, b, int64(0))
	assert.Equal(t, c, int64(0))

	//恢复超过兵力上限
	fmt.Println(ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 5000,
	}, 6000))
	a1, b1, c1 := ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 5000,
	}, 6000)
	assert.Equal(t, a1, int64(5000))
	assert.Equal(t, b1, int64(5000))
	assert.Equal(t, c1, int64(10000))

	//恢复不超过兵力上限
	fmt.Println(ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 5000,
	}, 1000))
	a2, b2, c2 := ResumeSoldierNum(&vo.BattleGeneral{
		SoldierNum: 5000,
	}, 1000)
	assert.Equal(t, a2, int64(1000))
	assert.Equal(t, b2, int64(5000))
	assert.Equal(t, c2, int64(6000))
}
