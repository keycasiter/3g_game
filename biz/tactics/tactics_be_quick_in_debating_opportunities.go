package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 才辩机捷
// 使自身施加的灼烧、水攻、中毒、溃逃、沙暴、叛逃状态伤害提高90%，休整和急救的恢复量提升30%
type BeQuickInDebatingOpportunitiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BeQuickInDebatingOpportunitiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BeQuickInDebatingOpportunitiesTactic) Prepare() {
	currentGeneral := b.tacticsParams.CurrentGeneral

	//使自身施加的灼烧、水攻、中毒、溃逃、沙暴、叛逃状态伤害提高90%，休整和急救的恢复量提升30%
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_DebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		return triggerResp
	})
}

func (b BeQuickInDebatingOpportunitiesTactic) Id() consts.TacticId {
	return consts.BeQuickInDebatingOpportunities
}

func (b BeQuickInDebatingOpportunitiesTactic) Name() string {
	return "才辩机捷"
}

func (b BeQuickInDebatingOpportunitiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BeQuickInDebatingOpportunitiesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeQuickInDebatingOpportunitiesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeQuickInDebatingOpportunitiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BeQuickInDebatingOpportunitiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeQuickInDebatingOpportunitiesTactic) Execute() {
}

func (b BeQuickInDebatingOpportunitiesTactic) IsTriggerPrepare() bool {
	return false
}

func (a BeQuickInDebatingOpportunitiesTactic) SetTriggerPrepare(triggerPrepare bool) {
}
