package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 神射
// 战斗中，使自己获得连击状态，每回合可以普通攻击2次，并使普通攻击目标统率降低25，可叠加，持续2回合
// 被动，100%
type DivineEjaculationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivineEjaculationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DivineEjaculationTactic) Prepare() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	//施加连击状态
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
		FromTactic: d.Id(),
	})
	//使普通攻击目标统率降低25，可叠加，持续2回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		sufferGeneral := d.tacticsParams.CurrentSufferGeneral

		//施加效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
			EffectRound: 2,
			EffectValue: 25,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_DecrCommand,
					TacticId:   d.Id(),
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}

func (d DivineEjaculationTactic) Id() consts.TacticId {
	return consts.DivineEjaculation
}

func (d DivineEjaculationTactic) Name() string {
	return "神射"
}

func (d DivineEjaculationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DivineEjaculationTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DivineEjaculationTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DivineEjaculationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (d DivineEjaculationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DivineEjaculationTactic) Execute() {
}

func (d DivineEjaculationTactic) IsTriggerPrepare() bool {
	return false
}
