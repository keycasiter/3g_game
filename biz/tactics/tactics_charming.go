package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：魅惑
// 战法描述：自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
// 持续1回合，自身为女性时，触发几率额外受智力影响
type CharmingTactic struct {
	tacticsParams model.TacticsParams
}

func (c CharmingTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams

	//效果施加
	tacticsParams.CurrentGeneral.TacticsTriggerMap[consts.BattleAction_SufferAttack] = &CharmingTactic{}

	return c
}

func (c CharmingTactic) Execute() {

}

func (c CharmingTactic) Name() string {
	return "魅惑"
}

func (c CharmingTactic) Id() int64 {
	return Charming
}

func (c CharmingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c CharmingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CharmingTactic) TriggerRate() float64 {
	return 1.0
}

func (c CharmingTactic) Trigger() {
	if !util.GenerateRate(0.45) {
		hlog.CtxInfof(c.tacticsParams.Ctx, "[%s]战法[%s]因几率没有生效",
			c.tacticsParams.CurrentSufferGeneral.BaseInfo.Name,
			c.Name(),
		)
		return
	}
	hitIdx := util.GenerateHitOneIdx(3)
	debuffs := []consts.DebuffEffectType{
		consts.DebuffEffectType_Chaos,
		consts.DebuffEffectType_NoStrategy,
		consts.DebuffEffectType_PoorHealth,
	}
	//施加负面效果
	c.tacticsParams.CurrentGeneral.DeBuffEffectHolderMap[debuffs[hitIdx]] = 1.0
	hlog.CtxInfof(c.tacticsParams.Ctx, "[%s]执行来自【魅惑】的「魅惑」效果",
		c.tacticsParams.CurrentSufferGeneral.BaseInfo.Name)
	hlog.CtxInfof(c.tacticsParams.Ctx, "[%s]的「%s」效果已施加",
		c.tacticsParams.CurrentGeneral.BaseInfo.Name,
		debuffs[hitIdx],
	)
	return
}
