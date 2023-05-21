package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//忠勇义烈
//战斗中，自身每回合有60%概率获得以下效果：
//主动战法发动率提升6%（受武力影响）
//武力、统率、智力提升45；28%倒戈，持续1回合，每种效果独立判定
type LoyalAndBraveMartyrs struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LoyalAndBraveMartyrs) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LoyalAndBraveMartyrs) Prepare() {
}

func (l LoyalAndBraveMartyrs) Id() consts.TacticId {
	return consts.LoyalAndBraveMartyrs
}

func (l LoyalAndBraveMartyrs) Name() string {
	return "忠勇义烈"
}

func (l LoyalAndBraveMartyrs) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LoyalAndBraveMartyrs) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LoyalAndBraveMartyrs) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LoyalAndBraveMartyrs) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (l LoyalAndBraveMartyrs) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LoyalAndBraveMartyrs) Execute() {
}

func (l LoyalAndBraveMartyrs) IsTriggerPrepare() bool {
	return false
}
