package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 密计诛逆
// 战斗中，我军主将造成伤害（大于300）后，有50%概率使敌军单体造成最终伤害降低15%（受统率影响，持续2回合，可叠加3次）
// 且前2回合有50%概率（受统率影响）令主将对敌军全体造成1次兵刃伤害（伤害率25%，受主将统率影响）；
// 战斗第6回合，自身对敌军单体造成斩杀伤害（伤害率100% +25%*最终降伤施加次数）
type MiJiZhuNiTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a MiJiZhuNiTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a MiJiZhuNiTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

}

func (a MiJiZhuNiTactic) Id() consts.TacticId {
	return consts.MiJiZhuNi
}

func (a MiJiZhuNiTactic) Name() string {
	return "密计诛逆"
}

func (a MiJiZhuNiTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a MiJiZhuNiTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a MiJiZhuNiTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a MiJiZhuNiTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a MiJiZhuNiTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a MiJiZhuNiTactic) Execute() {
}

func (a MiJiZhuNiTactic) IsTriggerPrepare() bool {
	return false
}
