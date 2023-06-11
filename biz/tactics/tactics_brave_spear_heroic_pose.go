package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 奋矛英姿
// 每次普通攻击后，降低敌方15统率，将其转化为自身的武力，直到战斗结束，每第4次普通攻击时，伤害提高100%并对敌军全体造成伤害
// 被动 100%
type BraveSpearHeroicPoseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BraveSpearHeroicPoseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BraveSpearHeroicPoseTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//注册效果器
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		sufferGenerl := b.tacticsParams.CurrentSufferGeneral
		triggerGeneral := b.tacticsParams.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//每次普通攻击后，降低敌方15统率，将其转化为自身的武力，直到战斗结束
		util.DebuffEffectWrapSet(ctx, sufferGenerl, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
			EffectValue: 15,
			FromTactic:  b.Id(),
		})
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue: 15,
			FromTactic:  b.Id(),
		})

		return triggerResp
	})
	//每第4次普通攻击时，伤害提高100%并对敌军全体造成伤害
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if triggerGeneral.ExecuteGeneralAttckNum != 0 && triggerGeneral.ExecuteGeneralAttckNum%4 == 0 {
			//普通攻击伤害提高
			if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:  1.0,
				EffectRound: 1,
				FromTactic:  b.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
						TacticId:   b.Id(),
					})

					return revokeResp
				})
			}
			//获得群攻效果
			if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
				EffectRate:  1.0,
				EffectRound: 1,
				FromTactic:  b.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_GroupAttack,
						TacticId:   b.Id(),
					})

					return revokeResp
				})
			}

			triggerResp.IsTerminate = true
		}

		return triggerResp
	})
}

func (b BraveSpearHeroicPoseTactic) Id() consts.TacticId {
	return consts.BraveSpearHeroicPose
}

func (b BraveSpearHeroicPoseTactic) Name() string {
	return "奋矛英姿"
}

func (b BraveSpearHeroicPoseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BraveSpearHeroicPoseTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BraveSpearHeroicPoseTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BraveSpearHeroicPoseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BraveSpearHeroicPoseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveSpearHeroicPoseTactic) Execute() {
}

func (b BraveSpearHeroicPoseTactic) IsTriggerPrepare() bool {
	return false
}
