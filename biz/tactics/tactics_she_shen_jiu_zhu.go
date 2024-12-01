package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 舍身救主
// 战斗中，自身受到伤害降低90%，每次受到伤害后，该效果降低3%，该效果降低5次后，
// 自身受到伤害时有35%概率（受统率影响）视为2次（可额外触发反击、急救等效果）
type SheShenJiuZhuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a SheShenJiuZhuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a SheShenJiuZhuTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，自身受到伤害降低90%，
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.9,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.9,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	})

	// 每次受到伤害后，该效果降低3%
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if effectParams, ok := util.BuffEffectGet(triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce); ok {
			for _, param := range effectParams {
				if param.FromTactic == a.Id() {
					param.EffectRate -= 0.03
				}
			}
		}
		if effectParams, ok := util.BuffEffectGet(triggerGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce); ok {
			for _, param := range effectParams {
				if param.FromTactic == a.Id() {
					param.EffectRate -= 0.03
				}
			}
		}

		return triggerResp
	})
}

func (a SheShenJiuZhuTactic) Id() consts.TacticId {
	return consts.SheShenJiuZhu
}

func (a SheShenJiuZhuTactic) Name() string {
	return "舍身救主"
}

func (a SheShenJiuZhuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a SheShenJiuZhuTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a SheShenJiuZhuTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a SheShenJiuZhuTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a SheShenJiuZhuTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a SheShenJiuZhuTactic) Execute() {
}

func (a SheShenJiuZhuTactic) IsTriggerPrepare() bool {
	return false
}

func (a SheShenJiuZhuTactic) SetTriggerPrepare(triggerPrepare bool) {
}
