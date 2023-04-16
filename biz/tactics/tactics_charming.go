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
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CharmingTactic) IsTriggerPrepare() bool {
	return false
}

func (c CharmingTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CharmingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CharmingTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CharmingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CharmingTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
	// 持续1回合，自身为女性时，触发几率额外受智力影响
	//效果施加
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Charming] = 1.0
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_Charming,
	)
	//触发效果注册
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}
		attackGeneral := params.AttackGeneral
		//有45%几率
		triggerRate := 0.45
		//自身为女性时，触发几率额外受智力影响
		if triggerGeneral.BaseInfo.Gender == consts.Gender_Female {
			//TODO
		}
		if !util.GenerateRate(triggerRate) {
			hlog.CtxInfof(ctx, "[%s]执行来自[%s]【%s】的「魅惑」效果因几率没有生效",
				triggerGeneral.BaseInfo.Name,
				triggerGeneral.BaseInfo.Name,
				c.Name(),
			)
			return triggerResp
		} else {
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「魅惑」效果",
				triggerGeneral.BaseInfo.Name,
				c.Name(),
			)
			//进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种
			debuffs := []consts.DebuffEffectType{
				consts.DebuffEffectType_Chaos,
				consts.DebuffEffectType_NoStrategy,
				consts.DebuffEffectType_PoorHealth,
			}
			hitIdx := util.GenerateHitOneIdx(len(debuffs))
			debuff := debuffs[hitIdx]

			if !util.DebuffEffectWrapSet(ctx, attackGeneral, debuff, 1.0) {
				return triggerResp
			}
			//次数设置
			if !util.TacticsDebuffEffectCountWrapIncr(ctx, attackGeneral, debuff, 1, 1, false) {
				return triggerResp
			}
			//消失效果注册
			util.TacticsTriggerWrapRegister(attackGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeGeneral := params.CurrentGeneral
				revokeRound := params.CurrentRound
				resp := &vo.TacticsTriggerResult{}
				//持续一回合
				if triggerRound+2 == revokeRound {
					util.DebuffEffectWrapRemove(ctx, revokeGeneral, debuff)
				}

				return resp
			})
		}

		return triggerResp
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
