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
	currentRound := s.tacticsParams.CurrentRound

	//发动概率55%
	if !util.GenerateRate(0.55) {
		return
	} else {
		hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
			currentGeneral.BaseInfo.Name,
			s.Name(),
		)
		//最多叠加两次
		//判断叠加次数
		if mm, ok := currentGeneral.BuffEffectCountMap[consts.BuffEffectType_SeizeTheSoul]; ok {
			if _, okk := mm[2]; okk {
				hlog.CtxInfof(ctx, "[%s][%s]叠加次数已达上限",
					currentGeneral.BaseInfo.Name,
					s.Name(),
				)
				return
			}
		}

		enemyGeneral := util.GetEnemyOneGeneral(s.tacticsParams)
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
		//降低敌军武将
		enemyGeneral.BaseInfo.AbilityAttr.ForceBase -= v
		enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= v
		enemyGeneral.BaseInfo.AbilityAttr.SpeedBase -= v
		enemyGeneral.BaseInfo.AbilityAttr.CommandBase -= v
		hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
			currentGeneral.BaseInfo.Name,
			v)
		hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
			currentGeneral.BaseInfo.Name,
			v)
		hlog.CtxInfof(ctx, "[%s]的速度降低了%.2f",
			currentGeneral.BaseInfo.Name,
			v)
		hlog.CtxInfof(ctx, "[%s]的统率降低了%.2f",
			currentGeneral.BaseInfo.Name,
			v)
		//持续2回合，可叠加2次
		//注册效果
		util.TacticsTriggerWrapSet(currentGeneral, consts.BattleAction_Attack, func(params vo.TacticsTriggerParams) {
			if params.CurrentRound == currentRound+2 {
				currentGeneral.BaseInfo.AbilityAttr.ForceBase -= v
				currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= v
				currentGeneral.BaseInfo.AbilityAttr.SpeedBase -= v
				currentGeneral.BaseInfo.AbilityAttr.CommandBase -= v
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					currentGeneral.BaseInfo.Name,
					consts.BuffEffectType_IncrForce)
				hlog.CtxInfof(ctx, "[%s]的武力降低了%.2f",
					currentGeneral.BaseInfo.Name,
					v)
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					currentGeneral.BaseInfo.Name,
					consts.BuffEffectType_IncrIntelligence)
				hlog.CtxInfof(ctx, "[%s]的智力降低了%.2f",
					currentGeneral.BaseInfo.Name,
					v)
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					currentGeneral.BaseInfo.Name,
					consts.BuffEffectType_IncrSpeed)
				hlog.CtxInfof(ctx, "[%s]的速度降低了%.2f",
					currentGeneral.BaseInfo.Name,
					v)
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					currentGeneral.BaseInfo.Name,
					consts.BuffEffectType_IncrCommand)
				hlog.CtxInfof(ctx, "[%s]的统率降低了%.2f",
					currentGeneral.BaseInfo.Name,
					v)
			}
		})

		//叠加次数
		util.BuffEffectCountWrapAdd(currentGeneral.BuffEffectCountMap,
			consts.BuffEffectType_SeizeTheSoul,
			1,
			1.0,
		)
	}
}
