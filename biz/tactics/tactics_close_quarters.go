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

// 短兵相见
// 主动 40%
// 使敌军单体降低30点统率，持续1回合，并对其造成一次兵刃攻击（伤害率210%）
type CloseQuartersTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CloseQuartersTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.4
	return c
}

func (c CloseQuartersTactic) Prepare() {

}

func (c CloseQuartersTactic) Id() consts.TacticId {
	return consts.CloseQuarters
}

func (c CloseQuartersTactic) Name() string {
	return "短兵相见"
}

func (c CloseQuartersTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c CloseQuartersTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CloseQuartersTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CloseQuartersTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c CloseQuartersTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CloseQuartersTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//使敌军单体降低30点统率，持续1回合，并对其造成一次兵刃攻击（伤害率210%）
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneral(currentGeneral, c.tacticsParams)

	//施加效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectValue: 30,
		EffectRound: 1,
		FromTactic:  c.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   c.Id(),
			})

			return revokeResp
		})
	}
	//伤害
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     c.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 2.1,
		TacticId:          c.Id(),
		TacticName:        c.Name(),
	})
}

func (c CloseQuartersTactic) IsTriggerPrepare() bool {
	return false
}
