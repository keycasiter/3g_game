package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//锦帆军
//将弓兵进阶为气盖千夫的锦帆军：
//部队普通攻击时，有45%概率使目标进入溃逃状态（伤害率64%，受武力影响），持续2回合
//若目标已经溃逃则造成兵刃攻击（伤害率110%）并恢复伤害量的30%兵力；
//若甘宁统领，提高友军6%会心
type JinFanArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JinFanArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JinFanArmyTactic) Prepare() {

}

func (j JinFanArmyTactic) Id() consts.TacticId {
	return consts.JinFanArmy
}

func (j JinFanArmyTactic) Name() string {
	return "锦帆军"
}

func (j JinFanArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (j JinFanArmyTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JinFanArmyTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JinFanArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (j JinFanArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (j JinFanArmyTactic) Execute() {
}

func (j JinFanArmyTactic) IsTriggerPrepare() bool {
	return false
}
