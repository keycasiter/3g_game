package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 长驱直入
// 战斗中，每次造成兵刃伤害后，使自己造成的兵刃伤害提升15%，最大叠加5次，
// 叠加5次后，使我军全体受到伤害降低16%（受武力影响），持续2回合
// 被动 100%
type MarchIntoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MarchIntoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 1.0
	return m
}

func (m MarchIntoTactic) Prepare() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	// 战斗中，每次造成兵刃伤害后，使自己造成的兵刃伤害提升15%，最大叠加5次，
	// 叠加5次后，使我军全体受到伤害降低16%（受武力影响），持续2回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_WeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.15,
			EffectTimes:    1,
			MaxEffectTimes: 5,
			FromTactic:     m.Id(),
			ProduceGeneral: triggerGeneral,
		})
		if effectParams, ok := util.BuffEffectOfTacticGet(triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, m.Id()); ok {
			effectTimes := int64(0)
			for _, effectParam := range effectParams {
				effectTimes += effectParam.EffectTimes
			}
			if effectTimes == 5 {
				pairGenerals := util.GetPairGeneralArr(currentGeneral, m.tacticsParams)
				for _, pairGeneral := range pairGenerals {
					effectRate := 0.16 + triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
					//受到兵刃伤害
					if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
						EffectRate:     effectRate,
						EffectRound:    2,
						FromTactic:     m.Id(),
						ProduceGeneral: triggerGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revoekResp := &vo.TacticsTriggerResult{}
							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    pairGeneral,
								EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
								TacticId:   m.Id(),
							})
							return revoekResp
						})
					}
					//受到谋略伤害
					if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
						EffectRate:     effectRate,
						EffectRound:    2,
						FromTactic:     m.Id(),
						ProduceGeneral: triggerGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revoekResp := &vo.TacticsTriggerResult{}
							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    pairGeneral,
								EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
								TacticId:   m.Id(),
							})
							return revoekResp
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (m MarchIntoTactic) Id() consts.TacticId {
	return consts.MarchInto
}

func (m MarchIntoTactic) Name() string {
	return "长驱直入"
}

func (m MarchIntoTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (m MarchIntoTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MarchIntoTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MarchIntoTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (m MarchIntoTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MarchIntoTactic) Execute() {

}

func (m MarchIntoTactic) IsTriggerPrepare() bool {
	return false
}

func (a MarchIntoTactic) SetTriggerPrepare(triggerPrepare bool) {
}
