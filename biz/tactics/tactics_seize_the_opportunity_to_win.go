package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//临机制胜
//对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率120%，受智力影响），持续2回合
//若敌军已有中毒状态，则使其随机获得灼烧（受智力影响）、叛逃（受武力活智力最高一项影响，无视防御）、沙暴（受智力影响）状态中的一种
//每回合持续造成伤害（伤害率120%），持续2回合，该战法发动后进入1回合冷却
type SeizeTheOpportunityToWinTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SeizeTheOpportunityToWinTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.55
	return s
}

func (s SeizeTheOpportunityToWinTactic) Prepare() {
}

func (s SeizeTheOpportunityToWinTactic) Id() consts.TacticId {
	return consts.SeizeTheOpportunityToWin
}

func (s SeizeTheOpportunityToWinTactic) Name() string {
	return "临机制胜"
}

func (s SeizeTheOpportunityToWinTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SeizeTheOpportunityToWinTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SeizeTheOpportunityToWinTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SeizeTheOpportunityToWinTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SeizeTheOpportunityToWinTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SeizeTheOpportunityToWinTactic) Execute() {
}

func (s SeizeTheOpportunityToWinTactic) IsTriggerPrepare() bool {
	return false
}
