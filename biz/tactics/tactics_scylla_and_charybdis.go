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

// 腹背受敌
// 自己及敌军智力最高的武将同时对敌军单体造成谋略攻击（伤害率118%，受智力影响）
// 主动，45%
type ScyllaAndCharybdisTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ScyllaAndCharybdisTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.45
	return s
}

func (s ScyllaAndCharybdisTactic) Prepare() {
}

func (s ScyllaAndCharybdisTactic) Id() consts.TacticId {
	return consts.ScyllaAndCharybdis
}

func (s ScyllaAndCharybdisTactic) Name() string {
	return "腹背受敌"
}

func (s ScyllaAndCharybdisTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s ScyllaAndCharybdisTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ScyllaAndCharybdisTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ScyllaAndCharybdisTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s ScyllaAndCharybdisTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ScyllaAndCharybdisTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 自己及敌军智力最高的武将同时对敌军单体造成谋略攻击（伤害率118%，受智力影响）
	effectGenerals := make([]*vo.BattleGeneral, 0)
	effectGenerals = append(effectGenerals, currentGeneral)
	effectGenerals = append(effectGenerals, util.GetEnemyGeneralWhoIsHighestIntelligence(currentGeneral, s.tacticsParams))
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, s.tacticsParams)
	for _, general := range effectGenerals {
		dmgRate := general.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.18
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     s.tacticsParams,
			AttackGeneral:     general,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          s.Id(),
			TacticName:        s.Name(),
		})
	}
}

func (s ScyllaAndCharybdisTactic) IsTriggerPrepare() bool {
	return false
}

func (a ScyllaAndCharybdisTactic) SetTriggerPrepare(triggerPrepare bool) {
}
