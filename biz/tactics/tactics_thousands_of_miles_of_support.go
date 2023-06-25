package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 千里驰援
// 援护全体友军，同时提高自身40统率，持续1回合
// 主动，40%
type ThousandsOfMilesOfSupportTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThousandsOfMilesOfSupportTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.4
	return t
}

func (t ThousandsOfMilesOfSupportTactic) Prepare() {

}

func (t ThousandsOfMilesOfSupportTactic) Id() consts.TacticId {
	return consts.ThousandsOfMilesOfSupport
}

func (t ThousandsOfMilesOfSupportTactic) Name() string {
	return "千里驰援"
}

func (t ThousandsOfMilesOfSupportTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThousandsOfMilesOfSupportTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThousandsOfMilesOfSupportTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThousandsOfMilesOfSupportTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ThousandsOfMilesOfSupportTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThousandsOfMilesOfSupportTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 援护全体友军，同时提高自身40统率，持续1回合
	pairGenerals := util.GetPairGeneralsNotSelf(t.tacticsParams, currentGeneral)
	for _, pairGeneral := range pairGenerals {
		//援护效果
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Intervene, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Intervene,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
		//统率提升
		if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     t.Id(),
			EffectValue:    40,
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrCommand,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (t ThousandsOfMilesOfSupportTactic) IsTriggerPrepare() bool {
	return false
}
