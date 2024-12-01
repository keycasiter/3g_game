package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 至柔动刚
// 战斗中，敌我全体普通攻击伤害降低35%，自身受到伤害时有50%概率（受智力影响）偷取伤害来源10点属性（智力、统率、速度随机一种，受智力影响，可叠加，持续5回合）
// 我军全体普通攻击后，自身有50%概率（受智力影响）对敌军单体造成一次谋略伤害（伤害率106%，受智力影响）
type ZhiRouDongGangTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a ZhiRouDongGangTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a ZhiRouDongGangTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，敌我全体普通攻击伤害降低35%
	allGenerals := util.GetAllGenerals(a.tacticsParams)
	for _, general := range allGenerals {
		//敌我全体普通攻击伤害降低35%
		util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchGeneralAttackDeduce, &vo.EffectHolderParams{
			EffectRate:     0.35,
			FromTactic:     a.Id(),
			ProduceGeneral: general,
		})
	}
	// 自身受到伤害时有50%概率（受智力影响）偷取伤害来源10点属性（智力、统率、速度随机一种，受智力影响，可叠加，持续5回合）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		triggerRate := 0.5 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
		if util.GenerateRate(triggerRate) {
			targetGeneral := params.AttackGeneral
			effectVal := 10 + cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100)

			attArr := []consts.AbilityAttr{
				consts.AbilityAttr_Intelligence,
				consts.AbilityAttr_Command,
				consts.AbilityAttr_Speed,
			}
			hitIdx := util.GenerateHitOneIdx(len(attArr))
			switch attArr[hitIdx] {
			case consts.AbilityAttr_Intelligence:
				//减少
				if util.DebuffEffectWrapSet(ctx, targetGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    targetGeneral,
						EffectType: consts.DebuffEffectType_DecrIntelligence,
						TacticId:   a.Id(),
					})
				}

				//增加
				if util.BuffEffectWrapSet(ctx, targetGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerGeneral,
						EffectType: consts.BuffEffectType_IncrIntelligence,
						TacticId:   a.Id(),
					})
				}
			case consts.AbilityAttr_Command:
				//减少
				if util.DebuffEffectWrapSet(ctx, targetGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    targetGeneral,
						EffectType: consts.DebuffEffectType_DecrCommand,
						TacticId:   a.Id(),
					})
				}

				//增加
				if util.BuffEffectWrapSet(ctx, targetGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerGeneral,
						EffectType: consts.BuffEffectType_IncrIntelligence,
						TacticId:   a.Id(),
					})
				}
			case consts.AbilityAttr_Speed:
				//减少
				if util.DebuffEffectWrapSet(ctx, targetGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    targetGeneral,
						EffectType: consts.DebuffEffectType_DecrSpeed,
						TacticId:   a.Id(),
					})
				}

				//增加
				if util.BuffEffectWrapSet(ctx, targetGeneral, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    5,
					EffectTimes:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerGeneral,
						EffectType: consts.BuffEffectType_IncrSpeed,
						TacticId:   a.Id(),
					})
				}
			}
		}

		return triggerResp
	})

	// 我军全体普通攻击后，自身有50%概率（受智力影响）对敌军单体造成一次谋略伤害（伤害率106%，受智力影响）
	pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
	for _, general := range pairGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRate := 0.5 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100

			if util.GenerateRate(triggerRate) {
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, a.tacticsParams)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     a.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: 1.06 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
					TacticId:          a.Id(),
					TacticName:        a.Name(),
				})
			}

			return triggerResp
		})
	}
}

func (a ZhiRouDongGangTactic) Id() consts.TacticId {
	return consts.ZhiRouDongGang
}

func (a ZhiRouDongGangTactic) Name() string {
	return "至柔动刚"
}

func (a ZhiRouDongGangTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a ZhiRouDongGangTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a ZhiRouDongGangTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a ZhiRouDongGangTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a ZhiRouDongGangTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a ZhiRouDongGangTactic) Execute() {
}

func (a ZhiRouDongGangTactic) IsTriggerPrepare() bool {
	return false
}

func (a ZhiRouDongGangTactic) SetTriggerPrepare(triggerPrepare bool) {
}
