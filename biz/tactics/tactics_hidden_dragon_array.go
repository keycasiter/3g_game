package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 潜龙阵
// 我军三名武将阵营均不相同时，主将提升15%武力、智力、速度、统率，造成伤害降低15%
// 副将造成伤害提高15%,受到伤害降低15%，且可触发战法的主将效果
// 若我军主将的任意战法拥有主将效果，使其失去该效果且属性提升值降低为5%
type HiddenDragonArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HiddenDragonArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 1.0
	return h
}

func (h HiddenDragonArrayTactic) Prepare() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	// 我军三名武将阵营均不相同时，主将提升15%武力、智力、速度、统率，造成伤害降低15%
	pairGenerals := util.GetPairGeneralArr(h.tacticsParams)
	groupMap := make(map[consts.Group]bool)
	for _, pairGeneral := range pairGenerals {
		groupMap[pairGeneral.BaseInfo.Group] = true
	}
	if len(groupMap) != 3 {
		hlog.CtxInfof(ctx, "[%s]【%s】无法生效，武将阵营未均不相同",
			currentGeneral.BaseInfo.Name,
			h.Name(),
		)
		return
	}
	// 副将造成伤害提高15%,受到伤害降低15%，且可触发战法的主将效果
	viceGenerals := util.GetPairViceGenerals(h.tacticsParams)
	for _, viceGeneral := range viceGenerals {
		//造成伤害提高
		util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.15,
			FromTactic:     h.Id(),
			ProduceGeneral: currentGeneral,
		})
		util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.15,
			FromTactic:     h.Id(),
			ProduceGeneral: currentGeneral,
		})
		//受到伤害降低
		util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.15,
			FromTactic:     h.Id(),
			ProduceGeneral: currentGeneral,
		})
		util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.15,
			FromTactic:     h.Id(),
			ProduceGeneral: currentGeneral,
		})
		//TODO 且可触发战法的主将效果
	}
	//TODO 若我军主将的任意战法拥有主将效果，使其失去该效果且属性提升值降低为5%

}

func (h HiddenDragonArrayTactic) Id() consts.TacticId {
	return consts.HiddenDragonArray
}

func (h HiddenDragonArrayTactic) Name() string {
	return "潜龙阵"
}

func (h HiddenDragonArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (h HiddenDragonArrayTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HiddenDragonArrayTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HiddenDragonArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (h HiddenDragonArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HiddenDragonArrayTactic) Execute() {
}

func (h HiddenDragonArrayTactic) IsTriggerPrepare() bool {
	return false
}
