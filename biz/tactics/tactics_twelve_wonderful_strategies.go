package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 十二奇策
// 移除敌军群体（1～2人）增益状态，提高我军全体1回合6%主动战法发动率（受智力影响）并使其下次发动主动战法后，对敌军单体造成谋略攻击（伤害率102%，受智力影响）
// 主动，45%
type TwelveWonderfulStrategiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TwelveWonderfulStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TwelveWonderfulStrategiesTactic) Prepare() {
}

func (t TwelveWonderfulStrategiesTactic) Id() consts.TacticId {
	return consts.TwelveWonderfulStrategies
}

func (t TwelveWonderfulStrategiesTactic) Name() string {
	return "十二奇策"
}

func (t TwelveWonderfulStrategiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TwelveWonderfulStrategiesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TwelveWonderfulStrategiesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TwelveWonderfulStrategiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TwelveWonderfulStrategiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TwelveWonderfulStrategiesTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//移除敌军群体（1～2人）增益状态
	enemyGenerals := util.GetEnemyGeneralsOneOrTwoArr(t.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//移除增益状态
		util.BuffEffectClean(ctx, enemyGeneral)
	}
	//提高我军全体1回合6%主动战法发动率（受智力影响）
	pairGenerals := util.GetPairGeneralArr(t.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		triggerRate := 0.06 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
			TriggerRate:    triggerRate,
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
	//并使其下次发动主动战法后，对敌军单体造成谋略攻击（伤害率102%，受智力影响）
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TwelveWonderfulStrategies_Prepare, &vo.EffectHolderParams{
		EffectTimes:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TwelveWonderfulStrategies_Prepare,
				TacticId:   t.Id(),
			}) {
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(revokeGeneral, t.tacticsParams)
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.02)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Strategy,
					Damage:        dmg,
					TacticId:      t.Id(),
					TacticName:    t.Name(),
				})
			}

			return revokeResp
		})
	}
}

func (t TwelveWonderfulStrategiesTactic) IsTriggerPrepare() bool {
	return false
}
