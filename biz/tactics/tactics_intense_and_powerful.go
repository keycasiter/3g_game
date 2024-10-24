package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 威谋靡亢
// 准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续2回合；如果目标已处于虚弱状态则使其陷入叛逃状态，
// 每回合持续造成伤害（伤害率158%，受武力或智力最高一项影响，无视防御），持续2回合
type IntenseAndPowerfulTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (i IntenseAndPowerfulTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.4
	return i
}

func (i IntenseAndPowerfulTactic) Prepare() {
}

func (i IntenseAndPowerfulTactic) Id() consts.TacticId {
	return consts.IntenseAndPowerful
}

func (i IntenseAndPowerfulTactic) Name() string {
	return "威谋靡亢"
}

func (i IntenseAndPowerfulTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (i IntenseAndPowerfulTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IntenseAndPowerfulTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IntenseAndPowerfulTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i IntenseAndPowerfulTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IntenseAndPowerfulTactic) Execute() {
	// 准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续2回合；如果目标已处于虚弱状态则使其陷入叛逃状态，
	// 每回合持续造成伤害（伤害率158%，受武力或智力最高一项影响，无视防御），持续2回合

	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral
	currentRound := i.tacticsParams.CurrentRound

	i.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			i.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if i.isTriggered {
				return triggerResp
			} else {
				i.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				i.Name(),
			)

			//准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续2回合；如果目标已处于虚弱状态则使其陷入叛逃状态，
			//每回合持续造成伤害（伤害率158%，受武力或智力最高一项影响，无视防御），持续2回合

			//找到敌军群体2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, i.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//是否已存在虚弱状态
				if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_PoorHealth) {
					if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Defect, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     i.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_Defect,
								TacticId:   i.Id(),
							}) {
								attrType, val := util.GetGeneralHighestBetweenForceOrIntelligence(currentGeneral)
								dmg := int64(0)
								dmgType := consts.DamageType_None
								switch attrType {
								case consts.AbilityAttr_Force:
									dmg = cast.ToInt64(val * 1.58)
									dmgType = consts.DamageType_Weapon
								case consts.AbilityAttr_Intelligence:
									dmg = cast.ToInt64(val * 1.58)
									dmgType = consts.DamageType_Strategy
								}
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:  i.tacticsParams,
									AttackGeneral:  currentGeneral,
									SufferGeneral:  revokeGeneral,
									DamageType:     dmgType,
									Damage:         dmg,
									TacticId:       i.Id(),
									TacticName:     i.Name(),
									EffectName:     fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
									IsIgnoreDefend: true,
								})
							}

							return revokeResp
						})
					}
				} else {
					if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     i.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_PoorHealth,
								TacticId:   i.Id(),
							})

							return revokeResp
						})
					}
				}
			}

		}

		return triggerResp
	})
}

func (i IntenseAndPowerfulTactic) IsTriggerPrepare() bool {
	return i.isTriggerPrepare
}
