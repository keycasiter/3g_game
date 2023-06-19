package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 高橹连营
// 战斗开始时，架起箭楼，填装10次（次数为0则失效），每回合发射2次，有50%概率（受武力影响）额外发射1次，对敌军单体造成兵刃伤害（每次目标独立判断，伤害率82%）
// 并有50%概率使其受到伤害提高6%，可叠加2次，首回合额外造成缴械效果，每回合有75%概率无法普通攻击，持续2回合；
// 自身为主将时，每回合会填装1次
// 被动 100%
type HighWoodenPaddlesConnectedToTheCampTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 1.0
	return h
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) Prepare() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)
	// 战斗开始时，架起箭楼，填装10次（次数为0则失效），每回合发射2次，有50%概率（受武力影响）额外发射1次，对敌军单体造成兵刃伤害（每次目标独立判断，伤害率82%）
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_HighWoodenPaddlesConnectedToTheCamp_Prepare, &vo.EffectHolderParams{
		EffectTimes:    10,
		FromTactic:     h.Id(),
		ProduceGeneral: currentGeneral,
	})

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeGeneral := params.CurrentGeneral
		revokeRound := params.CurrentRound

		// 自身为主将时，每回合会填装1次
		if currentGeneral.IsMaster {
			util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_HighWoodenPaddlesConnectedToTheCamp_Prepare, &vo.EffectHolderParams{
				EffectTimes:    1,
				FromTactic:     h.Id(),
				ProduceGeneral: currentGeneral,
			})
		}

		//每回合发射次数
		launchTimes := 2
		triggerRate := 0.5 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
		for i := 0; i < 2; i++ {
			if util.GenerateRate(triggerRate) {
				launchTimes++
			}
		}

		for i := 0; i < launchTimes; i++ {
			if util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_HighWoodenPaddlesConnectedToTheCamp_Prepare,
				CostTimes:  1,
				TacticId:   h.Id(),
			}) {
				//找到敌军单体
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, h.tacticsParams)
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.82)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: h.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticName:    h.Name(),
				})
				// 并有50%概率使其受到伤害提高6%，可叠加2次
				if util.GenerateRate(0.5) {
					util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
						EffectTimes:    1,
						MaxEffectTimes: 2,
						FromTactic:     h.Id(),
						ProduceGeneral: currentGeneral,
					})
					util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
						EffectTimes:    1,
						MaxEffectTimes: 2,
						FromTactic:     h.Id(),
						ProduceGeneral: currentGeneral,
					})
				}
				//首回合额外造成缴械效果，每回合有75%概率无法普通攻击，持续2回合；
				if revokeRound == consts.Battle_Round_First {
					if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
						TriggerRate:    0.75,
						EffectRound:    2,
						FromTactic:     h.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						//注册消失效果
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    enemyGeneral,
							EffectType: consts.DebuffEffectType_CancelWeapon,
							TacticId:   h.Id(),
						})
					}
				}
			}
		}

		return revokeResp
	})
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) Id() consts.TacticId {
	return consts.HighWoodenPaddlesConnectedToTheCamp
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) Name() string {
	return "高橹连营"
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HighWoodenPaddlesConnectedToTheCampTactic) Execute() {

}

func (h HighWoodenPaddlesConnectedToTheCampTactic) IsTriggerPrepare() bool {
	return false
}
