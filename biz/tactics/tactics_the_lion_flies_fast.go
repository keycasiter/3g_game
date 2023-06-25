package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 狮子奋迅
// 对敌军单体及额外1～2名敌人造成兵刃攻击（伤害率118%），并使自身主动战法发动几率提高10%，
// 自身为主将时，发动几率提高至15%，持续2回合，如果单体目标为敌军主将则使其陷入叛逃状态，每回合持续造成伤害（伤害率102%），持续2回合
// 主动，35%
type TheLionFliesFastTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheLionFliesFastTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TheLionFliesFastTactic) Prepare() {
}

func (t TheLionFliesFastTactic) Id() consts.TacticId {
	return consts.TheLionFliesFast
}

func (t TheLionFliesFastTactic) Name() string {
	return "狮子奋迅"
}

func (t TheLionFliesFastTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TheLionFliesFastTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheLionFliesFastTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheLionFliesFastTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheLionFliesFastTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheLionFliesFastTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	// 对敌军单体及额外1～2名敌人造成兵刃攻击（伤害率118%），并使自身主动战法发动几率提高10%，
	// 自身为主将时，发动几率提高至15%，持续2回合，如果单体目标为敌军主将则使其陷入叛逃状态，每回合持续造成伤害（伤害率102%），持续2回合
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
	enemyGenerals := util.GetEnemyGeneralsOneOrTwoArr(t.tacticsParams)
	enemyGenerals = append(enemyGenerals, enemyGeneral)
	for _, general := range enemyGenerals {
		//伤害
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.18)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: t.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: general,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      t.Id(),
			TacticName:    t.Name(),
		})
	}
	//并使自身主动战法发动几率提高10%，自身为主将时，发动几率提高至15%，持续2回合
	triggerRate := 0.1
	if currentGeneral.IsMaster {
		triggerRate = 0.15
	}
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove, &vo.EffectHolderParams{
		TriggerRate:    triggerRate,
		EffectRound:    2,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}

	//如果单体目标为敌军主将则使其陷入叛逃状态，每回合持续造成伤害（伤害率102%），持续2回合
	if enemyGeneral.IsMaster {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Defect, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Defect,
					TacticId:   t.Id(),
				}) {
					//伤害
					attrType, val := util.GetGeneralHighestBetweenForceOrIntelligence(revokeGeneral)
					dmgType := consts.DamageType_Weapon
					if attrType == consts.AbilityAttr_Intelligence {
						dmgType = consts.DamageType_Strategy
					}

					dmg := cast.ToInt64(val * 1.02)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams:  t.tacticsParams,
						AttackGeneral:  currentGeneral,
						SufferGeneral:  revokeGeneral,
						DamageType:     dmgType,
						Damage:         dmg,
						TacticId:       t.Id(),
						TacticName:     t.Name(),
						EffectName:     fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
						IsIgnoreDefend: true,
					})
				}

				return revokeResp
			})
		}
	}
}

func (t TheLionFliesFastTactic) IsTriggerPrepare() bool {
	return false
}
