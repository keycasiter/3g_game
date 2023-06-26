package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 白衣渡江
// 战斗首回合我军全体获得1次抵御
// 战斗中，自身造成兵刃伤害时有50%概率使敌军单体缴械，持续2回合，
// 造成谋略伤害时有50%概率使敌军单体计穷，持续1回合
// 该效果每回合最多分别触发一次。
// 指挥，100%
type CrossingTheRiverInWhiteClothesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CrossingTheRiverInWhiteClothesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CrossingTheRiverInWhiteClothesTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral
	triggerRoundHolderMap := map[consts.DebuffEffectType]map[consts.BattleRound]bool{}

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	pairGenerals := util.GetPairGeneralArr(c.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		// 战斗首回合我军全体获得1次抵御
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
			EffectTimes:    1,
			FromTactic:     c.Id(),
			ProduceGeneral: currentGeneral,
		})
		// 战斗中，自身造成兵刃伤害时有50%概率使敌军单体缴械，持续2回合
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_WeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRound := params.CurrentRound

			if util.GenerateRate(0.5) && !triggerRoundHolderMap[consts.DebuffEffectType_CancelWeapon][triggerRound] {
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, c.tacticsParams)
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     c.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_CancelWeapon,
							TacticId:   c.Id(),
						})

						return revokeResp
					})
				}
				triggerRoundHolderMap[consts.DebuffEffectType_CancelWeapon][triggerRound] = true
			}

			return triggerResp
		})
		// 造成谋略伤害时有50%概率使敌军单体计穷，持续1回合
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_StrategyDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRound := params.CurrentRound

			if util.GenerateRate(0.5) && !triggerRoundHolderMap[consts.DebuffEffectType_NoStrategy][triggerRound] {
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, c.tacticsParams)
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
					EffectRound:    1,
					FromTactic:     c.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_NoStrategy,
							TacticId:   c.Id(),
						})

						return revokeResp
					})
				}
				triggerRoundHolderMap[consts.DebuffEffectType_CancelWeapon][triggerRound] = true
			}

			return triggerResp
		})
	}

}

func (c CrossingTheRiverInWhiteClothesTactic) Id() consts.TacticId {
	return consts.CrossingTheRiverInWhiteClothes
}

func (c CrossingTheRiverInWhiteClothesTactic) Name() string {
	return "白衣渡江"
}

func (c CrossingTheRiverInWhiteClothesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c CrossingTheRiverInWhiteClothesTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CrossingTheRiverInWhiteClothesTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CrossingTheRiverInWhiteClothesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c CrossingTheRiverInWhiteClothesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CrossingTheRiverInWhiteClothesTactic) Execute() {

}

func (c CrossingTheRiverInWhiteClothesTactic) IsTriggerPrepare() bool {
	return false
}
