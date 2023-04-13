package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	triggerRate   float64
}

func (c CurettageTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CurettageTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CurettageTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CurettageTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.4
	return c
}

func (c CurettageTactic) Prepare() {
	return
}

func (c CurettageTactic) Name() string {
	return "刮骨疗毒"
}

func (c CurettageTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//找我我军损失兵力最多的武将
	maxLossSoldierNumGeneral := util.GetPairMaxLossSoldierNumGeneral(c.tacticsParams)

	//清除负面状态
	util.DebuffEffectClean(ctx, maxLossSoldierNumGeneral)

	//为其恢复兵力（治疗率256%，受智力影响）
	resumeNum := cast.ToInt64(2.56 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
	hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
		maxLossSoldierNumGeneral.BaseInfo.Name,
		resumeNum,
		maxLossSoldierNumGeneral.SoldierNum,
		maxLossSoldierNumGeneral.SoldierNum+resumeNum,
	)
	maxLossSoldierNumGeneral.SoldierNum += resumeNum
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
