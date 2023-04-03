package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
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
	tacticsParams model.TacticsParams

	//我军群体2人索引
	generalIdxMap map[int64]bool
}

func (a AppeaseArmyAndPeopleTactic) Prepare() {
	ctx := a.tacticsParams.Ctx

	//生成我军群体2人索引
	a.generalIdxMap = util.GenerateHitIdxMap(2, 3)

	hlog.CtxInfof(ctx, "[%s]发动战法[%s]",
		a.tacticsParams.CurrentGeneral.BaseInfo.Name,
		a.Name(),
	)
	//找到我军队伍
	pairGeneralArr := util.GetPairGeneralArr(a.tacticsParams)
	//使我军群体(2人)造成的伤害降低24%
	for idx, general := range pairGeneralArr {
		//按随机索引匹配武将进行生效
		if _, ok := a.generalIdxMap[cast.ToInt64(idx)]; ok {
			//造成谋略伤害降低
			general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] += 0.24
			hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了24.00%", general.BaseInfo.Name)
			//造成兵刃伤害降低
			general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] += 0.24
			hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害降低了24.00%", general.BaseInfo.Name)
		}
	}

	//受到的伤害降低24%
	// TODO（受统率影响）
	for idx, general := range pairGeneralArr {
		//按随机索引匹配武将进行生效
		if _, ok := a.generalIdxMap[cast.ToInt64(idx)]; ok {
			//受到谋略伤害降低
			general.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += 0.24
			hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了24.00%", general.BaseInfo.Name)
			//受到兵刃伤害降低
			general.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += 0.24
			hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了24.00%", general.BaseInfo.Name)
		}
	}

	//战斗第4回合时，恢复其兵力
	//TODO（治疗率126%，受智力影响）
	for idx, general := range pairGeneralArr {
		//按随机索引匹配武将进行生效
		if _, ok := a.generalIdxMap[cast.ToInt64(idx)]; ok {
			//恢复兵力
			general.BuffEffectTriggerMap[consts.BuffEffectType_Rest][consts.Battle_Round_Fourth] =
				general.BaseInfo.AbilityAttr.IntelligenceBase * 1.26
		}
	}
	hlog.CtxInfof(ctx, "[%s]的「%s[预备]」效果已施加", a.tacticsParams.CurrentGeneral.BaseInfo.Name,
		a.Name(),
	)
}

func (a AppeaseArmyAndPeopleTactic) Execute() {
	return
}

func (a AppeaseArmyAndPeopleTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	return a
}

func (a AppeaseArmyAndPeopleTactic) Id() int64 {
	return AppeaseArmyAndPeople
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

func (a AppeaseArmyAndPeopleTactic) Trigger() {
	return
}
