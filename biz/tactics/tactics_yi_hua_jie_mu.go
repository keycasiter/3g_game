package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 移花接木
// 使敌我全体受到治疗提升18%（受自身最高属性影响），并将敌军全体受到治疗的26%（受自身最高属性影响）转移到自身，持续1回合
type YiHuaJieMuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a YiHuaJieMuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.5
	return a
}

func (a YiHuaJieMuTactic) Prepare() {

}

func (a YiHuaJieMuTactic) Id() consts.TacticId {
	return consts.YiHuaJieMu
}

func (a YiHuaJieMuTactic) Name() string {
	return "移花接木"
}

func (a YiHuaJieMuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a YiHuaJieMuTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a YiHuaJieMuTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a YiHuaJieMuTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a YiHuaJieMuTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a YiHuaJieMuTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
}

func (a YiHuaJieMuTactic) IsTriggerPrepare() bool {
	return false
}
