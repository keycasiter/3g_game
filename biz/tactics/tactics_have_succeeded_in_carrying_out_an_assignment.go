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

// 不辱使命
// 对敌军单体造成一次兵刃攻击（伤害率220%），并有30%概率施加震慑状态（无法行动），持续1回合
type HaveSucceededInCarryingOutAnAssignmentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Prepare() {

}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Id() consts.TacticId {
	return consts.HaveSucceededInCarryingOutAnAssignment
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Name() string {
	return "不辱使命"
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	// 对敌军单体造成一次兵刃攻击（伤害率220%），并有30%概率施加震慑状态（无法行动），持续1回合
	// 找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, h.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.2)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: h.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    h.Name(),
	})
	//施加状态
	if util.GenerateRate(0.3) {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     h.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Awe,
					TacticId:   h.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) IsTriggerPrepare() bool {
	return false
}
