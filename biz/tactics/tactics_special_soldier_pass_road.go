package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 奇兵间道
// 战斗中，自身发动准备战法时，有75%几率（受武力影响）减少1回合准备时间，
// 战斗前4回合，主动战法造成的伤害提高30%（受武力影响），
// 第5回合时，若自身兵力低于50%，获得45%倒戈，直到战斗结束
// 被动，100%
type SpecialSoldierPassRoadTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SpecialSoldierPassRoadTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SpecialSoldierPassRoadTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 战斗中，自身发动准备战法时，有75%几率（受武力影响）减少1回合准备时间，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerTactic := params.CurrentTactic.(_interface.Tactics)

		if consts.ActivePrepareTacticsMap[triggerTactic.Id()] {

		}

		return triggerResp
	})
	// 战斗前4回合，主动战法造成的伤害提高30%（受武力影响），
	// 第5回合时，若自身兵力低于50%，获得45%倒戈，直到战斗结束
}

func (s SpecialSoldierPassRoadTactic) Id() consts.TacticId {
	return consts.SpecialSoldierPassRoad
}

func (s SpecialSoldierPassRoadTactic) Name() string {
	return "奇兵间道"
}

func (s SpecialSoldierPassRoadTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SpecialSoldierPassRoadTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SpecialSoldierPassRoadTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SpecialSoldierPassRoadTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s SpecialSoldierPassRoadTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SpecialSoldierPassRoadTactic) Execute() {
}

func (s SpecialSoldierPassRoadTactic) IsTriggerPrepare() bool {
	return false
}

func (a SpecialSoldierPassRoadTactic) SetTriggerPrepare(triggerPrepare bool) {
}
