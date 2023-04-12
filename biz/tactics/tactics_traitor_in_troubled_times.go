package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：乱世奸雄
// 战法描述：战斗中，使友军群体(2人)造成伤害提高16%（受智力影响），
// 自己受到伤害降低18%（受智力影响），如果自己为主将，副将造成伤害时，会为主将恢复其伤害量10%的兵力
type TraitorInTroubledTimesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TraitorInTroubledTimesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TraitorInTroubledTimesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TraitorInTroubledTimesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TraitorInTroubledTimesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TraitorInTroubledTimesTactic) Prepare() {
	currentGeneral := t.tacticsParams.CurrentGeneral
	ctx := t.tacticsParams.Ctx
	//战斗中，使友军群体(2人)造成伤害提高16%（受智力影响）
	//找到队友
	pairGenerals := util.GetPairGeneralsTwoArr(t.tacticsParams)
	for _, general := range pairGenerals {
		//造成伤害提高16% TODO （受智力影响）
		rate := 0.16
		rate += general.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
		general.BuffEffectHolderMap[consts.BuffEffectType_LaunchStrategyDamageImprove] += rate
		general.BuffEffectHolderMap[consts.BuffEffectType_LaunchWeaponDamageImprove] += rate
		hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害提高了%.2f%%",
			general.BaseInfo.Name,
			rate*100,
		)
		hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提高了%.2f%%",
			general.BaseInfo.Name,
			rate*100,
		)
	}

	//自己受到伤害降低18%  TODO（受智力影响）
	rate := 0.18
	rate += currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += rate
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += rate
	hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了%.2f%%",
		currentGeneral.BaseInfo.Name,
		rate*100,
	)
	hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了%.2f%%",
		currentGeneral.BaseInfo.Name,
		rate*100,
	)

	//TODO 如果自己为主将，副将造成伤害时，会为主将恢复其伤害量10%的兵力
}

func (t TraitorInTroubledTimesTactic) Name() string {
	return "乱世奸雄"
}

func (t TraitorInTroubledTimesTactic) Execute() {
	return
}

func (t TraitorInTroubledTimesTactic) Trigger() {
	return
}

func (t TraitorInTroubledTimesTactic) Id() consts.TacticId {
	return consts.TraitorInTroubledTimes
}

func (t TraitorInTroubledTimesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TraitorInTroubledTimesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
