package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 太平道法
// 获得28%奇谋并提高自带主动战法发动率(6%，若为准备战法则提高12%，受智力影响)，
// 自身为黄巾军主将时，使黄巾军副将同样获得自带战法发动率提升
type TaipingLawTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TaipingLawTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TaipingLawTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TaipingLawTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TaipingLawTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TaipingLawTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//获得28%奇谋并提高自带主动战法发动率(6%，若为准备战法则提高12%，受智力影响)，
	util.BuffEffectWrapSet(currentGeneral, consts.BuffEffectType_EnhanceStrategy, 0.28)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
			triggerGeneral.BaseInfo.Name,
			t.Name(),
			consts.BuffEffectType_HuangTianDangLi,
		)

		return triggerResp
	})
	//TODO 自身为黄巾军主将时，使黄巾军副将同样获得自带战法发动率提升
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
