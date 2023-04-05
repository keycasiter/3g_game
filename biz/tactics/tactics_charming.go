package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
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
	return c
}

func (c CharmingTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral
	//自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
	// 持续1回合，自身为女性时，触发几率额外受智力影响
	//效果施加
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Charming] = 1.0
	//触发效果注册
	util.TacticsTriggerWrapSet(currentGeneral, consts.BattleAction_SufferAttack, func(params vo.TacticsTriggerParams) {
		//有45%几率
		triggerRate := 0.45
		//自身为女性时，触发几率额外受智力影响
		if currentGeneral.BaseInfo.Gender == consts.Gender_Female {
			//TODO
		}
		if !util.GenerateRate(triggerRate) {
			hlog.CtxInfof(ctx, "[%s]执行来自[%s]【%s】的「魅惑」效果因几率没有生效",
				currentGeneral.BaseInfo.Name,
				currentGeneral.BaseInfo.Name,
				c.Name(),
			)
			return
		} else {
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「魅惑」效果",
				currentGeneral.BaseInfo.Name,
				c.Name(),
			)
			//进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种
			debuffs := []consts.DebuffEffectType{
				consts.DebuffEffectType_Chaos,
				consts.DebuffEffectType_NoStrategy,
				consts.DebuffEffectType_PoorHealth,
			}
			hitIdx := util.GenerateHitOneIdx(3)
			debuff := debuffs[hitIdx]
			//找到攻击者
			params.CurrentAttackGeneral.DeBuffEffectHolderMap[debuff] = 1.0
			//持续1回合
			params.CurrentAttackGeneral.DeBuffEffectCountMap[debuff][1] = 1.0
		}
	})
}

func (c CharmingTactic) Execute() {
	return
}

func (c CharmingTactic) Name() string {
	return "魅惑"
}

func (c CharmingTactic) Id() consts.TacticId {
	return consts.Charming
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
