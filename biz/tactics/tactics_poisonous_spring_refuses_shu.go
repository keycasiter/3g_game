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

// 毒泉拒蜀
// 对敌军单体（80%概率选择敌军统率最高的武将）施加猛毒，2回合后消失
// 拥有猛毒的敌军每次受到普通攻击会叠加一层猛毒，最多叠加3层，层数叠满时会立即消失并清除自身所有战法冷却时间，
// 猛毒消失时会对敌军全体造成谋略伤害（伤害率150%，每有一层猛毒伤害率提高40%，受智力影响），
// 我方蛮族造成的非普攻伤害也会增加猛毒层数，该战法发动后会进入1回合冷却
// 主动 60%
type PoisonousSpringRefusesShuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PoisonousSpringRefusesShuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p PoisonousSpringRefusesShuTactic) Prepare() {

}

func (p PoisonousSpringRefusesShuTactic) Id() consts.TacticId {
	return consts.PoisonousSpringRefusesShu
}

func (p PoisonousSpringRefusesShuTactic) Name() string {
	return "毒泉拒蜀"
}

func (p PoisonousSpringRefusesShuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PoisonousSpringRefusesShuTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PoisonousSpringRefusesShuTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PoisonousSpringRefusesShuTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PoisonousSpringRefusesShuTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PoisonousSpringRefusesShuTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral
	currentRound := p.tacticsParams.CurrentRound

	//判断是否冷却
	if ok := currentGeneral.TacticFrozenMap[p.Id()]; ok {
		hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效，无法发动",
			currentGeneral.BaseInfo.Name,
			p.Name(),
		)
		return
	}

	currentGeneral.TacticFrozenMap[p.Id()] = true
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)

	//注册冷却效果消失
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeRound := params.CurrentRound

		//1回合冷却，下下回合冷却结束
		if currentRound+2 == revokeRound {
			currentGeneral.TacticFrozenMap[p.Id()] = false

			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果已消失",
				currentGeneral.BaseInfo.Name,
				p.Name(),
			)
		}
		return revokeResp
	})

	// 对敌军单体（80%概率选择敌军统率最高的武将）施加猛毒，2回合后消失
	var enemyGeneral *vo.BattleGeneral
	if util.GenerateRate(0.8) {
		enemyGeneral = util.GetEnemyGeneralWhoIsHighestForce(p.tacticsParams)
	} else {
		enemyGeneral = util.GetEnemyOneGeneralByGeneral(currentGeneral, p.tacticsParams)
	}
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_StrongMethysis, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_StrongMethysis,
				TacticId:   p.Id(),
				CostOverTriggerFunc: func() {
					// 猛毒消失时会对敌军全体造成谋略伤害（伤害率150%，每有一层猛毒伤害率提高40%，受智力影响）
					allEnemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, p.tacticsParams)
					for _, general := range allEnemyGenerals {
						dmgRate := 1.5
						effectTimes := int64(0)
						if effectParams, ok := util.DeBuffEffectOfTacticGet(general, consts.DebuffEffectType_StrongMethysis, p.Id()); ok {
							for _, param := range effectParams {
								effectTimes += param.EffectTimes
							}
						}
						dmgRate += 0.4 * cast.ToFloat64(effectTimes)

						dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * dmgRate)
						damage.TacticDamage(&damage.TacticDamageParam{
							TacticsParams: p.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: general,
							DamageType:    consts.DamageType_Strategy,
							Damage:        dmg,
							TacticId:      p.Id(),
							TacticName:    p.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_StrongMethysis),
						})
					}
				},
			})

			return revokeResp
		})
	}
	// 拥有猛毒的敌军每次受到普通攻击会叠加一层猛毒，最多叠加3层，层数叠满时会立即消失并清除自身所有战法冷却时间
	allEnemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, p.tacticsParams)
	for _, general := range allEnemyGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.DeBuffEffectContains(triggerGeneral, consts.DebuffEffectType_StrongMethysis) {
				effectTimes := int64(0)
				if effectParams, ok := util.DeBuffEffectOfTacticGet(triggerGeneral, consts.DebuffEffectType_StrongMethysis, p.Id()); ok {
					for _, param := range effectParams {
						effectTimes += param.EffectTimes
					}
				}
				if effectTimes == 3 {
					//清除该效果
					util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_StrongMethysis, p.Id())
					//清除战法冷却
					util.TacticFrozenClean(ctx, triggerGeneral)
				}
			} else {
				util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_StrongMethysis, &vo.EffectHolderParams{
					EffectTimes:    1,
					MaxEffectTimes: 3,
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
			}

			return triggerResp
		})
	}

	// 我方蛮族造成的非普攻伤害也会增加猛毒层数，该战法发动后会进入1回合冷却
	barbarianPairGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range allEnemyGenerals {
		if util.IsContainsGeneralTag(general.BaseInfo.GeneralTag, consts.GeneralTag_Barbarian) {
			barbarianPairGenerals = append(barbarianPairGenerals, general)
		}
	}
	for _, general := range barbarianPairGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_TacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			sufferGeneral := p.tacticsParams.CurrentSufferGeneral

			if util.DeBuffEffectContains(sufferGeneral, consts.DebuffEffectType_StrongMethysis) {
				effectTimes := int64(0)
				if effectParams, ok := util.DeBuffEffectOfTacticGet(sufferGeneral, consts.DebuffEffectType_StrongMethysis, p.Id()); ok {
					for _, param := range effectParams {
						effectTimes += param.EffectTimes
					}
				}
				if effectTimes == 3 {
					//清除该效果
					util.DebuffEffectWrapRemove(ctx, sufferGeneral, consts.DebuffEffectType_StrongMethysis, p.Id())
					//清除战法冷却
					util.TacticFrozenClean(ctx, sufferGeneral)
				}
			} else {
				util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_StrongMethysis, &vo.EffectHolderParams{
					EffectTimes:    1,
					MaxEffectTimes: 3,
					FromTactic:     p.Id(),
					ProduceGeneral: triggerGeneral,
				})
			}

			return triggerResp
		})
	}
}

func (p PoisonousSpringRefusesShuTactic) IsTriggerPrepare() bool {
	return false
}
