package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 天下无双
// 对敌军单体发起决斗，决斗双方轮流向对方普通攻击3次，自己率先出手。
// 决斗途中，双手不受缴械和震慑状态影响，并且可以触发群攻和突击战法，
// 自身为主将时，决斗后自身受到兵刃伤害降低7%（受武力影响），持续2回合
// 主动35%
type UniqueTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (u UniqueTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	u.tacticsParams = tacticsParams
	u.triggerRate = 0.35
	return u
}

func (u UniqueTactic) Prepare() {

}

func (u UniqueTactic) Id() consts.TacticId {
	return consts.Unique
}

func (u UniqueTactic) Name() string {
	return "天下无双"
}

func (u UniqueTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (u UniqueTactic) GetTriggerRate() float64 {
	return u.triggerRate
}

func (u UniqueTactic) SetTriggerRate(rate float64) {
	u.triggerRate = rate
}

func (u UniqueTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (u UniqueTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (u UniqueTactic) Execute() {
	ctx := u.tacticsParams.Ctx
	currentGeneral := u.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		u.Name(),
	)
	// 对敌军单体发起决斗，决斗双方轮流向对方普通攻击3次，自己率先出手。
	//enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, u.tacticsParams)

	// 决斗途中，双手不受缴械和震慑状态影响，并且可以触发群攻和突击战法，
	// 自身为主将时，决斗后自身受到兵刃伤害降低7%（受武力影响），持续2回合

}

func (u UniqueTactic) IsTriggerPrepare() bool {
	return false
}

func (a UniqueTactic) SetTriggerPrepare(triggerPrepare bool) {
}
