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

// 据水断桥
// 对敌军群体（2-3人）造成溃逃状态，每回合持续造成伤害（伤害率78%，受武力影响），并使其造成伤害降低8%（受双方武力之差影响），
// 同时使自身获得16%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合，该战法发动后回进入1回合冷却
type BrokenBridgeByWaterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BrokenBridgeByWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BrokenBridgeByWaterTactic) Prepare() {

}

func (b BrokenBridgeByWaterTactic) Id() consts.TacticId {
	return consts.BrokenBridgeByWater
}

func (b BrokenBridgeByWaterTactic) Name() string {
	return "据水断桥"
}

func (b BrokenBridgeByWaterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BrokenBridgeByWaterTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BrokenBridgeByWaterTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BrokenBridgeByWaterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BrokenBridgeByWaterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BrokenBridgeByWaterTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//该战法发动后回进入1回合冷却
	//判断是否冷却
	if ok := currentGeneral.TacticFrozenMap[b.Id()]; ok {
		hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效，无法发动",
			currentGeneral.BaseInfo.Name,
			b.Name(),
		)
		return
	}

	currentGeneral.TacticFrozenMap[b.Id()] = true
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//注册冷却效果消失
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeRound := params.CurrentRound

		//1回合冷却，下下回合冷却结束
		if currentRound+2 == revokeRound {
			currentGeneral.TacticFrozenMap[b.Id()] = false

			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果已消失",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)
		}
		return revokeResp
	})

	// 对敌军群体（2-3人）造成溃逃状态，每回合持续造成伤害（伤害率78%，受武力影响），并使其造成伤害降低8%（受双方武力之差影响），
	// 同时使自身获得16%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合，该战法发动后回进入1回合冷却

	//找到敌军2～3人
	enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(currentGeneral, b.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		//施加溃逃状态
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Escape, &vo.EffectHolderParams{
			EffectRate:  1.0,
			EffectRound: 2,
			FromTactic:  b.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				//消耗回合
				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Escape,
					TacticId:   b.Id(),
				}) {
					//每回合持续造成伤害（伤害率78%，受武力影响）
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     b.tacticsParams,
						AttackGeneral:     currentGeneral,
						SufferGeneral:     revokeGeneral,
						DamageType:        consts.DamageType_Weapon,
						DamageImproveRate: 0.78,
						TacticId:          b.Id(),
						TacticName:        b.Name(),
						EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Escape),
					})
				}
				return revokeResp
			})
		}

		//并使其造成伤害降低8%（受双方武力之差影响）
		//兵刃伤害降低效果
		diff := util.CalculateAttrDiff(currentGeneral.BaseInfo.AbilityAttr.ForceBase, sufferGeneral.BaseInfo.AbilityAttr.ForceBase)
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:  0.08 + diff/100/100,
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
					EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
					TacticId:   b.Id(),
				})
				return revokeResp
			})
		}
		//谋略伤害降低效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:  0.08 + diff/100/100,
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
					EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
					TacticId:   b.Id(),
				})
				return revokeResp
			})
		}
	}

	//同时使自身获得16%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
		EffectRate:  0.16,
		EffectRound: 2,
		FromTactic:  b.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeGeneral := params.CurrentGeneral
			revokeResp := &vo.TacticsTriggerResult{}

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Defection,
				TacticId:   b.Id(),
			})

			return revokeResp
		})
	}
}

func (b BrokenBridgeByWaterTactic) IsTriggerPrepare() bool {
	return false
}

func (a BrokenBridgeByWaterTactic) SetTriggerPrepare(triggerPrepare bool) {
}
