package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 包扎
// 治疗随机我军单体（治疗率160%，受智力影响）
// 主动，30%
type DressTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DressTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.3
	return d
}

func (d DressTactic) Prepare() {

}

func (d DressTactic) Id() consts.TacticId {
	return consts.Dress
}

func (d DressTactic) Name() string {
	return "包扎"
}

func (d DressTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DressTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DressTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DressTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DressTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DressTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral
	//治疗随机我军单体（治疗率160%，受智力影响）
	pairGeneral := util.GetPairOneGeneral(d.tacticsParams, currentGeneral)
	resume := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.6)
	util.ResumeSoldierNum(ctx, pairGeneral, resume)
}

func (d DressTactic) IsTriggerPrepare() bool {
	return false
}
