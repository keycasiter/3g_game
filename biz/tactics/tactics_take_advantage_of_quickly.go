package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 速乘其利
// 普通攻击之后，对目标发动一次兵刃攻击（伤害率185%）并计穷（无法发动主动战法）1回合
type TakeAdvantageOfQuicklyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeAdvantageOfQuicklyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TakeAdvantageOfQuicklyTactic) Prepare() {
}

func (t TakeAdvantageOfQuicklyTactic) Id() consts.TacticId {
	return consts.TakeAdvantageOfQuickly
}

func (t TakeAdvantageOfQuicklyTactic) Name() string {
	return "速乘其利"
}

func (t TakeAdvantageOfQuicklyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeAdvantageOfQuicklyTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeAdvantageOfQuicklyTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeAdvantageOfQuicklyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t TakeAdvantageOfQuicklyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeAdvantageOfQuicklyTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//普通攻击之后，对目标发动一次兵刃攻击（伤害率185%）
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     t.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     t.tacticsParams.CurrentSufferGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 1.85,
		TacticId:          t.Id(),
		TacticName:        t.Name(),
	})
	//计穷（无法发动主动战法）1回合
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_NoStrategy,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TakeAdvantageOfQuicklyTactic) IsTriggerPrepare() bool {
	return false
}
