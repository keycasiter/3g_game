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

// 义胆雄心
// 战斗中，奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
// 偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
// 自身为主将时，降低属性效果受自身对应属性影响
type BraveAmbitionTactic struct {
	tacticsParams *model.TacticsParams
}

func (b BraveAmbitionTactic) TriggerRate() float64 {
	return 1.0
}

func (b BraveAmbitionTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	return b
}

func (b BraveAmbitionTactic) Prepare() {
	// 战斗中，奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
	// 偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
	// 自身为主将时，降低属性效果受自身对应属性影响
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//注册效果
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_BraveAmbition_Prepare,
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		currentRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}
		//奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
		if currentRound%2 != 0 {
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				b.Name(),
				consts.BuffEffectType_BraveAmbition_Prepare,
			)
			//找到敌军单体
			enemyGeneral := util.GetEnemyOneGeneral(b.tacticsParams)
			//造成伤害
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.84)
			finalDmg, holdNum, remainNum, isEffect := util.TacticDamage(b.tacticsParams, triggerGeneral, enemyGeneral, dmg, consts.BattleAction_PassiveTactic)
			if !isEffect {
				return triggerResp
			}
			hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d)",
				enemyGeneral.BaseInfo.Name,
				triggerGeneral.BaseInfo.Name,
				b.Name(),
				finalDmg,
				holdNum,
				remainNum,
			)
			//持续2回合
			if !util.TacticsDebuffEffectCountWrapIncr(ctx, enemyGeneral, consts.DebuffEffectType_BraveAmbition_DecrForce, 2, 2, true) {
				return triggerResp
			}
			//降低武力64点
			decrNum := float64(64)
			//自身为主将时，降低属性效果受自身对应属性影响
			if triggerGeneral.IsMaster {
				decrNum += triggerGeneral.BaseInfo.AbilityAttr.ForceBase / 100.00
			}
			enemyGeneral.BaseInfo.AbilityAttr.ForceBase -= decrNum
			hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
				enemyGeneral.BaseInfo.Name,
				decrNum)
			//注册效果恢复
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeGeneral := params.CurrentGeneral
				revokeResp := &vo.TacticsTriggerResult{}
				//次数为0
				if 0 == util.TacticsDebuffCountGet(revokeGeneral, consts.DebuffEffectType_BraveAmbition_DecrForce) {
					revokeGeneral.BaseInfo.AbilityAttr.ForceBase += decrNum

					hlog.CtxInfof(ctx, "[%s]的「%v」效果消失",
						revokeGeneral.BaseInfo.Name,
						consts.DebuffEffectType_BraveAmbition_DecrForce,
					)
					hlog.CtxInfof(ctx, "[%s]的武力提高了%.2f",
						revokeGeneral.BaseInfo.Name,
						decrNum)
					return revokeResp
				}
				//次数消耗
				util.TacticsDebuffEffectCountWrapDecr(revokeGeneral, consts.DebuffEffectType_BraveAmbition_DecrForce, 1)
				return triggerResp
			})
		}

		//偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
		if currentRound%2 == 0 {
			//找到敌军2人
			enemyGenerals := util.GetEnemyGeneralsTwoArr(b.tacticsParams)
			//造成伤害
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.76)
			for _, enemyGeneral := range enemyGenerals {
				finalDmg, holdNum, remainNum, isEffect := util.TacticDamage(b.tacticsParams, triggerGeneral, enemyGeneral, dmg, consts.BattleAction_PassiveTactic)
				if !isEffect {
					continue
				}
				hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d)",
					enemyGeneral.BaseInfo.Name,
					triggerGeneral.BaseInfo.Name,
					b.Name(),
					finalDmg,
					holdNum,
					remainNum,
				)
				//持续2回合
				if !util.TacticsDebuffEffectCountWrapIncr(ctx, enemyGeneral, consts.DebuffEffectType_BraveAmbition_DecrIntelligence, 2, 2, true) {
					return triggerResp
				}
				//降低智力34点，持续2回合；
				decrNum := float64(34)
				//自身为主将时，降低属性效果受自身对应属性影响
				if triggerGeneral.IsMaster {
					decrNum += triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100.00
				}
				enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= decrNum
				hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
					enemyGeneral.BaseInfo.Name,
					decrNum)

				//注册效果恢复
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeGeneral := params.CurrentGeneral
					revokeResp := &vo.TacticsTriggerResult{}
					//次数为0
					if 0 == util.TacticsDebuffCountGet(revokeGeneral, consts.DebuffEffectType_BraveAmbition_DecrIntelligence) {
						revokeGeneral.BaseInfo.AbilityAttr.ForceBase += decrNum

						hlog.CtxInfof(ctx, "[%s]的「%v」效果消失",
							revokeGeneral.BaseInfo.Name,
							consts.DebuffEffectType_BraveAmbition_DecrIntelligence,
						)
						hlog.CtxInfof(ctx, "[%s]的智力提高了%.2f",
							revokeGeneral.BaseInfo.Name,
							decrNum)
						return revokeResp
					}
					//次数消耗
					util.TacticsDebuffEffectCountWrapDecr(revokeGeneral, consts.DebuffEffectType_BraveAmbition_DecrIntelligence, 1)
					return triggerResp
				})
			}
		}

		return triggerResp
	})
}

func (b BraveAmbitionTactic) Id() consts.TacticId {
	return consts.BraveAmbition
}

func (b BraveAmbitionTactic) Name() string {
	return "义胆雄心"
}

func (b BraveAmbitionTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BraveAmbitionTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveAmbitionTactic) Execute() {
	return
}
