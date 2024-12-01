package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 梳妆掷戟
// 使敌军主将受到兵刃伤害提升15%，武力最高的副将智力降低100点（自身为女性时，受魅力影响），持续1回合，
// 战斗第2回合起，发动时额外使二者相互发动1次兵刃攻击（伤害率155%）
type DressingUpAndThrowingHalberdsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DressingUpAndThrowingHalberdsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.5
	return d
}

func (d DressingUpAndThrowingHalberdsTactic) Prepare() {

}

func (d DressingUpAndThrowingHalberdsTactic) Id() consts.TacticId {
	return consts.DressingUpAndThrowingHalberds
}

func (d DressingUpAndThrowingHalberdsTactic) Name() string {
	return "梳妆掷戟"
}

func (d DressingUpAndThrowingHalberdsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (d DressingUpAndThrowingHalberdsTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DressingUpAndThrowingHalberdsTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DressingUpAndThrowingHalberdsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DressingUpAndThrowingHalberdsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DressingUpAndThrowingHalberdsTactic) Execute() {
	// 使敌军主将受到兵刃伤害提升15%，武力最高的副将智力降低100点（自身为女性时，受魅力影响），持续1回合

	//（自身为女性时，受魅力影响）
	rate := 0.15
	if d.tacticsParams.CurrentGeneral.BaseInfo.Gender == consts.Gender_Female {
		rate += d.tacticsParams.CurrentGeneral.BaseInfo.AbilityAttr.CharmBase / 100 / 100
	}
	//使敌军主将受到兵刃伤害提升15%
	enemyGeneral := util.GetEnemyMasterGeneral(d.tacticsParams.CurrentGeneral, d.tacticsParams)
	util.DebuffEffectWrapSet(d.tacticsParams.Ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRound: 1,
		EffectRate:  rate,
		FromTactic:  d.Id(),
	})
	//武力最高的副将智力降低100点
	viceGeneral := util.GetEnemyOneViceGeneralHighestAttr(d.tacticsParams.CurrentGeneral, d.tacticsParams, consts.AbilityAttr_Force)
	if viceGeneral == nil {
		return
	}
	util.DebuffEffectWrapSet(d.tacticsParams.Ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
		EffectRound: 1,
		EffectValue: 100,
		FromTactic:  d.Id(),
	})

	// 战斗第2回合起，发动时额外使二者相互发动1次兵刃攻击（伤害率155%）
	if d.tacticsParams.CurrentRound >= consts.Battle_Round_Second {
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     d.tacticsParams,
			AttackGeneral:     enemyGeneral,
			SufferGeneral:     viceGeneral,
			DamageImproveRate: 1.55,
			DamageType:        consts.DamageType_Weapon,
			TacticId:          d.Id(),
			TacticName:        d.Name(),
		})

		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     d.tacticsParams,
			AttackGeneral:     viceGeneral,
			SufferGeneral:     enemyGeneral,
			DamageImproveRate: 1.55,
			DamageType:        consts.DamageType_Weapon,
			TacticId:          d.Id(),
			TacticName:        d.Name(),
		})
	}
}

func (d DressingUpAndThrowingHalberdsTactic) IsTriggerPrepare() bool {
	return false
}

func (a DressingUpAndThrowingHalberdsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
