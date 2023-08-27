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

// 后发制人
// 受到普通攻击时对攻击者进行一次反击(伤害率52%)
// 被动 100%
type GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Prepare() {
	currentGeneral := g.tacticsParams.CurrentGeneral
	ctx := g.tacticsParams.Ctx

	// 受到普通攻击时对攻击者进行一次反击(伤害率52%)
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		attackGeneral := params.AttackGeneral

		dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 0.52)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: g.tacticsParams,
			AttackGeneral: triggerGeneral,
			SufferGeneral: attackGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      g.Id(),
			TacticName:    g.Name(),
		})

		return triggerResp
	})
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Id() consts.TacticId {
	return consts.GainMasteryByStrikingOnlyAfterTheEnemyHasStruck
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Name() string {
	return "后发制人"
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Execute() {

}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) IsTriggerPrepare() bool {
	return false
}
