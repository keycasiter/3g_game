package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 乘胜长驱
// 战斗中，每回合使自己造成伤害提高11%，可叠加，直到战斗结束
// 被动，100%
type RidingOnTheVictoryDriveTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RidingOnTheVictoryDriveTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RidingOnTheVictoryDriveTactic) Prepare() {
}

func (r RidingOnTheVictoryDriveTactic) Id() consts.TacticId {
	return consts.RidingOnTheVictoryDrive
}

func (r RidingOnTheVictoryDriveTactic) Name() string {
	return "乘胜长驱"
}

func (r RidingOnTheVictoryDriveTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RidingOnTheVictoryDriveTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RidingOnTheVictoryDriveTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RidingOnTheVictoryDriveTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (r RidingOnTheVictoryDriveTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RidingOnTheVictoryDriveTactic) Execute() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)

	//战斗中，每回合使自己造成伤害提高11%，可叠加，直到战斗结束
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//兵刃伤害
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.11,
			FromTactic:     r.Id(),
			ProduceGeneral: triggerGeneral,
		})
		//谋略伤害
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.11,
			FromTactic:     r.Id(),
			ProduceGeneral: triggerGeneral,
		})

		return triggerResp
	})
}

func (r RidingOnTheVictoryDriveTactic) IsTriggerPrepare() bool {
	return false
}
