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

// 密计诛逆
// 战斗中，我军主将造成伤害（大于300）后，有50%概率使敌军单体造成最终伤害降低15%（受统率影响，持续2回合，可叠加3次）
// 且前2回合有50%概率（受统率影响）令主将对敌军全体造成1次兵刃伤害（伤害率25%，受主将统率影响）；
// 战斗第6回合，自身对敌军单体造成斩杀伤害（伤害率100% +25%*最终降伤施加次数）
type MiJiZhuNiTactic struct {
	tacticsParams      *model.TacticsParams
	triggerRate        float64
	deduceTriggerTimes int64
}

func (a MiJiZhuNiTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a MiJiZhuNiTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，我军主将造成伤害（大于300）后，有50%概率使敌军单体造成最终伤害降低15%（受统率影响，持续2回合，可叠加3次）
	pairMasterGeneral := util.GetPairMasterGeneral(currentGeneral, a.tacticsParams)
	util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if params.CurrentDamage > 300 {
			if util.GenerateRate(0.5) {
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, a.tacticsParams)

				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_LaunchFinalDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.15 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100,
					EffectRound:    2,
					FromTactic:     a.Id(),
					EffectTimes:    1,
					MaxEffectTimes: 3,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    enemyGeneral,
						EffectType: consts.DebuffEffectType_LaunchFinalDamageDeduce,
						TacticId:   a.Id(),
					})
				}
				a.deduceTriggerTimes++
			}

			// 且前2回合有50%概率（受统率影响）令主将对敌军全体造成1次兵刃伤害（伤害率25%，受主将统率影响）；
			if triggerRound <= consts.Battle_Round_Second {
				triggerRate := 0.5 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100
				if util.GenerateRate(triggerRate) {
					enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, a.tacticsParams)
					for _, general := range enemyGenerals {
						damage.TacticDamage(&damage.TacticDamageParam{
							TacticsParams:     a.tacticsParams,
							AttackGeneral:     pairMasterGeneral,
							SufferGeneral:     general,
							DamageType:        consts.DamageType_Weapon,
							DamageImproveRate: 0.25 + general.BaseInfo.AbilityAttr.CommandBase/100/100,
							TacticId:          a.Id(),
							TacticName:        a.Name(),
						})
					}
				}
			}
			// 战斗第6回合，自身对敌军单体造成斩杀伤害（伤害率100% +25%*最终降伤施加次数）
			if triggerRound == consts.Battle_Round_Sixth {
				enemeyGenearl := util.GetEnemyOneGeneralByGeneral(currentGeneral, a.tacticsParams)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     a.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     enemeyGenearl,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 1 + 0.25*cast.ToFloat64(a.deduceTriggerTimes),
					TacticId:          a.Id(),
					TacticName:        a.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (a MiJiZhuNiTactic) Id() consts.TacticId {
	return consts.MiJiZhuNi
}

func (a MiJiZhuNiTactic) Name() string {
	return "密计诛逆"
}

func (a MiJiZhuNiTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a MiJiZhuNiTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a MiJiZhuNiTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a MiJiZhuNiTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a MiJiZhuNiTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a MiJiZhuNiTactic) Execute() {
}

func (a MiJiZhuNiTactic) IsTriggerPrepare() bool {
	return false
}

func (a MiJiZhuNiTactic) SetTriggerPrepare(triggerPrepare bool) {
}
