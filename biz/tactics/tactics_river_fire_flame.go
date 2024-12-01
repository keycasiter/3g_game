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

// 江天长焰
// 战斗中，每回合使敌军全体受到谋略伤害提高4%（受智力影响），可叠加，
// 并有几率（32%，每有一名敌军处于灼烧或水攻状态时，提高4%）对敌军单体造成谋略伤害（伤害率146%，受智力影响）
// 自身为主将时，基础概率40%且该次攻击目标将锁定敌军兵力最低单体
// 指挥，100%
type RiverFireFlameTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RiverFireFlameTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RiverFireFlameTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	// 战斗中，每回合使敌军全体受到谋略伤害提高4%（受智力影响），可叠加，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, r.tacticsParams)
		for _, enemyGeneral := range enemyGenerals {
			util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.04 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
				FromTactic:     r.Id(),
				ProduceGeneral: triggerGeneral,
			})
		}
		// 并有几率（32%，每有一名敌军处于灼烧或水攻状态时，提高4%）对敌军单体造成谋略伤害（伤害率146%，受智力影响）
		triggerRate := 0.32
		//找到敌军单体
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, r.tacticsParams)
		// 自身为主将时，基础概率40%且该次攻击目标将锁定敌军兵力最低单体
		if currentGeneral.IsMaster {
			triggerRate = 0.4
			enemyGeneral = util.GetLowestSoliderNumGeneral(enemyGenerals)
		}
		allEnemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, r.tacticsParams)
		for _, enemyGeneral := range allEnemyGenerals {
			if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Firing) ||
				util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_WaterAttack) {
				triggerRate += 0.04
			}
		}
		dmgRate := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.46
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     r.tacticsParams,
			AttackGeneral:     triggerGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          r.Id(),
			TacticName:        r.Name(),
		})

		return triggerResp
	})
}

func (r RiverFireFlameTactic) Id() consts.TacticId {
	return consts.RiverFireFlame
}

func (r RiverFireFlameTactic) Name() string {
	return "江天长焰"
}

func (r RiverFireFlameTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (r RiverFireFlameTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RiverFireFlameTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RiverFireFlameTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (r RiverFireFlameTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RiverFireFlameTactic) Execute() {

}

func (r RiverFireFlameTactic) IsTriggerPrepare() bool {
	return false
}

func (a RiverFireFlameTactic) SetTriggerPrepare(triggerPrepare bool) {
}
