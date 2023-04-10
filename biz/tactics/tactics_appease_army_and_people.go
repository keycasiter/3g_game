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

// 战法名称：抚辑军民
// 战法描述：战斗前3回合，使我军群体(2人)造成的伤害降低24%，
// 受到的伤害降低24%（受统率影响），
// 战斗第4回合时，恢复其兵力（治疗率126%，受智力影响）
type AppeaseArmyAndPeopleTactic struct {
	tacticsParams *model.TacticsParams
}

func (a AppeaseArmyAndPeopleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (a AppeaseArmyAndPeopleTactic) TriggerRate() float64 {
	return 1.00
}

func (a AppeaseArmyAndPeopleTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	//找到我军队伍
	pairGeneralArr := util.GetPairGeneralsTwoArr(a.tacticsParams)
	//使我军群体(2人)造成的伤害降低24%
	launchDamageDeduceRate := 0.24
	for _, general := range pairGeneralArr {
		//造成谋略伤害降低
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] += launchDamageDeduceRate
		hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了%.2f%%", general.BaseInfo.Name,
			launchDamageDeduceRate*100)
		//造成兵刃伤害降低
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] += launchDamageDeduceRate
		hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害降低了%.2f%%", general.BaseInfo.Name,
			launchDamageDeduceRate*100)
		//注册效果
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			if params.CurrentRound == consts.Battle_Round_Fourth {
				triggerGeneral := params.CurrentGeneral
				//造成谋略伤害降低消失
				triggerGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] -= launchDamageDeduceRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					triggerGeneral.BaseInfo.Name,
					consts.DebuffEffectType_LaunchStrategyDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提升了%.2f%%", triggerGeneral.BaseInfo.Name,
					launchDamageDeduceRate*100)
				//造成兵刃伤害降低消失
				triggerGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] -= launchDamageDeduceRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					triggerGeneral.BaseInfo.Name,
					consts.DebuffEffectType_LaunchWeaponDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害提升了%.2f%%", triggerGeneral.BaseInfo.Name,
					launchDamageDeduceRate*100)
			}
			return &vo.TacticsTriggerResult{}
		})
	}

	//受到的伤害降低24%
	// TODO（受统率影响）
	sufferDamageDeduceRate := 0.24 + (currentGeneral.BaseInfo.AbilityAttr.CommandBase / 100 / 100)
	for _, general := range pairGeneralArr {
		//受到谋略伤害降低
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += sufferDamageDeduceRate
		hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了%.2f%%", general.BaseInfo.Name,
			sufferDamageDeduceRate*100)
		//受到兵刃伤害降低
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += sufferDamageDeduceRate
		hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了%.2f%%", general.BaseInfo.Name,
			sufferDamageDeduceRate*100)
		//注册效果
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			if params.CurrentRound == consts.Battle_Round_Fourth {
				triggerGeneral := params.CurrentGeneral

				//受到谋略伤害降低消失
				triggerGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] -= sufferDamageDeduceRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					triggerGeneral.BaseInfo.Name,
					consts.BuffEffectType_SufferStrategyDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]受到的谋略伤害提升了%.2f%%", triggerGeneral.BaseInfo.Name,
					sufferDamageDeduceRate*100)
				//受到兵刃伤害降低消失
				triggerGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] -= sufferDamageDeduceRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					triggerGeneral.BaseInfo.Name,
					consts.BuffEffectType_SufferWeaponDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害提升了%.2f%%", triggerGeneral.BaseInfo.Name,
					sufferDamageDeduceRate*100)
			}
			return &vo.TacticsTriggerResult{}
		})
	}

	//战斗第4回合时，恢复其兵力
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		//第四回合
		if params.CurrentRound == consts.Battle_Round_Fourth {
			triggerGeneral := params.CurrentGeneral
			//恢复兵力
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				a.Name(),
				consts.BuffEffectType_AppeaseArmyAndPeople_Prepare,
			)

			pairArr := util.GetPairGeneralsTwoArr(a.tacticsParams)
			for _, general := range pairArr {
				//恢复兵力
				//TODO（治疗率126%，受智力影响）
				resumeNum := cast.ToInt64(general.BaseInfo.AbilityAttr.IntelligenceBase * 1.26)
				resume, origin, final := util.ResumeSoldierNum(general, resumeNum)
				hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
					general.BaseInfo.Name,
					resume,
					origin,
					final,
				)
			}
		}
		return &vo.TacticsTriggerResult{}
	},
	)
	hlog.CtxInfof(ctx, "[%s]的「%s[预备]」效果已施加", a.tacticsParams.CurrentGeneral.BaseInfo.Name,
		a.Name(),
	)
}

func (a AppeaseArmyAndPeopleTactic) Execute() {
	return
}

func (a AppeaseArmyAndPeopleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	return a
}

func (a AppeaseArmyAndPeopleTactic) Id() consts.TacticId {
	return consts.AppeaseArmyAndPeople
}

func (a AppeaseArmyAndPeopleTactic) Name() string {
	return "抚辑军民"
}

func (a AppeaseArmyAndPeopleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a AppeaseArmyAndPeopleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
