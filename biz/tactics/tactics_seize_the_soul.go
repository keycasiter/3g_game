package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 夺魂挟魄
// 发动概率55%
// 偷取敌军单体38点武力、智力、速度、统率（受智力影响），
// 持续2回合，可叠加2次
type SeizeTheSoulTactic struct {
	tacticsParams *model.TacticsParams
}

func (s SeizeTheSoulTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SeizeTheSoulTactic) TriggerRate() float64 {
	return 0.55
}

func (s SeizeTheSoulTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	return s
}

func (s SeizeTheSoulTactic) Prepare() {
	return
}

func (s SeizeTheSoulTactic) Id() consts.TacticId {
	return consts.SeizeTheSoul
}

func (s SeizeTheSoulTactic) Name() string {
	return "夺魂挟魄"
}

func (s SeizeTheSoulTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SeizeTheSoulTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SeizeTheSoulTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	enemyGeneral := util.GetEnemyOneGeneral(s.tacticsParams)

	//最多叠加两次
	if !util.TacticsBuffEffectCountWrapIncr(currentGeneral, consts.BuffEffectType_SeizeTheSoul, 1, 2) {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			currentGeneral.BaseInfo.Name,
			consts.BuffEffectType_SeizeTheSoul,
		)
		return
	}
	if !util.TacticsDebuffEffectCountWrapIncr(ctx, enemyGeneral, consts.DebuffEffectType_SeizeTheSoul, 1, 2, false) {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			enemyGeneral.BaseInfo.Name,
			consts.DebuffEffectType_SeizeTheSoul,
		)
		return
	}

	//偷取敌军单体38点武力、智力、速度、统率（受智力影响）
	v := 38 + (currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100)
	//提高我军武将
	currentGeneral.BaseInfo.AbilityAttr.ForceBase += v
	currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase += v
	currentGeneral.BaseInfo.AbilityAttr.SpeedBase += v
	currentGeneral.BaseInfo.AbilityAttr.CommandBase += v
	hlog.CtxInfof(ctx, "[%s]的武力提高了%.2f",
		currentGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的智力提高了%.2f",
		currentGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的速度提高了%.2f",
		currentGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的统率提高了%.2f",
		currentGeneral.BaseInfo.Name,
		v)

	//注册撤销效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeGeneral := params.CurrentGeneral
		//次数为0
		if 0 == util.TacticsBuffCountGet(revokeGeneral, consts.BuffEffectType_SeizeTheSoul) {
			revokeGeneral.BaseInfo.AbilityAttr.ForceBase -= v
			revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= v
			revokeGeneral.BaseInfo.AbilityAttr.SpeedBase -= v
			revokeGeneral.BaseInfo.AbilityAttr.CommandBase -= v
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.BuffEffectType_IncrForce)
			hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.BuffEffectType_IncrIntelligence)
			hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.BuffEffectType_IncrSpeed)
			hlog.CtxInfof(ctx, "[%s]的速度降低了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.BuffEffectType_IncrCommand)
			hlog.CtxInfof(ctx, "[%s]的统率降低了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
		}
		//消耗次数-1
		util.TacticsBuffEffectCountWrapDecr(currentGeneral, consts.BuffEffectType_SeizeTheSoul, 1)
		return revokeResp
	})

	//降低敌军武将
	enemyGeneral.BaseInfo.AbilityAttr.ForceBase -= v
	enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= v
	enemyGeneral.BaseInfo.AbilityAttr.SpeedBase -= v
	enemyGeneral.BaseInfo.AbilityAttr.CommandBase -= v
	hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
		enemyGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
		enemyGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的速度降低了%.2f",
		enemyGeneral.BaseInfo.Name,
		v)
	hlog.CtxInfof(ctx, "[%s]的统率降低了%.2f",
		enemyGeneral.BaseInfo.Name,
		v)

	//注册撤销效果
	util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeGeneral := params.CurrentGeneral
		//次数为0
		if 0 == util.TacticsBuffCountGet(revokeGeneral, consts.BuffEffectType_SeizeTheSoul) {
			revokeGeneral.BaseInfo.AbilityAttr.ForceBase += v
			revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase += v
			revokeGeneral.BaseInfo.AbilityAttr.SpeedBase += v
			revokeGeneral.BaseInfo.AbilityAttr.CommandBase += v
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.DebuffEffectType_DecrForce)
			hlog.CtxInfof(ctx, "[%s]的武力提升了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.DebuffEffectType_DecrIntelligence)
			hlog.CtxInfof(ctx, "[%s]的智力提升了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.DebuffEffectType_DecrSpeed)
			hlog.CtxInfof(ctx, "[%s]的速度提升了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				revokeGeneral.BaseInfo.Name,
				consts.DebuffEffectType_DecrCommand)
			hlog.CtxInfof(ctx, "[%s]的统率提升了%.2f",
				revokeGeneral.BaseInfo.Name,
				v)
		}
		//消耗次数-1
		util.TacticsDebuffEffectCountWrapDecr(currentGeneral, consts.DebuffEffectType_SeizeTheSoul, 1)
		return revokeResp
	})
}
