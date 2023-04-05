package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：锋矢阵
// 战法描述：战斗中，使我军主将造成的伤害提升30%，受到的伤害提升20%，我军副将造成的伤害降低15%，受到的伤害降低25%
type FrontalVectorArrayTactic struct {
	tacticsParams model.TacticsParams
}

func (f FrontalVectorArrayTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	return f
}

func (f FrontalVectorArrayTactic) Prepare() {
	//使我军主将造成的伤害提升30%，受到的伤害提升20%
	masterGeneral := util.GetPairMasterGeneral(f.tacticsParams)
	masterGeneral.BuffEffectHolderMap[consts.BuffEffectType_LaunchStrategyDamageImprove] += 0.3
	masterGeneral.BuffEffectHolderMap[consts.BuffEffectType_LaunchWeaponDamageImprove] += 0.3
	masterGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_SufferWeaponDamageImprove] += 0.2
	masterGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_SufferStrategyDamageImprove] += 0.2
	//我军副将造成的伤害降低15%，受到的伤害降低25%
	viceGenerals := util.GetPairViceGenerals(f.tacticsParams)
	for _, general := range viceGenerals {
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] += 0.15
		general.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] += 0.15
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += 0.25
		general.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += 0.25
	}
}

func (f FrontalVectorArrayTactic) Id() consts.TacticId {
	return consts.FrontalVectorArray
}

func (f FrontalVectorArrayTactic) Name() string {
	return "锋矢阵"
}

func (f FrontalVectorArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (f FrontalVectorArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FrontalVectorArrayTactic) Execute() {
	return
}

func (f FrontalVectorArrayTactic) Trigger() {
	return
}
