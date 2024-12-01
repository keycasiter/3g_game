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

// 南蛮渠魁
// 战斗中，每回合行动时有49%几率对敌军全体造成兵刃伤害（伤害率106%），若未生效则提高7%发动概率，
// 自身受到7次普通攻击后会进入1回合震慑（无法行动）状态；
// 自身为主将时，每回合有15%（部队每多1名蛮族武将额外提高15%）概率使全体蛮族造成伤害提高15%（受自身损失兵力影响，最多80%），持续1回合
// 指挥 100%
type NanManQuKuiTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NanManQuKuiTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	n.tacticsParams = tacticsParams
	n.triggerRate = 1.0
	return n
}

func (n NanManQuKuiTactic) Prepare() {

	ctx := n.tacticsParams.Ctx
	currentGeneral := n.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		n.Name(),
	)

	// 战斗中，每回合行动时有49%几率对敌军全体造成兵刃伤害（伤害率106%），若未生效则提高7%发动概率，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.49) {
			enemyGenerals := util.GetEnemyGeneralArr(currentGeneral, n.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     n.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 1.06,
					TacticId:          n.Id(),
					TacticName:        n.Name(),
				})
			}
		} else {
			//提高发动概率
			n.SetTriggerRate(n.GetTriggerRate() + 0.07)
		}

		// 自身为主将时，每回合有15%（部队每多1名蛮族武将额外提高15%）概率使全体蛮族造成伤害提高15%（受自身损失兵力影响，最多80%），持续1回合
		if triggerGeneral.IsMaster {
			barbarianCnt := 0
			barbarianGenerals := make([]*vo.BattleGeneral, 0)
			allGenerals := util.GetAllGenerals(n.tacticsParams)
			for _, general := range allGenerals {
				if util.IsContainsGeneralTag(general.BaseInfo.GeneralTag, consts.GeneralTag_Barbarian) {
					barbarianCnt++
					barbarianGenerals = append(barbarianGenerals, general)
				}
			}
			triggerRate := 0.15 * cast.ToFloat64(1+barbarianCnt)
			if util.GenerateRate(triggerRate) {
				for _, barbarianGeneral := range barbarianGenerals {
					effectRate := 0.15 + cast.ToFloat64(triggerGeneral.LossSoldierNum/100/100)
					if effectRate > 0.8 {
						effectRate = 0.8
					}
					if util.BuffEffectWrapSet(ctx, barbarianGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
						EffectRate:     effectRate,
						EffectRound:    1,
						FromTactic:     n.Id(),
						ProduceGeneral: triggerGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(barbarianGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral
							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
								TacticId:   n.Id(),
							})

							return revokeResp
						})
					}
				}
			}
		}

		return triggerResp
	})

	// 自身受到7次普通攻击后会进入1回合震慑（无法行动）状态；
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if triggerGeneral.SufferExecuteGeneralAttackNum == 7 {
			if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
				EffectRound:    1,
				FromTactic:     n.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Awe,
						TacticId:   n.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (n NanManQuKuiTactic) Id() consts.TacticId {
	return consts.NanManQuKui
}

func (n NanManQuKuiTactic) Name() string {
	return "南蛮渠魁"
}

func (n NanManQuKuiTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (n NanManQuKuiTactic) GetTriggerRate() float64 {
	return n.triggerRate
}

func (n NanManQuKuiTactic) SetTriggerRate(rate float64) {
	n.triggerRate = rate
}

func (n NanManQuKuiTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (n NanManQuKuiTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (n NanManQuKuiTactic) Execute() {

}

func (n NanManQuKuiTactic) IsTriggerPrepare() bool {
	return false
}

func (a NanManQuKuiTactic) SetTriggerPrepare(triggerPrepare bool) {
}
