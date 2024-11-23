package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 至柔动刚
// 战斗中，敌我全体普通攻击伤害降低35%，自身受到伤害时有50%概率（受智力影响）偷取伤害来源10点属性（智力、统率、速度随机一种，受智力影响，可叠加，持续5回合）
// 我军全体普通攻击后，自身有50%概率（受智力影响）对敌军单体造成一次谋略伤害（伤害率106%，受智力影响）
type ZhiRouDongGangTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a ZhiRouDongGangTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a ZhiRouDongGangTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

}

func (a ZhiRouDongGangTactic) Id() consts.TacticId {
	return consts.ZhiRouDongGang
}

func (a ZhiRouDongGangTactic) Name() string {
	return "至柔动刚"
}

func (a ZhiRouDongGangTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a ZhiRouDongGangTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a ZhiRouDongGangTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a ZhiRouDongGangTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a ZhiRouDongGangTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a ZhiRouDongGangTactic) Execute() {
}

func (a ZhiRouDongGangTactic) IsTriggerPrepare() bool {
	return false
}
