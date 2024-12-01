package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 强攻
// 使自己进入连击（每回合可以普通攻击2次）状态，持续1回合
type TakeByStormTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeByStormTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TakeByStormTactic) Prepare() {
}

func (t TakeByStormTactic) Id() consts.TacticId {
	return consts.TakeByStorm
}

func (t TakeByStormTactic) Name() string {
	return "强攻"
}

func (t TakeByStormTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeByStormTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeByStormTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeByStormTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TakeByStormTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeByStormTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	// 使自己进入连击（每回合可以普通攻击2次）状态，持续1回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ContinuousAttack,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TakeByStormTactic) IsTriggerPrepare() bool {
	return false
}

func (a TakeByStormTactic) SetTriggerPrepare(triggerPrepare bool) {
}
