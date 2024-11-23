package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 蕙质兰心
// 战斗中，我军全体统率提升20（受各自智力影响），且自身获得7层兰心效果，每层使自身造成和受到伤害降低10%；
// 自身受到伤害时，消耗1层兰心效果并有50%概率（受智力影响）治疗我军单体（治疗率222%，受智力影响）
// 兰心首次降至4层时，敌我全体造成的兵刃伤害降低20%（受智力影响）
type HuiZhiLanXinTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a HuiZhiLanXinTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a HuiZhiLanXinTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
}

func (a HuiZhiLanXinTactic) Id() consts.TacticId {
	return consts.HuiZhiLanXin
}

func (a HuiZhiLanXinTactic) Name() string {
	return "蕙质兰心"
}

func (a HuiZhiLanXinTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a HuiZhiLanXinTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a HuiZhiLanXinTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a HuiZhiLanXinTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a HuiZhiLanXinTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a HuiZhiLanXinTactic) Execute() {
}

func (a HuiZhiLanXinTactic) IsTriggerPrepare() bool {
	return false
}
