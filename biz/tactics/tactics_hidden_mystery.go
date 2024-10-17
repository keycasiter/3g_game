package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/spf13/cast"
)

// 暗藏玄机
// 普通攻击之后，对攻击目标在此发起一次兵刃攻击（伤害率144%），如果目标为敌军主将则额外造成一次谋略攻击（伤害率92%，受智力影响）
type HiddenMysteryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HiddenMysteryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HiddenMysteryTactic) Prepare() {
}

func (h HiddenMysteryTactic) Id() consts.TacticId {
	return consts.HiddenMystery
}

func (h HiddenMysteryTactic) Name() string {
	return "暗藏玄机"
}

func (h HiddenMysteryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HiddenMysteryTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HiddenMysteryTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HiddenMysteryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (h HiddenMysteryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HiddenMysteryTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	//普通攻击之后，对攻击目标在此发起一次兵刃攻击（伤害率144%）
	weaponDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.44)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: h.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: h.tacticsParams.CurrentSufferGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        weaponDmg,
		TacticName:    h.Name(),
	})
	//如果目标为敌军主将则额外造成一次谋略攻击（伤害率92%，受智力影响）
	if h.tacticsParams.CurrentSufferGeneral.IsMaster {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.92)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: h.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: h.tacticsParams.CurrentSufferGeneral,
			DamageType:    consts.DamageType_Strategy,
			Damage:        dmg,
			TacticName:    h.Name(),
		})
	}
}

func (h HiddenMysteryTactic) IsTriggerPrepare() bool {
	return false
}
