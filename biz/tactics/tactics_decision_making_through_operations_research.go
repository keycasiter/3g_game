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

// 运筹决算
// 准备1回合，对敌军全体发动一次谋略攻击（伤害率176%，受智力影响）
type DecisionMakingThroughOperationsResearchTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (d DecisionMakingThroughOperationsResearchTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.45
	return d
}

func (d DecisionMakingThroughOperationsResearchTactic) Prepare() {
}

func (d DecisionMakingThroughOperationsResearchTactic) Id() consts.TacticId {
	return consts.DecisionMakingThroughOperationsResearch
}

func (d DecisionMakingThroughOperationsResearchTactic) Name() string {
	return "运筹决算"
}

func (d DecisionMakingThroughOperationsResearchTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DecisionMakingThroughOperationsResearchTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DecisionMakingThroughOperationsResearchTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DecisionMakingThroughOperationsResearchTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DecisionMakingThroughOperationsResearchTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DecisionMakingThroughOperationsResearchTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral
	currentRound := d.tacticsParams.CurrentRound

	//准备1回合，对敌军全体发动一次谋略攻击（伤害率176%，受智力影响）
	d.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			d.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if d.isTriggered {
				return triggerResp
			} else {
				d.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				d.Name(),
			)
			//对敌军全体发动一次谋略攻击（伤害率176%，受智力影响）
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, d.tacticsParams)
			dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.76

			for _, sufferGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     d.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     sufferGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticName:        d.Name(),
					TacticId:          d.Id(),
				})
			}
		}

		return triggerResp
	})
}

func (d DecisionMakingThroughOperationsResearchTactic) IsTriggerPrepare() bool {
	return d.isTriggerPrepare
}

func (a DecisionMakingThroughOperationsResearchTactic) SetTriggerPrepare(triggerPrepare bool) {
}
