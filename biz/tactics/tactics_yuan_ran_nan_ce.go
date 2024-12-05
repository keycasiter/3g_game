package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 渊然难测
// 战斗中，我军群体(2-3)人受到普通攻击时，有50%概率（受统率影响，自身为主将时，基础概率提升至60%）使伤害来源受到伤害提升6%（受统率影响），可叠加3次
// 首回合触发时，若伤害来源武将武力高于智力，则使其进入缴械状态，否则使其进入计穷状态，持续2回合
type YuanRanNanCeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a YuanRanNanCeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a YuanRanNanCeTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，我军群体(2-3)人受到普通攻击时，有50%概率（受统率影响，自身为主将时，基础概率提升至60%）使伤害来源受到伤害提升6%（受统率影响），可叠加3次
	// 首回合触发时，若伤害来源武将武力高于智力，则使其进入缴械状态，否则使其进入计穷状态，持续2回合

	pairGenerals := util.GetPairGeneralsTwoOrThreeMap(currentGeneral, a.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferGeneralAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.SufferAttackGeneral
			triggerRound := params.CurrentRound

			triggerRate := 0.5
			if triggerGeneral.IsMaster {
				triggerRate = 0.6
			}
			triggerRate += currentGeneral.BaseInfo.AbilityAttr.CommandBase / 100 / 100

			if util.GenerateRate(triggerRate) {
				effectRate := 0.06 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100

				//受到谋略伤害提升
				if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     effectRate,
					EffectRound:    2,
					EffectTimes:    1,
					MaxEffectTimes: 3,
					FromTactic:     a.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    params.AttackGeneral,
						EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
						TacticId:   a.Id(),
					})
				}

				//受到兵刃伤害提升
				if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     effectRate,
					EffectRound:    2,
					EffectTimes:    1,
					MaxEffectTimes: 3,
					FromTactic:     a.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    params.AttackGeneral,
						EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
						TacticId:   a.Id(),
					})
				}
			}

			//首回合触发时，若伤害来源武将武力高于智力，则使其进入缴械状态，否则使其进入计穷状态，持续2回合
			if triggerRound == consts.Battle_Round_First {
				attr, _ := util.GetGeneralHighestBetweenForceOrIntelligence(params.AttackGeneral)
				if attr == consts.AbilityAttr_Force {
					//缴械
					if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     a.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    params.AttackGeneral,
							EffectType: consts.DebuffEffectType_CancelWeapon,
							TacticId:   a.Id(),
						})
					}
				} else {
					//计穷
					if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     a.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    params.AttackGeneral,
							EffectType: consts.DebuffEffectType_NoStrategy,
							TacticId:   a.Id(),
						})
					}
				}
			}

			return triggerResp
		})
	}
}

func (a YuanRanNanCeTactic) Id() consts.TacticId {
	return consts.YuanRanNanCe
}

func (a YuanRanNanCeTactic) Name() string {
	return "渊然难测"
}

func (a YuanRanNanCeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a YuanRanNanCeTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a YuanRanNanCeTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a YuanRanNanCeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a YuanRanNanCeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a YuanRanNanCeTactic) Execute() {
}

func (a YuanRanNanCeTactic) IsTriggerPrepare() bool {
	return false
}

func (a YuanRanNanCeTactic) SetTriggerPrepare(triggerPrepare bool) {
}
