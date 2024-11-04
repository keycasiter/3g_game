package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 长者之风
// 使我军全体武力、智力提高28点
type TheWindOfTheElderlyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheWindOfTheElderlyTactic) IsTriggerPrepare() bool {
	return false
}

func (t TheWindOfTheElderlyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TheWindOfTheElderlyTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	pairGenerals := util.GetPairGeneralArr(currentGeneral, t.tacticsParams)
	for _, general := range pairGenerals {
		general.BaseInfo.AbilityAttr.ForceBase += 28
		general.BaseInfo.AbilityAttr.IntelligenceBase += 28
		hlog.CtxInfof(ctx, "[%s]的武力提高了28点",
			general.BaseInfo.Name,
		)
		hlog.CtxInfof(ctx, "[%s]的智力提高了28点",
			general.BaseInfo.Name,
		)
	}
}

func (t TheWindOfTheElderlyTactic) Id() consts.TacticId {
	return consts.TheWindOfTheElderly
}

func (t TheWindOfTheElderlyTactic) Name() string {
	return "长者之风"
}

func (t TheWindOfTheElderlyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheWindOfTheElderlyTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheWindOfTheElderlyTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheWindOfTheElderlyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TheWindOfTheElderlyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheWindOfTheElderlyTactic) Execute() {
}
