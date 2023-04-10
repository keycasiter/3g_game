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

// 连环计
// 发动概率35%
// 准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，
// 并发动谋略攻击（伤害率156%，受智力影响）
type InterlockedStratagemsTactic struct {
	tacticsParams *model.TacticsParams
}

func (i InterlockedStratagemsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i InterlockedStratagemsTactic) TriggerRate() float64 {
	return 0.35
}

func (i InterlockedStratagemsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	return i
}

func (i InterlockedStratagemsTactic) Prepare() {
	return
}

func (i InterlockedStratagemsTactic) Id() consts.TacticId {
	return consts.InterlockedStratagems
}

func (i InterlockedStratagemsTactic) Name() string {
	return "连环计"
}

func (i InterlockedStratagemsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InterlockedStratagemsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InterlockedStratagemsTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral
	currentRound := i.tacticsParams.CurrentRound

	//准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，
	//获取敌军全体
	allGenerals := util.GetEnemyGeneralMap(i.tacticsParams)
	for _, general := range allGenerals {
		//次数判断
		if !util.TacticsDebuffEffectCountWrapIncr(ctx, general, consts.DebuffEffectType_InterlockedStratagems, 2, 2, false) {
			return
		}
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
			general.BaseInfo.Name,
			consts.DebuffEffectType_InterlockedStratagems,
		)
		//注册效果
		registerFunc := func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			//准备一回合
			if !(currentRound+1 == params.CurrentRound) {
				return triggerResp
			}
			//次数消耗
			if !util.TacticsDebuffEffectCountWrapDecr(general,
				consts.DebuffEffectType_InterlockedStratagems,
				1,
			) {
				return triggerResp
			}
			//释放效果
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				i.Name(),
				consts.DebuffEffectType_InterlockedStratagems,
			)
			//找到2个队友进行伤害反弹
			pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, i.tacticsParams)
			for _, reboundGeneral := range pairGenerals {
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.15)
				finalNum, holdNum, remainNum, isEffect := util.TacticDamage(i.tacticsParams, currentGeneral, reboundGeneral, dmg, consts.BattleAction_Unknow)
				if !isEffect {
					return triggerResp
				}
				hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的「%v」效果，损失了兵力%d(%d↘%d)",
					reboundGeneral.BaseInfo.Name,
					currentGeneral.BaseInfo.Name,
					i.Name(),
					consts.DebuffEffectType_InterlockedStratagems,
					finalNum,
					holdNum,
					remainNum,
				)
			}
			return triggerResp
		}
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDamage, registerFunc)

	}
	//并发动谋略攻击（伤害率156%，受智力影响）
}
