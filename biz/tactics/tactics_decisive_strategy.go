package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 沉断机谋
// 使敌军群体(2人)统率降低、智力降低30%，持续2回合，并造成谋略伤害（伤害156%，受智力影响）
// 主动 40%
type DecisiveStrategyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DecisiveStrategyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.4
	return d
}

func (d DecisiveStrategyTactic) Prepare() {
}

func (d DecisiveStrategyTactic) Id() consts.TacticId {
	return consts.DecisiveStrategy
}

func (d DecisiveStrategyTactic) Name() string {
	return "沉断机谋"
}

func (d DecisiveStrategyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DecisiveStrategyTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DecisiveStrategyTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DecisiveStrategyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DecisiveStrategyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DecisiveStrategyTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	//使敌军群体(2人)统率降低、智力降低30%，持续2回合，并造成谋略伤害（伤害156%，受智力影响）

	//找到敌军群体2人
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, d.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//降低统率
		decrCommandVal := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.CommandBase * 0.3)
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
			EffectValue: decrCommandVal,
			EffectRound: 2,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_DecrCommand,
					TacticId:   d.Id(),
				})

				return revokeResp
			})
		}
		//降低智力
		decrIntelligenceVal := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.3)
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
			EffectValue: decrIntelligenceVal,
			EffectRound: 2,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_DecrIntelligence,
					TacticId:   d.Id(),
				})

				return revokeResp
			})
		}
		//造成谋略伤害（伤害156%，受智力影响）
		dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.56
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     d.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          d.Id(),
			TacticName:        d.Name(),
		})
	}
}

func (d DecisiveStrategyTactic) IsTriggerPrepare() bool {
	return false
}
