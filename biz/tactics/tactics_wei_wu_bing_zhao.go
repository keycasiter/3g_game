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

// 威武并昭
// 战斗中，自身每回合获得33%看破（受速度影响，持续1回合，看破：造成伤害时无视目标一定比例的受到伤害降低效果）；
// 敌军主将每回合行动前，自身有33%概率（受双方速度差影响）对敌军速度低于自身的武将发动1次普通攻击；
// 自身为主将时，没成判定成功后自身速度提升9点（受双方速度差影响），可叠加
type WeiWuBingZhaoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a WeiWuBingZhaoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.0
	return a
}

func (a WeiWuBingZhaoTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，自身每回合获得33%看破（受速度影响，持续1回合，看破：造成伤害时无视目标一定比例的受到伤害降低效果）；
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		kanPoEffectRate := 0.33 + triggerGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100

		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_KanPo, &vo.EffectHolderParams{
			EffectRate:     kanPoEffectRate,
			EffectRound:    1,
			FromTactic:     a.Id(),
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    triggerGeneral,
				EffectType: consts.BuffEffectType_KanPo,
				TacticId:   a.Id(),
			})
		}

		return triggerResp
	})

	// 敌军主将每回合行动前，自身有33%概率（受双方速度差影响）对敌军速度低于自身的武将发动1次普通攻击；
	enemyMasterGeneral := util.GetEnemyMasterGeneral(currentGeneral, a.tacticsParams)
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, a.tacticsParams)
	util.TacticsTriggerWrapRegister(enemyMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		for _, enemyGeneral := range enemyGenerals {
			//速度低于自身的武将
			if enemyGeneral.BaseInfo.AbilityAttr.SpeedBase < currentGeneral.BaseInfo.AbilityAttr.SpeedBase {

				//自身有33%概率（受双方速度差影响）对敌军速度低于自身的武将发动1次普通攻击；
				triggerRate := 0.33
				speedDiff := (currentGeneral.BaseInfo.AbilityAttr.SpeedBase - enemyMasterGeneral.BaseInfo.AbilityAttr.SpeedBase) / 100 / 100
				triggerRate += speedDiff
				if triggerRate < 0 {
					triggerRate = 0
				}
				if util.GenerateRate(triggerRate) {
					damage.AttackDamage(a.tacticsParams, currentGeneral, enemyGeneral, 0)
				} else {
					// 自身为主将时，没成判定成功后自身速度提升9点（受双方速度差影响），可叠加
					if currentGeneral.IsMaster {
						util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
							EffectValue:    cast.ToInt64(9 * (1 + speedDiff)),
							FromTactic:     a.Id(),
							ProduceGeneral: currentGeneral,
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (a WeiWuBingZhaoTactic) Id() consts.TacticId {
	return consts.WeiWuBingZhao
}

func (a WeiWuBingZhaoTactic) Name() string {
	return "威武并昭"
}

func (a WeiWuBingZhaoTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a WeiWuBingZhaoTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a WeiWuBingZhaoTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a WeiWuBingZhaoTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a WeiWuBingZhaoTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Spearman,
		consts.ArmType_Archers,
		consts.ArmType_Mauler,
		consts.ArmType_Apparatus,
	}
}

func (a WeiWuBingZhaoTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

}

func (a WeiWuBingZhaoTactic) IsTriggerPrepare() bool {
	return false
}
