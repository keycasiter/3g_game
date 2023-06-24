package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 虎豹骑
// 将骑兵进阶为天下骁锐的虎豹骑：
// 我军全体提高40武力，战斗前3回合，我军全体突击战法发动率提高10%，
// 若曹纯统领时，提升的发动概率额外受武力影响
type TigerAndLeopardCavalry struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerAndLeopardCavalry) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerAndLeopardCavalry) Prepare() {
}

func (t TigerAndLeopardCavalry) Id() consts.TacticId {
	return consts.TigerAndLeopardCavalry
}

func (t TigerAndLeopardCavalry) Name() string {
	return "虎豹骑"
}

func (t TigerAndLeopardCavalry) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TigerAndLeopardCavalry) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerAndLeopardCavalry) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerAndLeopardCavalry) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (t TigerAndLeopardCavalry) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (t TigerAndLeopardCavalry) Execute() {
}

func (t TigerAndLeopardCavalry) IsTriggerPrepare() bool {
	return false
}
