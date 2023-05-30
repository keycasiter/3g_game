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

// 血刃争锋
// 战斗中，提高75%普通攻击伤害，普通攻击之后对目标造成酣斗效果，每回合最多触发1次，
// 若目标身上存在3次酣斗效果，消耗全部酣斗效果额外提高自身（11%x酣斗次数）普通攻击伤害，额外提升效果不可叠加，持续到战斗结束
type BloodBladeBattle struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BloodBladeBattle) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BloodBladeBattle) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	//提高75%普通攻击伤害，
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GeneralAttackDamageImprove, &vo.EffectHolderParams{
		EffectRate: 0.75,
		FromTactic: b.Id(),
	})
	//普通攻击之后对目标造成酣斗效果
	guardRound := consts.Battle_Round_Unknow
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//每回合最多触发1次
		if guardRound != params.CurrentRound {
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_FightHard, &vo.EffectHolderParams{
				TriggerRate:    1.0,
				EffectTimes:    1,
				MaxEffectTimes: cast.ToInt64(consts.INT_MAX),
				FromTactic:     b.Id(),
			})
			//更新回合统计
			guardRound = params.CurrentRound
		}

		//若目标身上存在3次酣斗效果，消耗全部酣斗效果额外提高自身（11%x酣斗次数）普通攻击伤害，额外提升效果不可叠加，持续到战斗结束

		return triggerResp
	})
}

func (b BloodBladeBattle) Id() consts.TacticId {
	return consts.BloodBladeBattle
}

func (b BloodBladeBattle) Name() string {
	return "血刃争锋"
}

func (b BloodBladeBattle) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BloodBladeBattle) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BloodBladeBattle) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BloodBladeBattle) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BloodBladeBattle) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BloodBladeBattle) Execute() {

}

func (b BloodBladeBattle) IsTriggerPrepare() bool {
	return false
}
