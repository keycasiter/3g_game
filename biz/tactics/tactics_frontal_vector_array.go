package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：锋矢阵
// 战法描述：战斗中，使我军主将造成的伤害提升30%，受到的伤害提升20%，我军副将造成的伤害降低15%，受到的伤害降低25%
type FrontalVectorArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FrontalVectorArrayTactic) IsTriggerPrepare() bool {
	return false
}

func (f FrontalVectorArrayTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FrontalVectorArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FrontalVectorArrayTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FrontalVectorArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FrontalVectorArrayTactic) Prepare() {
	currentGeneral := f.tacticsParams.CurrentGeneral
	ctx := f.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	//使我军主将造成的伤害提升30%，受到的伤害提升20%
	masterGeneral := util.GetPairMasterGeneral(f.tacticsParams)

	util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate: 0.3,
		FromTactic: f.Id(),
	})
	util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate: 0.3,
		FromTactic: f.Id(),
	})
	util.DebuffEffectWrapSet(ctx, masterGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate: 0.2,
		FromTactic: f.Id(),
	})
	util.DebuffEffectWrapSet(ctx, masterGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate: 0.2,
		FromTactic: f.Id(),
	})
	//我军副将造成的伤害降低15%，受到的伤害降低25%
	viceGenerals := util.GetPairViceGenerals(f.tacticsParams)
	for _, general := range viceGenerals {
		util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.15,
			FromTactic: f.Id(),
		})
		util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.15,
			FromTactic: f.Id(),
		})
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.25,
			FromTactic: f.Id(),
		})
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.25,
			FromTactic: f.Id(),
		})
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
