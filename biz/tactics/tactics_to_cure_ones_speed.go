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

// 将行其疾
// 普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率158%）；
// 若命中敌军主将，则使其进入计穷状态，持续2回合
// 突击 60%
type ToCureOnesSpeedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToCureOnesSpeedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.6
	return t
}

func (t ToCureOnesSpeedTactic) Prepare() {

}

func (t ToCureOnesSpeedTactic) Id() consts.TacticId {
	return consts.ToCureOnesSpeed
}

func (t ToCureOnesSpeedTactic) Name() string {
	return "将行其疾"
}

func (t ToCureOnesSpeedTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t ToCureOnesSpeedTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToCureOnesSpeedTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToCureOnesSpeedTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t ToCureOnesSpeedTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToCureOnesSpeedTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率158%）；
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     t.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 1.58,
		TacticId:          t.Id(),
		TacticName:        t.Name(),
	})
	// 若命中敌军主将，则使其进入计穷状态，持续2回合
	if enemyGeneral.IsMaster {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_NoStrategy,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (t ToCureOnesSpeedTactic) IsTriggerPrepare() bool {
	return false
}
