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

// 金丹秘术
// 战斗前2回合，使我军全体获得35%规避（可以免疫伤害）效果，并在战斗第3回合开始，获得休整状态（每回合恢复一次兵力，回复率58%，受智力影响），持续3回合
type GoldenPillSecretTechniqueTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GoldenPillSecretTechniqueTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GoldenPillSecretTechniqueTactic) Prepare() {
	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)

	//战斗前2回合，使我军全体获得35%规避（可以免疫伤害）效果，并在战斗第3回合开始，获得休整状态（每回合恢复一次兵力，回复率58%，受智力影响），持续3回合

	//找到我军全体
	pairGenerals := util.GetPairGeneralArr(g.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//规避效果施加
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Evade, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     g.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Evade,
					TacticId:   g.Id(),
				})

				return revokeResp
			})
		}
		//并在战斗第3回合开始，获得休整状态（每回合恢复一次兵力，回复率58%，受智力影响），持续3回合
		//施加效果
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
			EffectRound:    3,
			FromTactic:     g.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Rest,
					TacticId:   g.Id(),
				}) {
					resumeNum := cast.ToInt64(revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.58)
					util.ResumeSoldierNum(&util.ResumeParams{
						Ctx:            ctx,
						TacticsParams:  g.tacticsParams,
						ProduceGeneral: revokeGeneral,
						SufferGeneral:  revokeGeneral,
						ResumeNum:      resumeNum,
						TacticId:       g.Id(),
					})
				}

				return revokeResp
			})
		}
	}
}

func (g GoldenPillSecretTechniqueTactic) Id() consts.TacticId {
	return consts.GoldenPillSecretTechnique
}

func (g GoldenPillSecretTechniqueTactic) Name() string {
	return "金丹秘术"
}

func (g GoldenPillSecretTechniqueTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (g GoldenPillSecretTechniqueTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GoldenPillSecretTechniqueTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GoldenPillSecretTechniqueTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (g GoldenPillSecretTechniqueTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GoldenPillSecretTechniqueTactic) Execute() {

}

func (g GoldenPillSecretTechniqueTactic) IsTriggerPrepare() bool {
	return false
}
