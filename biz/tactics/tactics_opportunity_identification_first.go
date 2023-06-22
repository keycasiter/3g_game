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

// 机鉴先识
// 准备回合使我军群体（2～3人）获得2次警戒，随后每回合有42%概率（受智力影响）使我军群体（2～3人）获得1次警戒（持续3回合，可重复获得4次）；
// 受到伤害超过自身可携带最大兵力的6%时（最低100兵力），使该次伤害降低40%（受智力影响）并消耗一次警戒；
// 自身为主将时，战斗前2回合，使我军全体受到缴械、计穷、混乱、震慑状态时，有75%概率同时施加给敌军单体，每回合每人最多生效一次
// 指挥 100%
type OpportunityIdentificationFirstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OpportunityIdentificationFirstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	o.tacticsParams = tacticsParams
	o.triggerRate = 1.0
	return o
}

func (o OpportunityIdentificationFirstTactic) Prepare() {
	ctx := o.tacticsParams.Ctx
	currentGeneral := o.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		o.Name(),
	)
	// 准备回合使我军群体（2～3人）获得2次警戒
	pairGenerals := util.GetPairGeneralsTwoOrThreeMap(o.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Alert, &vo.EffectHolderParams{
			EffectTimes:    2,
			EffectRound:    3,
			MaxEffectTimes: 4,
			FromTactic:     o.Id(),
			ProduceGeneral: currentGeneral,
		})
	}
	//随后每回合有42%概率（受智力影响）使我军群体（2～3人）获得1次警戒（持续3回合，可重复获得4次）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRate := 0.42 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100

		if util.GenerateRate(triggerRate) {
			generals := util.GetPairGeneralsTwoOrThreeMap(o.tacticsParams)
			for _, general := range generals {
				util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Alert, &vo.EffectHolderParams{
					EffectTimes:    1,
					EffectRound:    3,
					MaxEffectTimes: 4,
					FromTactic:     o.Id(),
					ProduceGeneral: triggerGeneral,
				})
			}
		}

		return triggerResp
	})

	// 受到伤害超过自身可携带最大兵力的6%时（最低100兵力），使该次伤害降低40%（受智力影响）并消耗一次警戒；
	allPairGenerals := util.GetPairGeneralArr(o.tacticsParams)
	for _, general := range allPairGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDamage, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			attackGeneral := params.AttackGeneral

			if cast.ToFloat64(o.tacticsParams.CurrentDamageNum/triggerGeneral.SoldierNum) > 0.06 {
				//消耗一次警戒
				if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    triggerGeneral,
					EffectType: consts.BuffEffectType_Alert,
					TacticId:   o.Id(),
				}) {
					if util.DebuffEffectWrapSet(ctx, attackGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
						EffectRate:     0.4,
						FromTactic:     o.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    attackGeneral,
							EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
							TacticId:   o.Id(),
						})
					}
				}
			}

			return triggerResp
		})
	}
	// 自身为主将时，战斗前2回合，使我军全体受到缴械、计穷、混乱、震慑状态时，有75%概率同时施加给敌军单体，每回合每人最多生效一次
	perGeneralRoundCntMap := make(map[consts.General_Id]map[consts.BattleRound]bool, 0)
	if currentGeneral.IsMaster {
		for _, pairGeneral := range allPairGenerals {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferDebuffEffectEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				effectType := params.DebuffEffect
				effectRound := params.EffectRound
				currentRound := params.CurrentRound

				if currentRound <= consts.Battle_Round_Second {
					//每回合每人最多生效一次
					if roundMap, ok := perGeneralRoundCntMap[consts.General_Id(pairGeneral.BaseInfo.Id)]; ok {
						if roundMap[currentRound] {
							return triggerResp
						}
					}
					//有75%概率同时施加给敌军单体
					if effectType == consts.DebuffEffectType_CancelWeapon ||
						effectType == consts.DebuffEffectType_NoStrategy ||
						effectType == consts.DebuffEffectType_Chaos ||
						effectType == consts.DebuffEffectType_Awe {
						if util.GenerateRate(0.75) {
							enemyGeneral := util.GetEnemyOneGeneralByGeneral(pairGeneral, o.tacticsParams)
							util.DebuffEffectWrapSet(ctx, enemyGeneral, effectType, &vo.EffectHolderParams{
								EffectRound:    effectRound,
								FromTactic:     o.Id(),
								ProduceGeneral: pairGeneral,
							})
						}
						if perGeneralRoundCntMap == nil {
							perGeneralRoundCntMap = map[consts.General_Id]map[consts.BattleRound]bool{}
						}
						perGeneralRoundCntMap[consts.General_Id(pairGeneral.BaseInfo.Id)][currentRound] = true
					}
				}

				return triggerResp
			})
		}
	}

}

func (o OpportunityIdentificationFirstTactic) Id() consts.TacticId {
	return consts.OpportunityIdentificationFirst
}

func (o OpportunityIdentificationFirstTactic) Name() string {
	return "机鉴先识"
}

func (o OpportunityIdentificationFirstTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (o OpportunityIdentificationFirstTactic) GetTriggerRate() float64 {
	return o.triggerRate
}

func (o OpportunityIdentificationFirstTactic) SetTriggerRate(rate float64) {
	o.triggerRate = rate
}

func (o OpportunityIdentificationFirstTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (o OpportunityIdentificationFirstTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (o OpportunityIdentificationFirstTactic) Execute() {
}

func (o OpportunityIdentificationFirstTactic) IsTriggerPrepare() bool {
	return false
}
