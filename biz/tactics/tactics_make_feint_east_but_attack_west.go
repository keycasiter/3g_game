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

// 声东击西
// 准备1回合，对敌军群体（2人）造成谋略攻击（伤害率175%，受智力影响），并降低30点速度，持续2回合
type MakeFeintToTheEastButAttackInTheWestTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.4
	return m
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) Prepare() {
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) Id() consts.TacticId {
	return consts.MakeFeintToTheEastButAttackInTheWest
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) Name() string {
	return "声东击西"
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) Execute() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral
	currentRound := m.tacticsParams.CurrentRound

	m.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			m.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if m.isTriggered {
				return triggerResp
			} else {
				m.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				m.Name(),
			)

			//准备1回合，对敌军群体（2人）造成谋略攻击（伤害率175%，受智力影响），并降低30点速度，持续2回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, m.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//伤害
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.75)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: m.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Strategy,
					Damage:        dmg,
					TacticName:    m.Name(),
				})
				//降低速度
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
					EffectValue:    30,
					EffectRound:    2,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_DecrSpeed,
							TacticId:   m.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (m MakeFeintToTheEastButAttackInTheWestTactic) IsTriggerPrepare() bool {
	return m.isTriggerPrepare
}
