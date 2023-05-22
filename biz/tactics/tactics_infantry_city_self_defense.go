package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 婴城自守
// 恢复我军群体（2人）兵力（治疗率92%，收智力影响），
// 并使其获得休整状态（每回合恢复一次兵力，治疗率62%），持续1回合
type InfantryCitySelfDefenseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InfantryCitySelfDefenseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.5
	return i
}

func (i InfantryCitySelfDefenseTactic) Prepare() {
}

func (i InfantryCitySelfDefenseTactic) Id() consts.TacticId {
	return consts.InfantryCitySelfDefense
}

func (i InfantryCitySelfDefenseTactic) Name() string {
	return "婴城自守"
}

func (i InfantryCitySelfDefenseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (i InfantryCitySelfDefenseTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i InfantryCitySelfDefenseTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i InfantryCitySelfDefenseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InfantryCitySelfDefenseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InfantryCitySelfDefenseTactic) Execute() {

}

func (i InfantryCitySelfDefenseTactic) IsTriggerPrepare() bool {
	return false
}
