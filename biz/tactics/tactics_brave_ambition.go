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

// 义胆雄心
// 战斗中，奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
// 偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
// 自身为主将时，降低属性效果受自身对应属性影响
type BraveAmbitionTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BraveAmbitionTactic) IsTriggerPrepare() bool {
	return false
}

func (b BraveAmbitionTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BraveAmbitionTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BraveAmbitionTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BraveAmbitionTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BraveAmbitionTactic) Prepare() {
	// 战斗中，奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
	// 偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
	// 自身为主将时，降低属性效果受自身对应属性影响
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//注册效果
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_BraveAmbition_Prepare, &vo.EffectHolderParams{
		FromTactic: b.Id(),
	}).IsSuccess {
		//注册触发效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerGeneral := params.CurrentGeneral
			currentRound := params.CurrentRound
			triggerResp := &vo.TacticsTriggerResult{}

			//奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
			if currentRound%2 != 0 {
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					triggerGeneral.BaseInfo.Name,
					b.Name(),
					consts.BuffEffectType_BraveAmbition_Prepare,
				)
				//找到敌军单体
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, b.tacticsParams)
				//造成伤害
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.84)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: b.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Weapon,
					TacticName:    b.Name(),
				})

				//施加效果
				debuffSetResp := util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  b.Id(),
				})
				//降低武力64点
				decrNum := float64(64)
				if debuffSetResp.IsSuccess && !debuffSetResp.IsRefreshEffect {
					//自身为主将时，降低属性效果受自身对应属性影响
					if triggerGeneral.IsMaster {
						decrNum += triggerGeneral.BaseInfo.AbilityAttr.ForceBase / 100.00
					}
					enemyGeneral.BaseInfo.AbilityAttr.ForceBase -= decrNum
					hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
						enemyGeneral.BaseInfo.Name,
						decrNum)
				}
				//效果注册
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					triggerCostGeneral := params.CurrentGeneral
					triggerCostResp := &vo.TacticsTriggerResult{}
					//效果消耗
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerCostGeneral,
						EffectType: consts.DebuffEffectType_DecrForce,
						TacticId:   b.Id(),
					})
					//效果恢复
					if util.DeBuffEffectOfTacticIsDeplete(enemyGeneral, consts.DebuffEffectType_DecrForce, b.Id()) {
						util.DebuffEffectWrapRemove(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, b.Id())

						enemyGeneral.BaseInfo.AbilityAttr.ForceBase += decrNum
						hlog.CtxInfof(ctx, "[%s]的武力提高了%.2f",
							enemyGeneral.BaseInfo.Name,
							decrNum)
					}
					return triggerCostResp
				})
			}

			//偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
			if currentRound%2 == 0 {
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					triggerGeneral.BaseInfo.Name,
					b.Name(),
					consts.BuffEffectType_BraveAmbition_Prepare,
				)
				//找到敌军2人
				enemyGenerals := util.GetEnemyGeneralsTwoArr(b.tacticsParams)
				//造成伤害
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.76)
				for _, enemyGeneral := range enemyGenerals {
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: b.tacticsParams,
						AttackGeneral: triggerGeneral,
						SufferGeneral: enemyGeneral,
						Damage:        dmg,
						DamageType:    consts.DamageType_Strategy,
						TacticName:    b.Name(),
					})
					//降低智力34点，持续2回合；
					decrNum := float64(34)
					//施加效果
					debuffSetResp := util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
						EffectRound: 2,
						FromTactic:  b.Id(),
					})
					if debuffSetResp.IsSuccess && !debuffSetResp.IsRefreshEffect {
						//自身为主将时，降低属性效果受自身对应属性影响
						if triggerGeneral.IsMaster {
							decrNum += triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100.00
						}
						enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= decrNum
						hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
							enemyGeneral.BaseInfo.Name,
							decrNum)
					}
					//效果注册
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						triggerCostGeneral := params.CurrentGeneral
						triggerCostResp := &vo.TacticsTriggerResult{}
						//效果消耗
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    triggerCostGeneral,
							EffectType: consts.DebuffEffectType_DecrIntelligence,
							TacticId:   b.Id(),
						})

						//效果恢复
						if util.DeBuffEffectOfTacticIsDeplete(enemyGeneral, consts.DebuffEffectType_DecrIntelligence, b.Id()) {
							util.DebuffEffectWrapRemove(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, b.Id())

							enemyGeneral.BaseInfo.AbilityAttr.ForceBase += decrNum
							hlog.CtxInfof(ctx, "[%s]的智力提高了%.2f",
								enemyGeneral.BaseInfo.Name,
								decrNum)
						}
						return triggerCostResp
					})
				}
			}

			return triggerResp
		})
	}
}

func (b BraveAmbitionTactic) Id() consts.TacticId {
	return consts.BraveAmbition
}

func (b BraveAmbitionTactic) Name() string {
	return "义胆雄心"
}

func (b BraveAmbitionTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BraveAmbitionTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveAmbitionTactic) Execute() {
	return
}
