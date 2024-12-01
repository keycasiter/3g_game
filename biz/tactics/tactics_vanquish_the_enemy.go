package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 克敌制胜
// 普通攻击之后，对攻击目标再次造成一次谋略伤害(伤害率180%，受智力影响)；
// 若目标处于溃逃或中毒状态，则有85%概率使目标进入虚弱（无法造成伤害）状态，持续1回合
type VanquishTheEnemyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (v VanquishTheEnemyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	v.tacticsParams = tacticsParams
	v.triggerRate = 0.4
	return v
}

func (v VanquishTheEnemyTactic) Prepare() {
}

func (v VanquishTheEnemyTactic) Id() consts.TacticId {
	return consts.VanquishTheEnemy
}

func (v VanquishTheEnemyTactic) Name() string {
	return "克敌制胜"
}

func (v VanquishTheEnemyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (v VanquishTheEnemyTactic) GetTriggerRate() float64 {
	return v.triggerRate
}

func (v VanquishTheEnemyTactic) SetTriggerRate(rate float64) {
	v.triggerRate = rate
}

func (v VanquishTheEnemyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (v VanquishTheEnemyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (v VanquishTheEnemyTactic) Execute() {
	ctx := v.tacticsParams.Ctx
	currentGeneral := v.tacticsParams.CurrentGeneral
	sufferGeneral := v.tacticsParams.CurrentSufferGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		v.Name(),
	)
	// 普通攻击之后，对攻击目标再次造成一次谋略伤害(伤害率180%，受智力影响)；
	dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.8
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     v.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     sufferGeneral,
		DamageType:        consts.DamageType_Strategy,
		DamageImproveRate: dmgRate,
		TacticId:          v.Id(),
		TacticName:        v.Name(),
	})
	// 若目标处于溃逃或中毒状态，则有85%概率使目标进入虚弱（无法造成伤害）状态，持续1回合
	if util.DeBuffEffectContains(sufferGeneral, consts.DebuffEffectType_Methysis) ||
		util.DeBuffEffectContains(sufferGeneral, consts.DebuffEffectType_Escape) {
		if util.GenerateRate(0.85) {
			if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
				EffectRound:    1,
				FromTactic:     v.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_PoorHealth,
						TacticId:   v.Id(),
					})

					return revokeResp
				})
			}
		}
	}
}

func (v VanquishTheEnemyTactic) IsTriggerPrepare() bool {
	return false
}

func (a VanquishTheEnemyTactic) SetTriggerPrepare(triggerPrepare bool) {
}
