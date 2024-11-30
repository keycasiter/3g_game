package execute

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type WarBookExecutor struct {
	Ctx     context.Context
	General *vo.BattleGeneral
	//对战战法全局holder
	TacticsParams *model.TacticsParams
}

func NewWarBookExecutor(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) *WarBookExecutor {
	return &WarBookExecutor{
		Ctx:           ctx,
		General:       general,
		TacticsParams: tacticParams,
	}
}

func (w *WarBookExecutor) Execute() {
	for _, warbook := range w.General.WarBooks {
		hlog.CtxInfof(w.Ctx, "[%v]兵书[%v]效果处理...", w.General.BaseInfo.Name, warbook.Name)
		if handler, ok := WarbookHandlerMap[consts.WarBookDetailType(warbook.Id)]; ok {
			handler.Handle(w.Ctx, w.General, w.TacticsParams)
		}
	}
}
