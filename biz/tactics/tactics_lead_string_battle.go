package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 引弦力战
// 普通攻击之后，有45%概率获得群攻（普通攻击时对目标同部队其他武将造成伤害）状态（伤害率52%），
// 若已处于群攻状态，则提高16武力，持续3回合，最多可叠加6次
type LeadStringBattle struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LeadStringBattle) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LeadStringBattle) Prepare() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)
	//普通攻击之后，有45%概率获得群攻（普通攻击时对目标同部队其他武将造成伤害）状态（伤害率52%），
	//若已处于群攻状态，则提高16武力，持续3回合，最多可叠加6次
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.45) {
			if util.BuffEffectContains(triggerGeneral, consts.BuffEffectType_GroupAttack) {
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
					EffectValue:    16,
					EffectRound:    3,
					EffectTimes:    1,
					MaxEffectTimes: 6,
					FromTactic:     l.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//注册消失效果
				}
			} else {
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
					EffectRate:     0.52,
					FromTactic:     l.Id(),
					ProduceGeneral: currentGeneral,
				})
			}
		}

		return triggerResp
	})
}

func (l LeadStringBattle) Id() consts.TacticId {
	return consts.LeadStringBattle
}

func (l LeadStringBattle) Name() string {
	return "引弦力战"
}

func (l LeadStringBattle) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (l LeadStringBattle) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LeadStringBattle) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LeadStringBattle) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (l LeadStringBattle) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (l LeadStringBattle) Execute() {

}

func (l LeadStringBattle) IsTriggerPrepare() bool {
	return false
}
