package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 倾国倾城
// 准备1回合，使敌军群体（1～2人）进入混乱状态，并有50%概率使我军群体（2人）受到伤害降低16%，持续2回合，自身为女性时，必定选择敌军群体2人
// 40%
type BeautyWhichOverthrowsStatesAndCitiesTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	//是否已经触发
	isTriggered bool
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Prepare() {
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Id() consts.TacticId {
	return consts.BeautyWhichOverthrowsStatesAndCities
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Name() string {
	return "倾国倾城"
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//准备1回合
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
			//使敌军群体（1～2人）进入混乱状态，并有50%概率使我军群体（2人）受到伤害降低16%，持续2回合，自身为女性时，必定选择敌军群体2人

			//自身为女性时，必定选择敌军群体2人
			enemyGenerals := make([]*vo.BattleGeneral, 0)
			if currentGeneral.BaseInfo.Gender == consts.Gender_Female {
				enemyGenerals = util.GetEnemyTwoGeneralByGeneral(triggerGeneral, b.tacticsParams)
			} else {
				enemyGenerals = util.GetEnemyGeneralsOneOrTwoArr(b.tacticsParams)
			}
			//施加混乱效果
			for _, general := range enemyGenerals {
				if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  b.Id(),
				}).IsSuccess {
					//注册消失效果
				}
			}
			//并有50%概率使我军群体（2人）受到伤害降低16%
			if util.GenerateRate(0.5) {
				pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, b.tacticsParams)
				for _, general := range pairGenerals {
					if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
						EffectRate:     0.16,
						EffectRound:    2,
						FromTactic:     b.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						//注册消失效果
						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    general,
							EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
							TacticId:   b.Id(),
						})
					}
					if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
						EffectRate:     0.16,
						EffectRound:    2,
						FromTactic:     b.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						//注册消失效果
						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    general,
							EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
							TacticId:   b.Id(),
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
