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

// 兴云布雨
// 战斗第2回合开始，使敌军全体进入水攻状态，每回合持续造成伤害（伤害率72%，受智力影响），并使其受到伤害增加10%，持续5回合
// 指挥，100%
type MakeCloudAndRainTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MakeCloudAndRainTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 1.0
	return m
}

func (m MakeCloudAndRainTactic) Prepare() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	// 战斗第2回合开始，使敌军全体进入水攻状态，每回合持续造成伤害（伤害率72%，受智力影响），并使其受到伤害增加10%，持续5回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Second {
			enemyGenerals := util.GetEnemyGeneralArr(m.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//水攻
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_WaterAttack, &vo.EffectHolderParams{
					EffectRound:    5,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_WaterAttack,
							TacticId:   m.Id(),
						}) {
							dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.72)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: m.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: revokeGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        dmg,
								TacticName:    m.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
							})
						}

						return revokeResp
					})
				}
				//受到兵刃伤害提升
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.1,
					EffectRound:    5,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
							TacticId:   m.Id(),
						})

						return revokeResp
					})
				}
				//受到谋略伤害提升
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.1,
					EffectRound:    5,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
							TacticId:   m.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (m MakeCloudAndRainTactic) Id() consts.TacticId {
	return consts.MakeCloudAndRain
}

func (m MakeCloudAndRainTactic) Name() string {
	return "兴云布雨"
}

func (m MakeCloudAndRainTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (m MakeCloudAndRainTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MakeCloudAndRainTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MakeCloudAndRainTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (m MakeCloudAndRainTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MakeCloudAndRainTactic) Execute() {

}

func (m MakeCloudAndRainTactic) IsTriggerPrepare() bool {
	return false
}
