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
)

// 决水溃城
// 准备1回合，对敌军群体（2～3人）造成破坏（禁用装备）状态及水攻状态，每回合持续造成伤害（伤害率112%，受智力影响）
// 持续2回合，若该战法首回合发动则无需准备
type BreakingThroughTheWaterAndCrushingTheCityTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.45
	return b
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Prepare() {
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheWaterAndCrushingTheCity
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Name() string {
	return "决水溃城"
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//若该战法首回合发动则无需准备
	if currentRound == consts.Battle_Round_First {
		b.tacticTrigger(currentGeneral)
		return
	} else {
		b.isTriggerPrepare = true
		hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
			currentGeneral.BaseInfo.Name,
			b.Name(),
		)
	}
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			b.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if b.isTriggered {
				return triggerResp
			} else {
				b.isTriggered = true
			}

			b.tacticTrigger(triggerGeneral)
		}

		return triggerResp
	})
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) IsTriggerPrepare() bool {
	return false
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) tacticTrigger(currentGeneral *vo.BattleGeneral) {
	ctx := b.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	//对敌军群体（2～3人）造成破坏（禁用装备）状态及水攻状态，每回合持续造成伤害（伤害率112%，受智力影响）,持续2回合
	//找到敌军2~3人
	enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(currentGeneral, b.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		//施加破坏效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Break, &vo.EffectHolderParams{
			EffectRate:  1.0,
			EffectRound: 2,
			FromTactic:  b.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Break,
					TacticId:   b.Id(),
				})

				return revokeResp
			})
		}
		//施加水攻效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_WaterAttack, &vo.EffectHolderParams{
			EffectRate:  1.12,
			EffectRound: 2,
			FromTactic:  b.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_WaterAttack,
					TacticId:   b.Id(),
				}) {
					//每回合持续造成伤害（伤害率112%，受智力影响）
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     b.tacticsParams,
						AttackGeneral:     currentGeneral,
						SufferGeneral:     revokeGeneral,
						DamageType:        consts.DamageType_Strategy,
						DamageImproveRate: 1.12,
						TacticId:          b.Id(),
						TacticName:        b.Name(),
						EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
					})
				}

				return revokeResp
			})
		}
	}
}

func (a BreakingThroughTheWaterAndCrushingTheCityTactic) SetTriggerPrepare(triggerPrepare bool) {
}
