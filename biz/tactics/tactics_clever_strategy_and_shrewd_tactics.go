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

// 神机妙算
// 敌军群体(2人)发动主动战法时，有35%几率令其失败并对其造成谋略伤害(伤害率100%，受智力影响)，
// 自身为主将时，该次伤害会基于双方智力之差额外提高
type CleverStrategyAndShrewdTacticsTactic struct {
	tacticsParams *model.TacticsParams
}

func (c CleverStrategyAndShrewdTacticsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	return c
}

func (c CleverStrategyAndShrewdTacticsTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	//敌军群体(2人)发动主动战法时，有35%几率令其失败并对其造成谋略伤害(伤害率100%，受智力影响)，
	//找到两个敌军
	enemyGenerals := util.GetEnemyGeneralsTwoArr(c.tacticsParams)
	//注册触发效果
	for _, sufferGeneral := range enemyGenerals {
		util.TacticsTriggerWrapSet(sufferGeneral, consts.BattleAction_ExecuteActiveTactic, func(params *vo.TacticsTriggerParams) {
			//35%几率
			if !util.GenerateRate(0.35) {
				hlog.CtxInfof(ctx, "[%s]执行来自[%s]【%s】的「神机妙算」效果因几率没有生效",
					sufferGeneral.BaseInfo.Name,
					currentGeneral.BaseInfo.Name,
					c.Name(),
				)
				return
			} else {
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「神机妙算」效果",
					sufferGeneral.BaseInfo.Name,
					c.Name(),
				)
				dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.00)
				//自身为主将时，该次伤害会基于双方智力之差额外提高
				if currentGeneral.IsMaster {
					diff := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase -
						sufferGeneral.BaseInfo.AbilityAttr.IntelligenceBase
					if diff > 0 {
						dmgNum += cast.ToInt64(diff)
					}
				}
				finalDmg, originNum, remaindNum := util.TacticDamage(c.tacticsParams, currentGeneral, sufferGeneral, dmgNum)
				hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的「神机妙算」效果，损失了兵力%d(%d↘%d️️️)",
					sufferGeneral.BaseInfo.Name,
					currentGeneral.BaseInfo.Name,
					c.Name(),
					finalDmg,
					originNum,
					remaindNum,
				)
			}
		})
	}
	//自身为主将时，该次伤害会基于双方智力之差额外提高
}

func (c CleverStrategyAndShrewdTacticsTactic) Id() consts.TacticId {
	return consts.CleverStrategyAndShrewdTactics
}

func (c CleverStrategyAndShrewdTacticsTactic) Name() string {
	return "神机妙算"
}

func (c CleverStrategyAndShrewdTacticsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c CleverStrategyAndShrewdTacticsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CleverStrategyAndShrewdTacticsTactic) Execute() {
	return
}
