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

//非攻制胜
//指挥 100%
// 战斗第3回合起，我军武力最高单体每次造成兵刃伤害时，会治疗我军兵力最低单体，
// 治疗量相当于此次伤害量的10%（受造成伤害武将的武力影响）
type NonOffensiveVictoryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NonOffensiveVictoryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	n.tacticsParams = tacticsParams
	n.triggerRate = 1.0
	return n
}

func (n NonOffensiveVictoryTactic) Prepare() {
	currentGeneral := n.tacticsParams.CurrentGeneral
	ctx := n.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		n.Name(),
	)
	//找到我军最高武力单体
	pairGeneral := util.GetPairGeneralWhoIsHighestForce(n.tacticsParams)

	// 战斗第3回合起，我军武力最高单体每次造成兵刃伤害时，会治疗我军兵力最低单体，治疗量相当于此次伤害量的10%（受造成伤害武将的武力影响）
	util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral
		damageNum := params.CurrentDamage

		if triggerRound >= consts.Battle_Round_Third {
			pairGenerals := util.GetPairGeneralArr(n.tacticsParams)
			lowSoliderNumGeneral := util.GetLowestSoliderNumGeneral(pairGenerals)
			resumeNum := cast.ToInt64(cast.ToFloat64(damageNum)*0.1) + cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100)

			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  n.tacticsParams,
				ProduceGeneral: triggerGeneral,
				SufferGeneral:  lowSoliderNumGeneral,
				ResumeNum:      resumeNum,
			})
		}

		return triggerResp
	})
}

func (n NonOffensiveVictoryTactic) Id() consts.TacticId {
	return consts.NonOffensiveVictory
}

func (n NonOffensiveVictoryTactic) Name() string {
	return "非攻制胜"
}

func (n NonOffensiveVictoryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (n NonOffensiveVictoryTactic) GetTriggerRate() float64 {
	return n.triggerRate
}

func (n NonOffensiveVictoryTactic) SetTriggerRate(rate float64) {
	n.triggerRate = rate
}

func (n NonOffensiveVictoryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (n NonOffensiveVictoryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (n NonOffensiveVictoryTactic) Execute() {

}

func (n NonOffensiveVictoryTactic) IsTriggerPrepare() bool {
	return false
}
