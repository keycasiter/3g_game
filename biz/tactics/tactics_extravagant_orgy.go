package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 酒池肉林
// 战斗中，每回合使我军男性武将获得4%倒戈及10点武力，同时降低10点智力，可叠加，持续到战斗结束；
// 战斗第5回合起，每回合对敌军全体和友军男性武将造成兵刃伤害（伤害率60%）
// 自身为主将时，伤害率提升至68%
// 被动 ，100%
type ExtravagantOrgyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e ExtravagantOrgyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 1.0
	return e
}

func (e ExtravagantOrgyTactic) Prepare() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral

	// 战斗中，每回合使我军男性武将获得4%倒戈及10点武力，同时降低10点智力，可叠加，持续到战斗结束；
	malePairGenerals := make([]*vo.BattleGeneral, 0)
	//找到我军全体
	pairGenerals := util.GetPairGeneralArr(e.tacticsParams)
	for _, general := range pairGenerals {
		if general.BaseInfo.Gender == consts.Gender_Male {
			malePairGenerals = append(malePairGenerals, general)
		}
	}
	// 战斗中，每回合使我军男性武将获得4%倒戈及10点武力，同时降低10点智力，可叠加，持续到战斗结束；
	if len(malePairGenerals) > 0 {
		for _, general := range malePairGenerals {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := e.tacticsParams.CurrentGeneral

				//施加倒戈效果
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
					EffectRate:     0.04,
					FromTactic:     e.Id(),
					ProduceGeneral: currentGeneral,
				})
				//提高武力
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
					EffectValue:    10,
					FromTactic:     e.Id(),
					ProduceGeneral: currentGeneral,
				})
				//降低智力
				util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
					EffectValue:    10,
					FromTactic:     e.Id(),
					ProduceGeneral: currentGeneral,
				})

				return triggerResp
			})
		}
	}
	// 战斗第5回合起，每回合对敌军全体和友军男性武将造成兵刃伤害（伤害率60%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if triggerRound >= consts.Battle_Round_Fifth {
			//找到全体男性武将
			maleAllGenerals := make([]*vo.BattleGeneral, 0)
			allGenerals := util.GetAllGeneralsNotSelfByGeneral(triggerGeneral, e.tacticsParams)
			for _, general := range allGenerals {
				if general.BaseInfo.Gender == consts.Gender_Male {
					maleAllGenerals = append(maleAllGenerals, general)
				}
			}

			for _, general := range maleAllGenerals {
				// 自身为主将时，伤害率提升至68%
				rate := 0.6
				if triggerGeneral.IsMaster {
					rate = 0.68
				}
				dmg := cast.ToInt64(general.BaseInfo.AbilityAttr.ForceBase * rate)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: e.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: general,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      e.Id(),
					TacticName:    e.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (e ExtravagantOrgyTactic) Id() consts.TacticId {
	return consts.ExtravagantOrgy
}

func (e ExtravagantOrgyTactic) Name() string {
	return "酒池肉林"
}

func (e ExtravagantOrgyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (e ExtravagantOrgyTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e ExtravagantOrgyTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e ExtravagantOrgyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (e ExtravagantOrgyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e ExtravagantOrgyTactic) Execute() {

}

func (e ExtravagantOrgyTactic) IsTriggerPrepare() bool {
	return false
}
