package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 监统震军
// 战斗中，友军群体造成负面状态时，有38%概率（受智力影响）使负面状态持续施加增加1回合（混乱及伪报不生效）
// 自身普通攻击后，对负面状态最多的敌军单体造成谋略攻击（伤害率114%，受智力影响）
// 若敌军都没有负面状态则为损失兵力最多的我军单体恢复兵力（治疗率92%，受智力影响）
type SuperviseLeadAndSeizureArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SuperviseLeadAndSeizureArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SuperviseLeadAndSeizureArmyTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 战斗中，友军群体造成负面状态时，有38%概率（受智力影响）使负面状态持续施加增加1回合（混乱及伪报不生效）
	pairGenerals := util.GetPairGeneralArr(s.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_DebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}

			triggerRate := 0.38 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			if util.GenerateRate(triggerRate) {
				if params.DebuffEffect == consts.DebuffEffectType_Chaos ||
					params.DebuffEffect == consts.DebuffEffectType_FalseReport {
					return triggerResp
				}

				params.EffectHolderParams.EffectRound++
			}

			return triggerResp
		})
	}
	// 自身普通攻击后，对负面状态最多的敌军单体造成谋略攻击（伤害率114%，受智力影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, s.tacticsParams)
		enemyGeneral := enemyGenerals[0]
		debuffMaxCnt := util.DeBuffEffectHolderCount(enemyGeneral)
		for _, general := range enemyGenerals {
			if util.DeBuffEffectHolderCount(general) > debuffMaxCnt {
				enemyGeneral = general
				debuffMaxCnt = util.DeBuffEffectHolderCount(general)
			}
		}
		if debuffMaxCnt > 0 {
			dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.14
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     s.tacticsParams,
				AttackGeneral:     currentGeneral,
				SufferGeneral:     enemyGeneral,
				DamageType:        consts.DamageType_Strategy,
				DamageImproveRate: dmgRate,
				TacticId:          s.Id(),
				TacticName:        s.Name(),
			})
		} else {
			// 若敌军都没有负面状态则为损失兵力最多的我军单体恢复兵力（治疗率92%，受智力影响）
			general := util.GetPairLowestSoldierNumGeneral(s.tacticsParams, currentGeneral)
			resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.92)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  s.tacticsParams,
				ProduceGeneral: currentGeneral,
				SufferGeneral:  general,
				ResumeNum:      resumeNum,
				TacticId:       s.Id(),
			})
		}

		return triggerResp
	})

}

func (s SuperviseLeadAndSeizureArmyTactic) Id() consts.TacticId {
	return consts.SuperviseLeadAndSeizureArmy
}

func (s SuperviseLeadAndSeizureArmyTactic) Name() string {
	return "监统震军"
}

func (s SuperviseLeadAndSeizureArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SuperviseLeadAndSeizureArmyTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SuperviseLeadAndSeizureArmyTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SuperviseLeadAndSeizureArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SuperviseLeadAndSeizureArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SuperviseLeadAndSeizureArmyTactic) Execute() {
}

func (s SuperviseLeadAndSeizureArmyTactic) IsTriggerPrepare() bool {
	return false
}
