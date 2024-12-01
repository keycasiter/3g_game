package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 奋突
// 普通攻击之后，使自己造成兵刃伤害提高12%，最多叠加3次，并且有35%概率使目标缴械（无法进行普通攻击），持续1回合
type RiseUpBravelyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RiseUpBravelyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RiseUpBravelyTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral
	sufferGeneral := r.tacticsParams.CurrentSufferGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		// 普通攻击之后，使自己造成兵刃伤害提高12%，最多叠加3次，并且有35%概率使目标缴械（无法进行普通攻击），持续1回合
		if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.12,
			EffectRound:    1,
			EffectTimes:    1,
			MaxEffectTimes: 3,
			FromTactic:     r.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
					TacticId:   r.Id(),
				})

				return revokeResp
			})
		}
		//缴械
		if util.GenerateRate(0.35) {
			if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
				EffectRound:    1,
				FromTactic:     r.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_CancelWeapon,
						TacticId:   r.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (r RiseUpBravelyTactic) Id() consts.TacticId {
	return consts.RiseUpBravely
}

func (r RiseUpBravelyTactic) Name() string {
	return "奋突"
}

func (r RiseUpBravelyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RiseUpBravelyTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RiseUpBravelyTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RiseUpBravelyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (r RiseUpBravelyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RiseUpBravelyTactic) Execute() {

}

func (r RiseUpBravelyTactic) IsTriggerPrepare() bool {
	return false
}

func (a RiseUpBravelyTactic) SetTriggerPrepare(triggerPrepare bool) {
}
