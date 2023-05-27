package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 杯蛇鬼车
// 准备1回合，对敌军群体(2人)发动一次谋略攻击（伤害率153%，受智力影响），
// 并为我军群体(2人)恢复一定兵力（恢复率102%，受智力影响）
type CupSnakeGhostCarTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (c CupSnakeGhostCarTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.5
	return c
}

func (c CupSnakeGhostCarTactic) Prepare() {
}

func (c CupSnakeGhostCarTactic) Id() consts.TacticId {
	return consts.CupSnakeGhostCar
}

func (c CupSnakeGhostCarTactic) Name() string {
	return "杯蛇鬼车"
}

func (c CupSnakeGhostCarTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CupSnakeGhostCarTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CupSnakeGhostCarTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CupSnakeGhostCarTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c CupSnakeGhostCarTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CupSnakeGhostCarTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral
	currentRound := c.tacticsParams.CurrentRound

	//准备1回合，对敌军群体(2人)发动一次谋略攻击（伤害率153%，受智力影响）
	c.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if currentRound+1 == triggerRound {
			//准备回合释放
			c.isTriggerPrepare = false
			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				c.Name(),
			)

			//找到敌军2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, c.tacticsParams)
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.53)
			for _, enemyGeneral := range enemyGenerals {
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: c.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					Damage:        dmg,
					TacticName:    c.Name(),
				})
			}
		}
		//并为我军群体(2人)恢复一定兵力（恢复率102%，受智力影响）
		//找到我军2人
		pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, c.tacticsParams)
		for _, pairGeneral := range pairGenerals {
			resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.02)
			resume, origin, final := util.ResumeSoldierNum(ctx, pairGeneral, resumeNum)
			hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
				pairGeneral.BaseInfo.Name,
				resume,
				origin,
				final,
			)
		}

		return triggerResp
	})
}

func (c CupSnakeGhostCarTactic) IsTriggerPrepare() bool {
	return c.isTriggerPrepare
}
