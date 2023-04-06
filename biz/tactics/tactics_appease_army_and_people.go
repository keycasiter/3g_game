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

func (a AppeaseArmyAndPeopleTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法[%s]",
		a.tacticsParams.CurrentGeneral.BaseInfo.Name,
		a.Name(),
	)
	//找到我军队伍
	pairGeneralArr := util.GetPairGeneralsTwoArr(a.tacticsParams)
	//使我军群体(2人)造成的伤害降低24%
	for _, general := range pairGeneralArr {
		//造成谋略伤害降低
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] += 0.24
		hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了24.00%%", general.BaseInfo.Name)
		//造成兵刃伤害降低
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] += 0.24
		hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害降低了24.00%%", general.BaseInfo.Name)
	}

	//受到的伤害降低24%
	// TODO（受统率影响）
	for _, general := range pairGeneralArr {
		//受到谋略伤害降低
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += 0.24
		hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了24.00%%", general.BaseInfo.Name)
		//受到兵刃伤害降低
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += 0.24
		hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了24.00%%", general.BaseInfo.Name)
	}

	//战斗第4回合时，恢复其兵力
	//注册效果
	util.TacticsTriggerWrapSet(currentGeneral,
		consts.BattleAction_Attack,
		func(params *vo.TacticsTriggerParams) {
			//第四回合
			if a.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					currentGeneral.BaseInfo.Name,
					a.Name(),
					consts.BuffEffectType_AppeaseArmyAndPeople_Prepare,
				)
				for _, general := range pairGeneralArr {
					//恢复兵力
					//TODO（治疗率126%，受智力影响）
					resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.26)
					resume, origin, final := util.ResumeSoldierNum(general, resumeNum)
					hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
						general.BaseInfo.Name,
						resume,
						origin,
						final,
					)
				}
			}
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
