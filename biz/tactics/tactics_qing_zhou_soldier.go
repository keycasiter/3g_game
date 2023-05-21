package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//青州兵
//将枪兵进阶为冲坚毁锐的青州兵：
//战斗2回合，使我军群体（2人）受到普通攻击时对攻击者进行一次反击（伤害率72%），
//第3回合，回合开始时依次为我军全体恢复兵力，优先完全恢复我军兵力最低单体，再恢复我军其他单体
//（总治疗率180%，受武力影响，额外受敌军造成伤害影响）
// 若曹操统领，治疗效果额外受统率影响
type QingZhouSoldierTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (q QingZhouSoldierTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	q.tacticsParams = tacticsParams
	q.triggerRate = 1.0
	return q
}

func (q QingZhouSoldierTactic) Prepare() {

}

func (q QingZhouSoldierTactic) Id() consts.TacticId {
	return consts.QingZhouSoldier
}

func (q QingZhouSoldierTactic) Name() string {
	return "青州兵"
}

func (q QingZhouSoldierTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (q QingZhouSoldierTactic) GetTriggerRate() float64 {
	return q.triggerRate
}

func (q QingZhouSoldierTactic) SetTriggerRate(rate float64) {
	q.triggerRate = rate
}

func (q QingZhouSoldierTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (q QingZhouSoldierTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (q QingZhouSoldierTactic) Execute() {

}

func (q QingZhouSoldierTactic) IsTriggerPrepare() bool {
	return false
}
