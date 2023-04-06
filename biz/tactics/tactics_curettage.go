package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 战法名称：刮骨疗毒
// 战法描述：为损失兵力最多的我军单体清除负面状态并为其恢复兵力（治疗率256%，受智力影响）
type CurettageTactic struct {
	tacticsParams *model.TacticsParams
}

func (c CurettageTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	return c
}

func (c CurettageTactic) Prepare() {
	return
}

func (c CurettageTactic) Name() string {
	return "刮骨疗毒"
}

func (c CurettageTactic) Execute() {
	currentGeneral := c.tacticsParams.CurrentGeneral
	//找我我军损失兵力最多的武将
	maxLossSoldierNumGeneral := util.GetPairMaxLossSoldierNumGeneral(c.tacticsParams)

	//清除负面状态
	maxLossSoldierNumGeneral.DeBuffEffectHolderMap = map[consts.DebuffEffectType]float64{}
	//清除负面触发器
	maxLossSoldierNumGeneral.DeBuffEffectTriggerMap = map[consts.DebuffEffectType]map[consts.BattleRound]float64{}

	//为其恢复兵力（治疗率256%，受智力影响）
	maxLossSoldierNumGeneral.SoldierNum += cast.ToInt64(2.56 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
}

func (c CurettageTactic) Trigger() {
	return
}

func (c CurettageTactic) Id() consts.TacticId {
	return consts.Curettage
}

func (c CurettageTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c CurettageTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
