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

// 破军威胜
// 降低敌军单体70点统率（受武力影响），持续2回合，并对其造成兵刃伤害（伤害率228%）
type BreakingThroughTheArmyAndWinningVictoriesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Prepare() {
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheArmyAndWinningVictories
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Name() string {
	return "破军威胜"
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	//降低敌军单体70点统率（受武力影响），持续2回合，并对其造成兵刃伤害（伤害率228%）
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneral(b.tacticsParams)

	//施加效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectRate:  1.0,
		EffectRound: 2,
		FromTactic:  b.Id(),
	}).IsSuccess {
		//降低统率
		util.DeduceGeneralAttr(enemyGeneral, consts.AbilityAttr_Command, 70)
		hlog.CtxInfof(ctx, "[%s]的统率降低了%.2d",
			enemyGeneral.BaseInfo.Name,
			70)
		//注册消失效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   b.Id(),
				CostOverTriggerFunc: func() {
					util.DeduceGeneralAttr(revokeGeneral, consts.AbilityAttr_Command, 70)
					hlog.CtxInfof(ctx, "[%s]的统率提升了%.2d",
						revokeGeneral.BaseInfo.Name,
						70)
				},
			})

			return revokeResp
		})
		//并对其造成兵刃伤害（伤害率228%）
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.28)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: b.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticName:    b.Name(),
		})
	}
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) IsTriggerPrepare() bool {
	return false
}
