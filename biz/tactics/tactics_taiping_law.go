package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 太平道法
// 获得28%奇谋并提高自带主动战法发动率(6%，若为准备战法则提高12%，受智力影响)，
// 自身为黄巾军主将时，使黄巾军副将同样获得自带战法发动率提升
type TaipingLawTactic struct {
	tacticsParams *model.TacticsParams
}

func (t TaipingLawTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t TaipingLawTactic) Prepare() {

}

func (t TaipingLawTactic) Id() consts.TacticId {
	return consts.TaipingLaw
}

func (t TaipingLawTactic) Name() string {
	return "太平道法"
}

func (t TaipingLawTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TaipingLawTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TaipingLawTactic) Execute() {
	return
}
