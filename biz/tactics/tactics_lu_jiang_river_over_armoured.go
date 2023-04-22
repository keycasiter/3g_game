package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//庐江上甲
//提高自身44点统率，并为友军单体分担40%伤害，持续2回合
type LuJiangRiverOverArmouredTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LuJiangRiverOverArmouredTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 0.35
	return l
}

func (l LuJiangRiverOverArmouredTactic) Prepare() {
}

func (l LuJiangRiverOverArmouredTactic) Id() consts.TacticId {
	return consts.LuJiangRiverOverArmoured
}

func (l LuJiangRiverOverArmouredTactic) Name() string {
	return "庐江上甲"
}

func (l LuJiangRiverOverArmouredTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LuJiangRiverOverArmouredTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LuJiangRiverOverArmouredTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LuJiangRiverOverArmouredTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (l LuJiangRiverOverArmouredTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LuJiangRiverOverArmouredTactic) Execute() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral
	currentRound := l.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)

	//施加效果
	currentGeneral.BaseInfo.AbilityAttr.CommandBase += 44
	hlog.CtxInfof(ctx, "[%s]的统率提升了44点",
		currentGeneral.BaseInfo.Name,
	)
	//找到友军单体
	pairGeneral := util.GetPairOneGeneralNotSelf(l.tacticsParams, currentGeneral)
	pairGeneral.ShareResponsibilityForByGeneral = currentGeneral
	//施加效果
	if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_ShareResponsibilityFor, 0.4) {
		//注册消失效果
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral
			revokeRound := params.CurrentRound
			//设置回合
			if currentRound+2 == revokeRound {
				util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_ShareResponsibilityFor)
			}

			return revokeResp
		})
	}
}

func (l LuJiangRiverOverArmouredTactic) IsTriggerPrepare() bool {
	return false
}
