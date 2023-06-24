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

// 傲睨王侯
// 战斗中，发现15个敌军破绽，破绽会分布在全体敌军中，敌军目标受到普通攻击时会触发1个破绽，使该目标降低3%武力、智力、统率、速度（受智力影响），可叠加
// 单个目标的破绽全部触发时，使其进入1回合虚弱状态且受到伤害提高15%（受智力影响），持续2回合，场上所有破绽触发后，敌军群体（2人）武力、智力、统率、速度降低20%（受智力影响）
// 指挥，100%
type ProudPrinceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p ProudPrinceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p ProudPrinceTactic) Prepare() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral
	allFlawCnt := 15

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 战斗中，发现15个敌军破绽，破绽会分布在全体敌军中，敌军目标受到普通攻击时会触发1个破绽，使该目标降低3%武力、智力、统率、速度（受智力影响），可叠加
	allEnemeyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, p.tacticsParams)
	//15个破绽不均匀分布
	for i := 0; i < allFlawCnt; i++ {
		hitIdx := util.GenerateHitOneIdx(len(allEnemeyGenerals))
		//施加破绽
		util.DebuffEffectWrapSet(ctx, allEnemeyGenerals[hitIdx], consts.DebuffEffectType_Flaw, &vo.EffectHolderParams{
			EffectTimes:    1,
			FromTactic:     p.Id(),
			ProduceGeneral: currentGeneral,
		})
	}
	for _, enemeyGeneral := range allEnemeyGenerals {
		util.TacticsTriggerWrapRegister(enemeyGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    triggerGeneral,
				EffectType: consts.DebuffEffectType_Flaw,
				TacticId:   p.Id(),
				//单个目标的破绽全部触发时，使其进入1回合虚弱状态且受到伤害提高15%（受智力影响），持续2回合
				CostOverTriggerFunc: func() {
					//虚弱
					if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_PoorHealth,
								TacticId:   p.Id(),
							})

							return revokeResp
						})
					}
					//受到兵刃伤害提升
					if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
						EffectRound:    2,
						EffectRate:     0.15 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
								TacticId:   p.Id(),
							})

							return revokeResp
						})
					}
					//受到谋略伤害提升
					if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
						EffectRound:    2,
						EffectRate:     0.15 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
								TacticId:   p.Id(),
							})

							return revokeResp
						})
					}
				},
			}) {
				//破绽减少
				allFlawCnt--
				//武力
				util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
					EffectValue:    cast.ToInt64(enemeyGeneral.BaseInfo.AbilityAttr.ForceBase * 0.03),
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
				//智力
				util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
					EffectValue:    cast.ToInt64(enemeyGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.03),
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
				//统率
				util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
					EffectValue:    cast.ToInt64(enemeyGeneral.BaseInfo.AbilityAttr.CommandBase * 0.03),
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
				//速度
				util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
					EffectValue:    cast.ToInt64(enemeyGeneral.BaseInfo.AbilityAttr.SpeedBase * 0.03),
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
			}

			//场上所有破绽触发后，敌军群体（2人）武力、智力、统率、速度降低20%（受智力影响）
			if allFlawCnt == 0 {
				for _, general := range allEnemeyGenerals {
					//武力
					util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
						EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.ForceBase * 0.2),
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					})
					//智力
					util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
						EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.IntelligenceBase * 0.2),
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					})
					//统率
					util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
						EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.CommandBase * 0.2),
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					})
					//速度
					util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
						EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.SpeedBase * 0.2),
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					})
				}
			}

			return triggerResp
		})
	}

	// 单个目标的破绽全部触发时，使其进入1回合虚弱状态且受到伤害提高15%（受智力影响），持续2回合，场上所有破绽触发后，敌军群体（2人）武力、智力、统率、速度降低20%（受智力影响）
}

func (p ProudPrinceTactic) Id() consts.TacticId {
	return consts.ProudPrince
}

func (p ProudPrinceTactic) Name() string {
	return "傲睨王侯"
}

func (p ProudPrinceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p ProudPrinceTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p ProudPrinceTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p ProudPrinceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (p ProudPrinceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p ProudPrinceTactic) Execute() {
}

func (p ProudPrinceTactic) IsTriggerPrepare() bool {
	return false
}
