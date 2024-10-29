package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 屠几上肉
// 对敌军单体造成一次兵刃攻击（伤害率150%），及谋略攻击（伤害率150%，受智力影响）
type SlaughterMeatOnTableTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SlaughterMeatOnTableTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.35
	return s
}

func (s SlaughterMeatOnTableTactic) Prepare() {
}

func (s SlaughterMeatOnTableTactic) Id() consts.TacticId {
	return consts.SlaughterMeatOnTable
}

func (s SlaughterMeatOnTableTactic) Name() string {
	return "屠几上肉"
}

func (s SlaughterMeatOnTableTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SlaughterMeatOnTableTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SlaughterMeatOnTableTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SlaughterMeatOnTableTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SlaughterMeatOnTableTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SlaughterMeatOnTableTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 对敌军单体造成一次兵刃攻击（伤害率150%），及谋略攻击（伤害率150%，受智力影响）
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, s.tacticsParams)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     s.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 1.5,
		TacticId:          s.Id(),
		TacticName:        s.Name(),
	})
	strategyDmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.5
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     s.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Strategy,
		DamageImproveRate: strategyDmgRate,
		TacticId:          s.Id(),
		TacticName:        s.Name(),
	})
}

func (s SlaughterMeatOnTableTactic) IsTriggerPrepare() bool {
	return false
}
