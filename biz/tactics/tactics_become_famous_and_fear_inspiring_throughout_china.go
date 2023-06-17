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

// 威震华夏
// 准备1回合，对敌军全体进行猛攻（伤害率146%），使其有50%概率进入缴械、计穷状态，独立判定，持续1回合，并使自己造成的兵刃伤害提升36%，持续2回合；
// 自身为主将时，造成控制效果的概率提高65%
type BecomeFamousAndFearInspiringThroughoutChinaTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.35
	return b
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Prepare() {
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Id() consts.TacticId {
	return consts.BecomeFamousAndFearInspiringThroughoutChina
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Name() string {
	return "威震华夏"
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	// 准备1回合，对敌军全体进行猛攻（伤害率146%），使其有50%概率进入缴械、计穷状态，独立判定，持续1回合，并使自己造成的兵刃伤害提升36%，持续2回合；
	// 自身为主将时，造成控制效果的概率提高65%
	b.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			b.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if b.isTriggered {
				return triggerResp
			} else {
				b.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)

			//对敌军全体进行猛攻（伤害率146%）
			//找到敌军全体
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, b.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.46)
			for _, enemyGeneral := range enemyGenerals {
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: b.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticName:    b.Name(),
				})

				//使其有50%概率进入缴械、计穷状态，独立判定，持续1回合
				//自身为主将时，造成控制效果的概率提高65%
				debuffEffects := []consts.DebuffEffectType{
					consts.DebuffEffectType_CancelWeapon,
					consts.DebuffEffectType_NoStrategy,
				}
				rate := 0.5
				if currentGeneral.IsMaster {
					rate = 0.65
				}
				for _, debuffEffect := range debuffEffects {
					if util.GenerateRate(rate) {
						util.DebuffEffectWrapSet(ctx, enemyGeneral, debuffEffect, &vo.EffectHolderParams{
							EffectRound: 1,
							FromTactic:  b.Id(),
						})
					}
				}
			}
			//并使自己造成的兵刃伤害提升36%，持续2回合
			util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.36,
				EffectRound:    2,
				FromTactic:     b.Id(),
				ProduceGeneral: currentGeneral,
			})
		}

		return triggerResp
	})
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
