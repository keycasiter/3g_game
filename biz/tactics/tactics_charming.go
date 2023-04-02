package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：魅惑
// 战法描述：自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
// 持续1回合，自身为女性时，触发几率额外受智力影响
type CharmingTactic struct {
	tacticsParams model.TacticsParams
}

func (c CharmingTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
}

func (c CharmingTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	sufferGeneralName := c.tacticsParams.SufferGeneral.BaseInfo.Name
	if !util.GenerateRate(0.45) {
		hlog.CtxInfof(ctx, "[%s]执行来自[%s][%s]的「魅惑」效果因几率没有生效",
			sufferGeneralName,
			sufferGeneralName,
			c.Name(),
		)
		return
	}
}

func (c CharmingTactic) Name() string {
	return "魅惑"
}

func (c CharmingTactic) Id() int64 {
	return Charming
}

func (c CharmingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c CharmingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CharmingTactic) TriggerRate() float64 {
	return 1.0
}
