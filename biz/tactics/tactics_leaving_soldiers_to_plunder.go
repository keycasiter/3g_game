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

// 纵兵劫掠
// 对敌军单体造成兵刃攻击（伤害率172%）及震慑状态（无法行动），持续1回合
type LeavingSoldiersToPlunderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LeavingSoldiersToPlunderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 0.35
	return l
}

func (l LeavingSoldiersToPlunderTactic) Prepare() {

}

func (l LeavingSoldiersToPlunderTactic) Id() consts.TacticId {
	return consts.LeavingSoldiersToPlunder
}

func (l LeavingSoldiersToPlunderTactic) Name() string {
	return "纵兵劫掠"
}

func (l LeavingSoldiersToPlunderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LeavingSoldiersToPlunderTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LeavingSoldiersToPlunderTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LeavingSoldiersToPlunderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (l LeavingSoldiersToPlunderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LeavingSoldiersToPlunderTactic) Execute() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)
	// 对敌军单体造成兵刃攻击（伤害率172%）及震慑状态（无法行动），持续1回合
	enemyGeneral := util.GetEnemyOneGeneral(currentGeneral, l.tacticsParams)
	if enemyGeneral == nil {
		return
	}
	//伤害
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     l.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 1.72,
		TacticName:        l.Name(),
	})

	//效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     l.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Awe,
				TacticId:   l.Id(),
			})

			return revokeResp
		})
	}
}

func (l LeavingSoldiersToPlunderTactic) IsTriggerPrepare() bool {
	return false
}

func (a LeavingSoldiersToPlunderTactic) SetTriggerPrepare(triggerPrepare bool) {
}
