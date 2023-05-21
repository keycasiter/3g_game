package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//卧薪尝胆
//对敌军群体（2人）发动一次兵刃攻击（伤害率96%），
//并有25几率（根据自身连击、洞察、先攻、必中、破阵的状态数，每多一种提高10%几率）造成震慑（无法行动）状态，持续1回合
type SleepOnTheBrushwoodAndTasteTheGallTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.45
	return s
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Prepare() {
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Id() consts.TacticId {
	return consts.SleepOnTheBrushwoodAndTasteTheGall
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Name() string {
	return "卧薪尝胆"
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Execute() {
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) IsTriggerPrepare() bool {
	return false
}
