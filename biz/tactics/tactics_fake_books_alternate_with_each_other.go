package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//伪书相间
//对敌军单体造成谋略伤害（伤害率206%，受智力影响）
//若目标处于混乱状态则使目标对其友军单体发动攻击（伤害率186%，类型取决于目标武力、智力较高的一项），
//否则施加混乱（攻击和战法无差别选择目标）状态，持续1回合
type FakeBooksAlternateWithEachOtherTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FakeBooksAlternateWithEachOtherTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.45
	return f
}

func (f FakeBooksAlternateWithEachOtherTactic) Prepare() {

}

func (f FakeBooksAlternateWithEachOtherTactic) Id() consts.TacticId {
	return consts.FakeBooksAlternateWithEachOther
}

func (f FakeBooksAlternateWithEachOtherTactic) Name() string {
	return "伪书相间"
}

func (f FakeBooksAlternateWithEachOtherTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FakeBooksAlternateWithEachOtherTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FakeBooksAlternateWithEachOtherTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FakeBooksAlternateWithEachOtherTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FakeBooksAlternateWithEachOtherTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FakeBooksAlternateWithEachOtherTactic) Execute() {
}

func (f FakeBooksAlternateWithEachOtherTactic) IsTriggerPrepare() bool {
	return false
}
