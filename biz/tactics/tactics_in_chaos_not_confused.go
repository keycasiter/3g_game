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
)

// 处兹不惑
// 使敌军群体（2人）有60%几率随机获得灼烧、中毒、溃逃状态，每回合持续造成伤害（伤害70%），每种状态独立判定，持续2回合
// 主动，35%
type InChaosNotConfusedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InChaosNotConfusedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.35
	return i
}

func (i InChaosNotConfusedTactic) Prepare() {

}

func (i InChaosNotConfusedTactic) Id() consts.TacticId {
	return consts.InChaosNotConfused
}

func (i InChaosNotConfusedTactic) Name() string {
	return "处兹不惑"
}

func (i InChaosNotConfusedTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i InChaosNotConfusedTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i InChaosNotConfusedTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i InChaosNotConfusedTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InChaosNotConfusedTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InChaosNotConfusedTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)

	// 使敌军群体（2人）有60%几率随机获得灼烧、中毒、溃逃状态，每回合持续造成伤害（伤害70%），每种状态独立判定，持续2回合
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, i.tacticsParams)
	debuffs := []consts.DebuffEffectType{
		consts.DebuffEffectType_Firing,
		consts.DebuffEffectType_Methysis,
		consts.DebuffEffectType_Escape,
	}
	for _, enemyGeneral := range enemyGenerals {
		for _, debuff := range debuffs {
			if util.GenerateRate(0.6) {
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, debuff, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     i.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//消失效果
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: debuff,
							TacticId:   i.Id(),
						}) {
							//伤害
							switch debuff {
							case consts.DebuffEffectType_Firing:
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:     i.tacticsParams,
									AttackGeneral:     currentGeneral,
									SufferGeneral:     enemyGeneral,
									DamageType:        consts.DamageType_Strategy,
									DamageImproveRate: 0.7,
									TacticId:          i.Id(),
									TacticName:        i.Name(),
									EffectName:        fmt.Sprintf("%v", debuff),
								})
							case consts.DebuffEffectType_Methysis:
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:     i.tacticsParams,
									AttackGeneral:     currentGeneral,
									SufferGeneral:     enemyGeneral,
									DamageType:        consts.DamageType_Strategy,
									DamageImproveRate: 0.7,
									TacticId:          i.Id(),
									TacticName:        i.Name(),
									EffectName:        fmt.Sprintf("%v", debuff),
								})
							case consts.DebuffEffectType_Defect:
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:     i.tacticsParams,
									AttackGeneral:     currentGeneral,
									SufferGeneral:     enemyGeneral,
									DamageType:        consts.DamageType_Weapon,
									DamageImproveRate: 0.7,
									TacticId:          i.Id(),
									TacticName:        i.Name(),
									EffectName:        fmt.Sprintf("%v", debuff),
									IsIgnoreDefend:    true,
								})
							}
						}

						return revokeResp
					})
				}
			}
		}
	}
}

func (i InChaosNotConfusedTactic) IsTriggerPrepare() bool {
	return false
}

func (a InChaosNotConfusedTactic) SetTriggerPrepare(triggerPrepare bool) {
}
