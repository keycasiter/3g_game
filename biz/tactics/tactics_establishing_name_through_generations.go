package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 累世立名
// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率126%），降低其24统率（受武力影响）并附加灼烧状态，
// 每回合持续造成伤害（伤害率60%，受智力影响），持续2回合
// 自身为主将时，额外提高我军全体80统率，持续3回合
type EstablishingNameThroughGenerationsTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (e EstablishingNameThroughGenerationsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 0.5
	return e
}

func (e EstablishingNameThroughGenerationsTactic) Prepare() {

}

func (e EstablishingNameThroughGenerationsTactic) Id() consts.TacticId {
	return consts.EstablishingNameThroughGenerations
}

func (e EstablishingNameThroughGenerationsTactic) Name() string {
	return "累世立名"
}

func (e EstablishingNameThroughGenerationsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (e EstablishingNameThroughGenerationsTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EstablishingNameThroughGenerationsTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EstablishingNameThroughGenerationsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (e EstablishingNameThroughGenerationsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EstablishingNameThroughGenerationsTactic) Execute() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral
	currentRound := e.tacticsParams.CurrentRound

	e.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		e.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		//准备回合释放
		if currentRound+2 == triggerRound {
			e.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if e.isTriggered {
				return triggerResp
			} else {
				e.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				e.Name(),
			)

			// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率126%）
			// 自身为主将时，额外提高我军全体80统率，持续3回合

			//找到敌军2人
			enemyGenerals := util.GetEnemyGeneralsTwoArr(e.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.26)
			for _, enemyGeneral := range enemyGenerals {
				//对敌军单体造成一次兵刃攻击（伤害率126%）
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: e.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      e.Id(),
					TacticName:    e.Name(),
				})
				//降低其24统率（受武力影响）
				//降低统率
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
					EffectValue:    24,
					EffectRound:    2,
					FromTactic:     e.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//注册消失效果
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        nil,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_DecrCommand,
							TacticId:   e.Id(),
						})

						return revokeResp
					})
				}
				//并附加灼烧状态，每回合持续造成伤害（伤害率60%，受智力影响），持续2回合
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     e.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//消失效果
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Firing,
							TacticId:   e.Id(),
						}) {
							//持续伤害
							fireDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.6)
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams: e.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: enemyGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        fireDmg,
								TacticId:      e.Id(),
								TacticName:    e.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
							})
						}

						return revokeResp
					})
				}
			}
			//自身为主将时，额外提高我军全体80统率，持续3回合
			if currentGeneral.IsMaster {
				//找到我军全体
				pairGenerals := util.GetPairGeneralArr(e.tacticsParams)
				for _, pairGeneral := range pairGenerals {
					if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
						EffectValue:    80,
						EffectRound:    3,
						FromTactic:     e.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						//注册消失
						util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.BuffEffectType_IncrCommand,
								TacticId:   e.Id(),
							})

							return revokeResp
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (e EstablishingNameThroughGenerationsTactic) IsTriggerPrepare() bool {
	return e.isTriggerPrepare
}
