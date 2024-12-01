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

// 挫志怒袭
// 准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续1回合；
// 如果目标已处于虚弱状态则改为造成一次猛击（伤害率188%）
type FrustrationAndAngerAttackTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (f FrustrationAndAngerAttackTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FrustrationAndAngerAttackTactic) Prepare() {
}

func (f FrustrationAndAngerAttackTactic) Id() consts.TacticId {
	return consts.FrustrationAndAngerAttack
}

func (f FrustrationAndAngerAttackTactic) Name() string {
	return "挫志怒袭"
}

func (f FrustrationAndAngerAttackTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FrustrationAndAngerAttackTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FrustrationAndAngerAttackTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FrustrationAndAngerAttackTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FrustrationAndAngerAttackTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FrustrationAndAngerAttackTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral
	currentRound := f.tacticsParams.CurrentRound

	// 准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续1回合；
	// 如果目标已处于虚弱状态则改为造成一次猛击（伤害率188%）

	f.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			f.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if f.isTriggered {
				return triggerResp
			} else {
				f.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				triggerGeneral.BaseInfo.Name,
				f.Name(),
			)

			//找到群体2人
			pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, f.tacticsParams)
			for _, general := range pairGenerals {
				//如果目标已处于虚弱状态则改为造成一次猛击（伤害率188%）
				if util.DeBuffEffectContains(general, consts.DebuffEffectType_PoorHealth) {
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     f.tacticsParams,
						AttackGeneral:     currentGeneral,
						SufferGeneral:     general,
						DamageType:        consts.DamageType_Weapon,
						DamageImproveRate: 1.88,
						TacticId:          f.Id(),
						TacticName:        f.Name(),
					})
				} else {
					//施加效果
					if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
						EffectRound:    1,
						FromTactic:     f.Id(),
						ProduceGeneral: general,
					}).IsSuccess {
						//消失效果
						util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_PoorHealth,
								TacticId:   f.Id(),
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

func (f FrustrationAndAngerAttackTactic) IsTriggerPrepare() bool {
	return f.isTriggerPrepare
}

func (a FrustrationAndAngerAttackTactic) SetTriggerPrepare(triggerPrepare bool) {
}
