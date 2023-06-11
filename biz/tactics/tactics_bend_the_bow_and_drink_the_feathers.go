package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 弯弓饮羽
// 普通攻击之后，使攻击目标降低150点统率，并造成计穷（无法发动主动战法）状态，持续1回合
type BendTheBowAndDrinkTheFeathersTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BendTheBowAndDrinkTheFeathersTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BendTheBowAndDrinkTheFeathersTactic) Prepare() {

}

func (b BendTheBowAndDrinkTheFeathersTactic) Id() consts.TacticId {
	return consts.BendTheBowAndDrinkTheFeathers
}

func (b BendTheBowAndDrinkTheFeathersTactic) Name() string {
	return "弯弓饮羽"
}

func (b BendTheBowAndDrinkTheFeathersTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BendTheBowAndDrinkTheFeathersTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BendTheBowAndDrinkTheFeathersTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BendTheBowAndDrinkTheFeathersTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (b BendTheBowAndDrinkTheFeathersTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BendTheBowAndDrinkTheFeathersTactic) Execute() {
	//普通攻击之后，使攻击目标降低150点统率，并造成计穷（无法发动主动战法）状态，持续1回合
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	sufferGeneral := b.tacticsParams.CurrentSufferGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//使攻击目标降低150点统率
	util.DeduceGeneralAttr(sufferGeneral, consts.AbilityAttr_Command, 150)
	hlog.CtxInfof(ctx, "[%s]的统率降低了%.2f",
		sufferGeneral.BaseInfo.Name,
		150)
	//造成计穷（无法发动主动战法）状态，持续1回合
	if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
		EffectRate:  1.0,
		EffectRound: 1,
		FromTactic:  b.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    sufferGeneral,
				EffectType: consts.DebuffEffectType_NoStrategy,
				TacticId:   b.Id(),
				CostOverTriggerFunc: func() {
					util.ImproveGeneralAttr(revokeGeneral, consts.AbilityAttr_Command, 150)
					hlog.CtxInfof(ctx, "[%s]的统率提高了%.2f",
						revokeGeneral.BaseInfo.Name,
						150)
				},
			})
			return revokeResp
		})
	}
}

func (b BendTheBowAndDrinkTheFeathersTactic) IsTriggerPrepare() bool {
	return false
}
