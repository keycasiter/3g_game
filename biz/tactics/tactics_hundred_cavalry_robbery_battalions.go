package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 百骑劫营
// 普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率162%）同时有50%几率对敌军主将额外发动一次兵刃攻击（伤害率120%）
type HundredCavalryRobberyBattalionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HundredCavalryRobberyBattalionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HundredCavalryRobberyBattalionsTactic) Prepare() {

}

func (h HundredCavalryRobberyBattalionsTactic) Id() consts.TacticId {
	return consts.HundredCavalryRobberyBattalions
}

func (h HundredCavalryRobberyBattalionsTactic) Name() string {
	return "百骑劫营"
}

func (h HundredCavalryRobberyBattalionsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HundredCavalryRobberyBattalionsTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HundredCavalryRobberyBattalionsTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HundredCavalryRobberyBattalionsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (h HundredCavalryRobberyBattalionsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (h HundredCavalryRobberyBattalionsTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)
	//普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率162%）同时有50%几率对敌军主将额外发动一次兵刃攻击（伤害率120%）
	//敌军单体伤害
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, h.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.62)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: h.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticId:      h.Id(),
		TacticName:    h.Name(),
	})
	//额外攻击
	if util.GenerateRate(0.5) {
		//敌军主将
		enemyMasterGeneral := util.GetEnemyMasterGeneral(h.tacticsParams)
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.2)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: h.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyMasterGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      h.Id(),
			TacticName:    h.Name(),
		})
	}
}

func (h HundredCavalryRobberyBattalionsTactic) IsTriggerPrepare() bool {
	return false
}
