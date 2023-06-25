package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 白马义从
// 我军全体战斗前2回合获得先攻，并提高10%主动战法发动率
// 若公孙瓒统领，提高发动率受速度影响
// 兵种，100%
type WhiteHorseFollowsWithLoyaltyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)
	// 我军全体战斗前2回合获得先攻，并提高10%主动战法发动率
	pairGenerals := util.GetPairGeneralArr(w.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//先攻
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_FirstAttack, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     w.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_FirstAttack,
					TacticId:   w.Id(),
				})

				return revokeResp
			})
		}
		//发动率
		triggerRate := 0.1
		//若公孙瓒统领，提高发动率受速度影响
		if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.GongSunZan {
			triggerRate += currentGeneral.BaseInfo.AbilityAttr.SpeedBase / 100 / 100
		}
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
			EffectRound:    2,
			TriggerRate:    triggerRate,
			FromTactic:     w.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
					TacticId:   w.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Id() consts.TacticId {
	return consts.WhiteHorseFollowsWithLoyalty
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Name() string {
	return "白马义从"
}

func (w WhiteHorseFollowsWithLoyaltyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WhiteHorseFollowsWithLoyaltyTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WhiteHorseFollowsWithLoyaltyTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WhiteHorseFollowsWithLoyaltyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (w WhiteHorseFollowsWithLoyaltyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Execute() {
}

func (w WhiteHorseFollowsWithLoyaltyTactic) IsTriggerPrepare() bool {
	return false
}
