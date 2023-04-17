package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//诈降
//战斗开始后首回合，自己获得混乱（攻击和战法无差别选择目标）状态，第2回合起获得休整状态（每回合恢复一次兵力，治疗率80%，受智力影响），持续3回合
type PretendToSurrenderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PretendToSurrenderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p PretendToSurrenderTactic) Prepare() {
	//战斗开始后首回合，自己获得混乱（攻击和战法无差别选择目标）状态

	//第2回合起获得休整状态（每回合恢复一次兵力，治疗率80%，受智力影响），持续3回合
}

func (p PretendToSurrenderTactic) Id() consts.TacticId {
	return consts.PretendToSurrender
}

func (p PretendToSurrenderTactic) Name() string {
	return "诈降"
}

func (p PretendToSurrenderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (p PretendToSurrenderTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PretendToSurrenderTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PretendToSurrenderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (p PretendToSurrenderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PretendToSurrenderTactic) Execute() {

}

func (p PretendToSurrenderTactic) IsTriggerPrepare() bool {
	return false
}
