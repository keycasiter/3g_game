package util

import (
	"encoding/json"

	"github.com/keycasiter/3g_game/biz/model/vo"
)

func DeepCopyBattleGenerals(src []*vo.BattleGeneral) ([]*vo.BattleGeneral, error) {
	arr := make([]*vo.BattleGeneral, 0)
	bytes, err := json.Marshal(src)
	if err != nil {
		return arr, err
	}
	err = json.Unmarshal(bytes, &arr)
	if err != nil {
		return arr, err
	}
	return arr, nil
}
