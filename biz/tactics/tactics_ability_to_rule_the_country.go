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

//经天纬地
//战斗中，我军全体发动主动战法及突击战法时，自身有35%概率（受智力影响）会对敌军单体发动谋略攻击（伤害率50%，受智力影响）
//若由自身触发，则额外由50%概率使我军群体（2人）获得5%攻心（造成谋略伤害时，恢复自身基于伤害量的一定兵力），持续2回合 ，可叠加3次攻心值；
//自身为主将时，攻心效果提高之10%
type AbilityToRuleTheCountryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AbilityToRuleTheCountryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a AbilityToRuleTheCountryTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//找到我军队伍
	pairGeneralArr := util.GetPairGeneralArr(a.tacticsParams)
	for _, pairGeneral := range pairGeneralArr {
		//施加效果
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_AbilityToRuleTheCountry_Prepare, &vo.EffectHolderParams{
			EffectRate: 1.0,
			FromTactic: a.Id(),
		})

		registerFunc := func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}

			//触发概率
			triggerRate := 0.35
			triggerRate += currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
			//概率判断
			if !util.GenerateRate(triggerRate) {
				return triggerResp
			}
			//对单体发动谋略攻击
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(pairGeneral, a.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.5)
			util.TacticDamage(&util.TacticDamageParam{
				TacticsParams: a.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: enemyGeneral,
				Damage:        dmg,
				TacticName:    a.Name(),
			})

			//若由自身触发，则额外有50%概率使我军群体（2人）获得5%攻心（造成谋略伤害时，恢复自身基于伤害量的一定兵力），持续2回合 ，可叠加3次攻心值；
			//自身为主将时，攻心效果提高至10%
			if pairGeneral.BaseInfo.Id == currentGeneral.BaseInfo.Id {
				//攻心效果
				effectRate := 0.05
				if currentGeneral.IsMaster {
					effectRate = 0.1
				}

				//概率判断
				if !util.GenerateRate(0.5) {
					return triggerResp
				}
				//我军2人
				pairGenerals := util.GetPairGeneralsTwoArrByGeneral(pairGeneral, a.tacticsParams)
				for _, general := range pairGenerals {
					//攻心效果施加
					if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_AttackHeart, &vo.EffectHolderParams{
						EffectRate:     effectRate,
						FromTactic:     a.Id(),
						EffectTimes:    1,
						EffectRound:    2,
						MaxEffectTimes: 3,
					}).IsSuccess {
						//攻心效果失效
						util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeGeneral := params.CurrentGeneral
							revokeResp := &vo.TacticsTriggerResult{}
							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.BuffEffectType_AttackHeart,
								TacticId:   a.Id(),
								CostOverTriggerFunc: func() {
									hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
										general.BaseInfo.Name,
										consts.BuffEffectType_AttackHeart,
									)
								},
							})
							return revokeResp
						})
					}
				}
			}
			return triggerResp
		}

		//注册发动主动战法效果
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_ActiveTacticEnd, registerFunc)
		//注册发动突击战法效果
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AssaultTacticEnd, registerFunc)
	}
}

func (a AbilityToRuleTheCountryTactic) Id() consts.TacticId {
	return consts.AbilityToRuleTheCountry
}

func (a AbilityToRuleTheCountryTactic) Name() string {
	return "经天纬地"
}

func (a AbilityToRuleTheCountryTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AbilityToRuleTheCountryTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AbilityToRuleTheCountryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a AbilityToRuleTheCountryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AbilityToRuleTheCountryTactic) Execute() {
}

func (a AbilityToRuleTheCountryTactic) IsTriggerPrepare() bool {
	return false
}
