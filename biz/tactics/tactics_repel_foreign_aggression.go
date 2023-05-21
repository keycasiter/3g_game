package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//折冲御侮
//普通攻击之后，使随机敌军单体降低100点统率和智力，持续2回合；
//若自己不是主将，则额外使我军主将获得2次抵御，可免疫伤害，持续2回合
//45%
type RepelForeignAggressionTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RepelForeignAggressionTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.45
	return r
}

func (r RepelForeignAggressionTactic) Prepare() {
}

func (r RepelForeignAggressionTactic) Id() consts.TacticId {
	return consts.RepelForeignAggression
}

func (r RepelForeignAggressionTactic) Name() string {
	return "折冲御侮"
}

func (r RepelForeignAggressionTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (r RepelForeignAggressionTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RepelForeignAggressionTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RepelForeignAggressionTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (r RepelForeignAggressionTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RepelForeignAggressionTactic) Execute() {

}

func (r RepelForeignAggressionTactic) IsTriggerPrepare() bool {
	return false
}
