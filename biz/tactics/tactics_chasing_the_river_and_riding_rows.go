package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 溯江摇橹
// 战斗中，每回合有50%概率净化自身，自身造成伤害时，有35%概率使随机敌军单体进入计穷或震慑状态，持续1回合，每回合最多触发一次，
// 受到伤害时，有60%概率对敌军单体造成对应类型的伤害（伤害率50%，受对应属性影响）
// 自身为主将时，净化的触发概率提高10%
// 被动，100%
type ChasingTheRiverAndRidingRowsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChasingTheRiverAndRidingRowsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c ChasingTheRiverAndRidingRowsTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral
	triggerRoundHolderMap := map[consts.BattleRound]bool{}

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	// 战斗中，每回合有50%概率净化自身，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRate := 0.5

		// 自身为主将时，净化的触发概率提高10%
		if triggerGeneral.IsMaster {
			triggerRate += 0.1
		}
		if util.GenerateRate(triggerRate) {
			util.DebuffEffectClean(ctx, triggerGeneral)
		}

		return triggerResp
	})
	//自身造成伤害时，有35%概率使随机敌军单体进入计穷或震慑状态，持续1回合，每回合最多触发一次，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if util.GenerateRate(0.35) && !triggerRoundHolderMap[triggerRound] {
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, c.tacticsParams)
			debuffs := []consts.DebuffEffectType{
				consts.DebuffEffectType_NoStrategy,
				consts.DebuffEffectType_Awe,
			}
			hitIdx := util.GenerateHitOneIdx(len(debuffs))
			debuff := debuffs[hitIdx]
			//施加效果
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, debuff, &vo.EffectHolderParams{
				EffectRound:    1,
				FromTactic:     c.Id(),
				ProduceGeneral: nil,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: debuff,
						TacticId:   c.Id(),
					})

					return revokeResp
				})
			}
			triggerRoundHolderMap[triggerRound] = true
		}

		return triggerResp
	})
	// 受到伤害时，有60%概率对敌军单体造成对应类型的伤害（伤害率50%，受对应属性影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		dmgType := params.DamageType

		if util.GenerateRate(0.6) {
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, c.tacticsParams)
			dmgRate := float64(0)
			switch dmgType {
			case consts.DamageType_Weapon:
				dmgRate = triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100/100 + 0.5
			case consts.DamageType_Strategy:
				dmgRate = triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.5
			}

			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     c.tacticsParams,
				AttackGeneral:     triggerGeneral,
				SufferGeneral:     enemyGeneral,
				DamageType:        dmgType,
				DamageImproveRate: dmgRate,
				TacticId:          c.Id(),
				TacticName:        c.Name(),
			})
		}

		return triggerResp
	})
}

func (c ChasingTheRiverAndRidingRowsTactic) Id() consts.TacticId {
	return consts.ChasingTheRiverAndRidingRows
}

func (c ChasingTheRiverAndRidingRowsTactic) Name() string {
	return "溯江摇橹"
}

func (c ChasingTheRiverAndRidingRowsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ChasingTheRiverAndRidingRowsTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c ChasingTheRiverAndRidingRowsTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c ChasingTheRiverAndRidingRowsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c ChasingTheRiverAndRidingRowsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ChasingTheRiverAndRidingRowsTactic) Execute() {

}

func (c ChasingTheRiverAndRidingRowsTactic) IsTriggerPrepare() bool {
	return false
}
