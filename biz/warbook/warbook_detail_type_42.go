package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 诱敌之策
// 自身受到兵刃伤害和谋略伤害提高16%，友军造成的兵刃伤害和谋略伤害提高5%
type WarBookDetailType_42 struct {
}

func (w *WarBookDetailType_42) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

}
