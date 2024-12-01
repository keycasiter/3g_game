package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 符命自立
// 战斗前2回合中任一回合，自身获得玉玺，提高100%会心几率及奇谋几率，每回合逐渐降低，直至战斗第8回合降至0，
// 自身为主将时，提高主动战法发动几率25%，准备战法为35%，同样会逐渐降低
// 被动 100%
type FumingSelfRelianceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FumingSelfRelianceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FumingSelfRelianceTactic) Prepare() {
	currentGeneral := f.tacticsParams.CurrentGeneral
	ctx := f.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	// 战斗前2回合中任一回合，自身获得玉玺，提高100%会心几率及奇谋几率，每回合逐渐降低，直至战斗第8回合降至0，
	// 自身为主将时，提高主动战法发动几率25%，准备战法为35%，同样会逐渐降低

	hitIdx := util.GenerateHitOneIdx(2)
	//触发回合
	hitRound := hitIdx + 1
	//剩余回合
	remainRound := 8 - hitRound

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		if triggerRound == consts.BattleRound(hitRound) {
			//自身获得玉玺，提高100%会心几率及奇谋几率
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
				TriggerRate:    1.0,
				FromTactic:     f.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//每回合逐渐降低
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					deduceRatePerRound := cast.ToFloat64(100 / remainRound)

					util.BuffEffectOfTacticDeduce(&util.BuffEffectOfTacticDeduceParams{
						Ctx:         ctx,
						General:     revokeGeneral,
						EffectType:  consts.BuffEffectType_EnhanceStrategy,
						TacticId:    f.Id(),
						TriggerRate: deduceRatePerRound,
					})

					return revokeResp
				})
			}

			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
				TriggerRate:    1.0,
				FromTactic:     f.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//每回合逐渐降低
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					deduceRatePerRound := cast.ToFloat64(100 / remainRound)

					util.BuffEffectOfTacticDeduce(&util.BuffEffectOfTacticDeduceParams{
						Ctx:         ctx,
						General:     revokeGeneral,
						EffectType:  consts.BuffEffectType_EnhanceWeapon,
						TacticId:    f.Id(),
						TriggerRate: deduceRatePerRound,
					})

					return revokeResp
				})
			}
			//自身为主将时，提高主动战法发动几率25%，准备战法为35%，同样会逐渐降低
			if triggerGeneral.IsMaster {
				//主动战法发动率
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
					TriggerRate:    0.25,
					FromTactic:     f.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//每回合逐渐降低
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						deduceRatePerRound := 0.25 / cast.ToFloat64(remainRound)

						util.BuffEffectOfTacticDeduce(&util.BuffEffectOfTacticDeduceParams{
							Ctx:         ctx,
							General:     revokeGeneral,
							EffectType:  consts.BuffEffectType_TacticsActiveTriggerImprove,
							TacticId:    f.Id(),
							TriggerRate: deduceRatePerRound,
						})

						return revokeResp
					})
				}
				//准备战法
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerPrepareImprove, &vo.EffectHolderParams{
					TriggerRate:    0.15,
					FromTactic:     f.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//每回合逐渐降低
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						deduceRatePerRound := 0.15 / cast.ToFloat64(remainRound)

						util.BuffEffectOfTacticDeduce(&util.BuffEffectOfTacticDeduceParams{
							Ctx:         ctx,
							General:     revokeGeneral,
							EffectType:  consts.BuffEffectType_TacticsActiveTriggerPrepareImprove,
							TacticId:    f.Id(),
							TriggerRate: deduceRatePerRound,
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (f FumingSelfRelianceTactic) Id() consts.TacticId {
	return consts.FumingSelfReliance
}

func (f FumingSelfRelianceTactic) Name() string {
	return "符命自立"
}

func (f FumingSelfRelianceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FumingSelfRelianceTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FumingSelfRelianceTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FumingSelfRelianceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (f FumingSelfRelianceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FumingSelfRelianceTactic) Execute() {

}

func (f FumingSelfRelianceTactic) IsTriggerPrepare() bool {
	return false
}

func (a FumingSelfRelianceTactic) SetTriggerPrepare(triggerPrepare bool) {
}
