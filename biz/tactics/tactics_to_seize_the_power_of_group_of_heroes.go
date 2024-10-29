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

// 搦战群雄
// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率200%），随后使自己造成兵刃伤害提高25%，受到兵刃伤害降低25%，（受武力影响），持续2回合
// 主动，35%
type ToSeizeThePowerOfGroupOfHeroesTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Prepare() {
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Id() consts.TacticId {
	return consts.ToSeizeThePowerOfGroupOfHeroes
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Name() string {
	return "搦战群雄"
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	currentRound := t.tacticsParams.CurrentRound

	t.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			t.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if t.isTriggered {
				return triggerResp
			} else {
				t.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				t.Name(),
			)
			//准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率200%）
			enemeyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, t.tacticsParams)
			for _, enemeyGeneral := range enemeyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     t.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemeyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 2.0,
					TacticId:          t.Id(),
					TacticName:        t.Name(),
				})
			}
			//随后使自己造成兵刃伤害提高25%，受到兵刃伤害降低25%，（受武力影响），持续2回合
			effectRate := 0.25 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
			if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRound:    2,
				EffectRate:     effectRate,
				FromTactic:     t.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
						TacticId:   t.Id(),
					})

					return revokeResp
				})
			}

			if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRound:    2,
				EffectRate:     effectRate,
				FromTactic:     t.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
						TacticId:   t.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) IsTriggerPrepare() bool {
	return false
}
