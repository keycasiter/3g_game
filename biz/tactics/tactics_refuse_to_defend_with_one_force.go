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

// 一力拒守
// 恢复自身兵力(治疗率268%)并提高21统率，最多叠加2次，持续3回合；
// 自身为副将时，兵力恢复受武力影响
type RefuseToDefendWithOneForceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RefuseToDefendWithOneForceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.55
	return r
}

func (r RefuseToDefendWithOneForceTactic) Prepare() {
}

func (r RefuseToDefendWithOneForceTactic) Id() consts.TacticId {
	return consts.RefuseToDefendWithOneForce
}

func (r RefuseToDefendWithOneForceTactic) Name() string {
	return "一力拒守"
}

func (r RefuseToDefendWithOneForceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RefuseToDefendWithOneForceTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RefuseToDefendWithOneForceTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RefuseToDefendWithOneForceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RefuseToDefendWithOneForceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RefuseToDefendWithOneForceTactic) Execute() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)

	//恢复自身兵力(治疗率268%)并提高21统率，最多叠加2次，持续3回合；
	//自身为副将时，兵力恢复受武力影响
	resumeRate := 2.68
	if !currentGeneral.IsMaster {
		resumeRate += currentGeneral.BaseInfo.AbilityAttr.ForceBase / 100 / 100
	}
	resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * resumeRate)
	util.ResumeSoldierNum(ctx, currentGeneral, resumeNum)

	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    21,
		EffectRound:    3,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     r.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrCommand,
				TacticId:   r.Id(),
			})

			return revokeResp
		})
	}
}

func (r RefuseToDefendWithOneForceTactic) IsTriggerPrepare() bool {
	return false
}
