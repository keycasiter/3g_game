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

// 持军毅重
// 使自己下次受到兵刃伤害提高40%，持续1回合，提高自身56点武力，持续3回合，并对敌军群体（2人）造成猛烈一击（伤害率184%）
// 主动 35%
type HoldTheArmyWithDeterminationAndDeterminationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.35
	return h
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Prepare() {
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Id() consts.TacticId {
	return consts.HoldTheArmyWithDeterminationAndDetermination
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Name() string {
	return "持军毅重"
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	//兵刃伤害提升
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRound:    1,
		EffectRate:     0.4,
		FromTactic:     h.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
				TacticId:   h.Id(),
			})

			return revokeResp
		})
	}
	//武力提升
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectValue:    56,
		EffectRound:    3,
		FromTactic:     h.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrForce,
				TacticId:   h.Id(),
			})

			return revokeResp
		})
	}
	//攻击
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, h.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     h.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Weapon,
			DamageImproveRate: 1.84,
			TacticId:          h.Id(),
			TacticName:        h.Name(),
		})
	}
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) IsTriggerPrepare() bool {
	return false
}
