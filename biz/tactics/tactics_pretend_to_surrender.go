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

// 诈降
// 战斗开始后首回合，自己获得混乱（攻击和战法无差别选择目标）状态，第2回合起获得休整状态（每回合恢复一次兵力，治疗率80%，受智力影响），持续3回合
type PretendToSurrenderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PretendToSurrenderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p PretendToSurrenderTactic) Prepare() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		//战斗开始后首回合，自己获得混乱（攻击和战法无差别选择目标）状态
		if triggerRound == consts.Battle_Round_First {
			//施加效果
			if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
				EffectRound:    3,
				FromTactic:     p.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Chaos,
						TacticId:   p.Id(),
					})

					return revokeResp
				})
			}
		}
		//第2回合起获得休整状态（每回合恢复一次兵力，治疗率80%，受智力影响），持续3回合
		if triggerRound == consts.Battle_Round_Second {
			//施加效果
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
				EffectRound:    3,
				FromTactic:     p.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_Rest,
						TacticId:   p.Id(),
					}) {
						resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.8)
						util.ResumeSoldierNum(&util.ResumeParams{
							Ctx:            ctx,
							TacticsParams:  p.tacticsParams,
							ProduceGeneral: currentGeneral,
							SufferGeneral:  revokeGeneral,
							ResumeNum:      resumeNum,
							TacticId:       p.Id(),
						})
					}

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (p PretendToSurrenderTactic) Id() consts.TacticId {
	return consts.PretendToSurrender
}

func (p PretendToSurrenderTactic) Name() string {
	return "诈降"
}

func (p PretendToSurrenderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (p PretendToSurrenderTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PretendToSurrenderTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PretendToSurrenderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (p PretendToSurrenderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PretendToSurrenderTactic) Execute() {

}

func (p PretendToSurrenderTactic) IsTriggerPrepare() bool {
	return false
}

func (a PretendToSurrenderTactic) SetTriggerPrepare(triggerPrepare bool) {
}
