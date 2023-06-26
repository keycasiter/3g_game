package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 追伤
// 普通攻击之后，对攻击目标造成禁疗，持续2回合
// 突击，40%
type ChasingInjuryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChasingInjuryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.4
	return c
}

func (c ChasingInjuryTactic) Prepare() {

}

func (c ChasingInjuryTactic) Id() consts.TacticId {
	return consts.ChasingInjury
}

func (c ChasingInjuryTactic) Name() string {
	return "追伤"
}

func (c ChasingInjuryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c ChasingInjuryTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c ChasingInjuryTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c ChasingInjuryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (c ChasingInjuryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ChasingInjuryTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral
	sufferGeneral := c.tacticsParams.CurrentSufferGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	// 普通攻击之后，对攻击目标造成禁疗，持续2回合
	if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     c.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_ProhibitionTreatment,
				TacticId:   c.Id(),
			})

			return revokeResp
		})
	}
}

func (c ChasingInjuryTactic) IsTriggerPrepare() bool {
	return false
}
