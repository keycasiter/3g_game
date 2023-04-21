package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//才器过人
//使我军群体（2人）造成的谋略伤害提高27%，持续2回合
type OutstandingTalentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OutstandingTalentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	o.tacticsParams = tacticsParams
	o.triggerRate = 0.4
	return o
}

func (o OutstandingTalentTactic) Prepare() {
}

func (o OutstandingTalentTactic) Id() consts.TacticId {
	return consts.OutstandingTalent
}

func (o OutstandingTalentTactic) Name() string {
	return "才器过人"
}

func (o OutstandingTalentTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (o OutstandingTalentTactic) GetTriggerRate() float64 {
	return o.triggerRate
}

func (o OutstandingTalentTactic) SetTriggerRate(rate float64) {
	o.triggerRate = rate
}

func (o OutstandingTalentTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (o OutstandingTalentTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (o OutstandingTalentTactic) Execute() {
	ctx := o.tacticsParams.Ctx
	currentGeneral := o.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		o.Name(),
	)

	//使我军群体（2人）造成的谋略伤害提高27%，持续2回合
	pairGeneral := util.GetPairGeneralsTwoArr(o.tacticsParams)
	for _, general := range pairGeneral {
		//设置效果
		if util.TacticsBuffEffectCountWrapIncr(ctx, general, consts.BuffEffectType_OutstandingTalent_Prepare, 2, 2, true) {
			rate := 0.27
			general.BuffEffectHolderMap[consts.BuffEffectType_OutstandingTalent_Prepare] += rate
			hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提高了%.2f%%",
				general.BaseInfo.Name,
				rate*100,
			)

			//注册消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.BuffEffectContains(revokeGeneral, consts.BuffEffectType_OutstandingTalent_Prepare) &&
					0 == util.TacticsBuffCountGet(revokeGeneral, consts.BuffEffectType_OutstandingTalent_Prepare) {
					revokeGeneral.BuffEffectHolderMap[consts.BuffEffectType_OutstandingTalent_Prepare] -= rate
					hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了%.2f%%",
						revokeGeneral.BaseInfo.Name,
						rate*100,
					)
				}

				util.TacticsBuffEffectCountWrapDecr(ctx, revokeGeneral, consts.BuffEffectType_OutstandingTalent_Prepare, 1)

				return revokeResp
			})
		}
	}
}

func (o OutstandingTalentTactic) IsTriggerPrepare() bool {
	return false
}
