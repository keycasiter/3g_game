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

// 兵无常势
// 战斗中，自己累积进行第3次普通攻击时，对攻击目标造成谋略伤害（伤害率240%，受智力影响），
// 并治疗自己（治疗率180%，受智力影响）
type WorkOutMeasuresToSuitLocalConditionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)

	//战斗中，自己累积进行第3次普通攻击时，对攻击目标造成谋略伤害（伤害率240%，受智力影响），并治疗自己（治疗率180%，受智力影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := w.tacticsParams.CurrentSufferGeneral

		if triggerGeneral.ExecuteGeneralAttackNum > 0 &&
			triggerGeneral.ExecuteGeneralAttackNum%3 == 0 {
			//攻击
			dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 2.4
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     w.tacticsParams,
				AttackGeneral:     triggerGeneral,
				SufferGeneral:     sufferGeneral,
				DamageType:        consts.DamageType_Strategy,
				DamageImproveRate: dmgRate,
				TacticId:          w.Id(),
				TacticName:        w.Name(),
			})
			//治疗
			resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.8)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  w.tacticsParams,
				ProduceGeneral: triggerGeneral,
				SufferGeneral:  triggerGeneral,
				ResumeNum:      resumeNum,
				TacticId:       w.Id(),
			})
		}

		return triggerResp
	})
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Id() consts.TacticId {
	return consts.WorkOutMeasuresToSuitLocalConditions
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Name() string {
	return "兵无常势"
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Execute() {
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) IsTriggerPrepare() bool {
	return false
}

func (a WorkOutMeasuresToSuitLocalConditionsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
