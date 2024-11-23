package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 灵机一动
// 使我军群体2-3人智力、统率、魅力提升15(受智力影响)，可叠加，持续3回合
type AFlushOfWitTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AFlushOfWitTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.55
	return a
}

func (a AFlushOfWitTactic) Prepare() {
}

func (a AFlushOfWitTactic) Id() consts.TacticId {
	return consts.AFlashOfWit
}

func (a AFlushOfWitTactic) Name() string {
	return "灵机一动"
}

func (a AFlushOfWitTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AFlushOfWitTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AFlushOfWitTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AFlushOfWitTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a AFlushOfWitTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Apparatus,
	}
}

func (a AFlushOfWitTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//使我军群体2-3人智力、统率、魅力提升15(受智力影响)，可叠加，持续3回合
	pairGenerals := util.GetPairGeneralsTwoOrThreeMap(currentGeneral, a.tacticsParams)
	for _, general := range pairGenerals {
		//智力
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
			EffectValue:    15,
			EffectRound:    3,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrIntelligence,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
		//统率
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectValue:    15,
			EffectRound:    3,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrCommand,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
		//魅力
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCharm, &vo.EffectHolderParams{
			EffectValue:    15,
			EffectRound:    3,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrCharm,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (a AFlushOfWitTactic) IsTriggerPrepare() bool {
	return false
}
