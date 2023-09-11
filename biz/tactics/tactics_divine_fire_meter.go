package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 神火计
// 战斗中，每次成功发动主动战法时，有80%几率对敌军全体造成谋略攻击（伤害率68%，受智力影响）；
// 自身为主将时，该次攻击有40%概率对目标施加灼烧状态（伤害率34%，受智力影响），持续1回合
// 被动，100%
type DivineFireMeterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivineFireMeterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DivineFireMeterTactic) Prepare() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//每次成功发动主动战法时，有80%几率对敌军全体造成谋略攻击（伤害率68%，受智力影响）
		if util.GenerateRate(0.8) {
			//找到敌军全体
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, d.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.68)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: d.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Strategy,
					Damage:        dmg,
					TacticName:    d.Name(),
					TacticId:      d.Id(),
				})
				//自身为主将时，该次攻击有40%概率对目标施加灼烧状态（伤害率34%，受智力影响），持续1回合
				if triggerGeneral.IsMaster {
					if util.GenerateRate(0.4) {
						//施加效果
						if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
							EffectRate:  0.34,
							EffectRound: 1,
							FromTactic:  d.Id(),
						}).IsSuccess {
							//注册消失效果
							util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}

								util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    enemyGeneral,
									EffectType: consts.DebuffEffectType_Firing,
									TacticId:   d.Id(),
								})

								return revokeResp
							})
						}
					}
				}
			}
		}

		return triggerResp
	})
}

func (d DivineFireMeterTactic) Id() consts.TacticId {
	return consts.DivineFireMeter
}

func (d DivineFireMeterTactic) Name() string {
	return "神火计"
}

func (d DivineFireMeterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DivineFireMeterTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DivineFireMeterTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DivineFireMeterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (d DivineFireMeterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DivineFireMeterTactic) Execute() {

}

func (d DivineFireMeterTactic) IsTriggerPrepare() bool {
	return false
}
