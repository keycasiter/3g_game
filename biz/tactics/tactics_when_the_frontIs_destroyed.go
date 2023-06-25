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

// 当锋摧决
// 普通攻击之后，对攻击目标再次造成一次谋略攻击（伤害率182%，受智力影响），并伪报（禁用被动战法及指挥战法）1回合
type WhenTheFrontIsDestroyedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WhenTheFrontIsDestroyedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 0.35
	return w
}

func (w WhenTheFrontIsDestroyedTactic) Prepare() {
}

func (w WhenTheFrontIsDestroyedTactic) Id() consts.TacticId {
	return consts.WhenTheFrontIsDestroyed
}

func (w WhenTheFrontIsDestroyedTactic) Name() string {
	return "当锋摧决"
}

func (w WhenTheFrontIsDestroyedTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (w WhenTheFrontIsDestroyedTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WhenTheFrontIsDestroyedTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WhenTheFrontIsDestroyedTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (w WhenTheFrontIsDestroyedTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WhenTheFrontIsDestroyedTactic) Execute() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral
	sufferGeneral := w.tacticsParams.CurrentSufferGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)
	//普通攻击之后，对攻击目标再次造成一次谋略攻击（伤害率182%，受智力影响），并伪报（禁用被动战法及指挥战法）1回合
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.82)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: w.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: sufferGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      w.Id(),
		TacticName:    w.Name(),
	})
	//伪报
	if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_FalseReport, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     w.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_FalseReport,
				TacticId:   w.Id(),
			})

			return revokeResp
		})
	}
}

func (w WhenTheFrontIsDestroyedTactic) IsTriggerPrepare() bool {
	return false
}
