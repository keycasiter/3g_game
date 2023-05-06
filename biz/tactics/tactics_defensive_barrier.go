package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 御敌屏障
// 战斗前4回合，使我军群体（2人）受到的伤害降低25%
type DefensiveBarrierTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DefensiveBarrierTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DefensiveBarrierTactic) Prepare() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)

	//找到随机友军2人
	pairGenerals := util.GetPairGeneralsTwoArr(d.tacticsParams)
	for _, general := range pairGenerals {
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.25,
			FromTactic: d.Id(),
		}).IsSuccess {
			hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了25%%",
				general.BaseInfo.Name,
			)
		}

		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.25,
			FromTactic: d.Id(),
		}).IsSuccess {
			hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了25%%",
				general.BaseInfo.Name,
			)
		}

		//注册消失效果
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeRound := params.CurrentRound
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			if revokeRound == consts.Battle_Round_Fifth {
				if util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, d.Id()) {
					hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害提高了25%%",
						general.BaseInfo.Name,
					)
				}
				if util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, d.Id()) {
					hlog.CtxInfof(ctx, "[%s]受到的谋略伤害提高了25%%",
						general.BaseInfo.Name,
					)
				}
			}

			return revokeResp
		})
	}
}

func (d DefensiveBarrierTactic) Id() consts.TacticId {
	return consts.DefensiveBarrier
}

func (d DefensiveBarrierTactic) Name() string {
	return "御敌屏障"
}

func (d DefensiveBarrierTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DefensiveBarrierTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DefensiveBarrierTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DefensiveBarrierTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (d DefensiveBarrierTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DefensiveBarrierTactic) Execute() {

}

func (d DefensiveBarrierTactic) IsTriggerPrepare() bool {
	return false
}
