package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//瞋目横矛
//使自己武力提升50点，并获得群攻（普通攻击时对目标同部队其他武将造成伤害）效果（伤害率70%），持续2回合
type AngryEyeHorizontalSpearTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AngryEyeHorizontalSpearTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.0
	return a
}

func (a AngryEyeHorizontalSpearTactic) Prepare() {
}

func (a AngryEyeHorizontalSpearTactic) Id() consts.TacticId {
	return consts.AngryEyeHorizontalSpear
}

func (a AngryEyeHorizontalSpearTactic) Name() string {
	return "瞋目横矛"
}

func (a AngryEyeHorizontalSpearTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a AngryEyeHorizontalSpearTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AngryEyeHorizontalSpearTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AngryEyeHorizontalSpearTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AngryEyeHorizontalSpearTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AngryEyeHorizontalSpearTactic) Execute() {
	//使自己武力提升50点，并获得群攻（普通攻击时对目标同部队其他武将造成伤害）效果（伤害率70%），持续2回合
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound
	//武力提升
	currentGeneral.BaseInfo.AbilityAttr.ForceBase += 50
	hlog.CtxInfof(ctx, "[%s]的武力提高了%d",
		currentGeneral.BaseInfo.Name,
		50,
	)
	//群攻效果
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
		EffectRate: 0.7,
		FromTactic: a.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRound := params.CurrentRound

			if currentRound+2 == triggerRound {
				//武力增益消失
				triggerGeneral.BaseInfo.AbilityAttr.ForceBase -= 50
				hlog.CtxInfof(ctx, "[%s]的武力降低了%d",
					currentGeneral.BaseInfo.Name,
					50,
				)
				//移除群攻效果
				util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_GroupAttack, a.Id())
			}

			return triggerResp
		})
	}
}

func (a AngryEyeHorizontalSpearTactic) IsTriggerPrepare() bool {
	return false
}
