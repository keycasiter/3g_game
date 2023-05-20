package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//横戈跃马
//战斗前3回合，使双方全体造成的谋略伤害降低30%，第三回合起，使我军全体造成的兵刃伤害提升20%（受速度影响），持续到战斗结束
type GallantTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GallantTactic) IsTriggerPrepare() bool {
	return false
}

func (g GallantTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GallantTactic) Prepare() {
	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)
	//战斗前3回合，使双方全体造成的谋略伤害降低30%
	allGenerals := util.GetAllGenerals(g.tacticsParams)
	for _, general := range allGenerals {
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate: 0.3,
			FromTactic: g.Id(),
		}).IsSuccess {
			hlog.CtxInfof(ctx, "[%s]造成谋略伤害降低30%%",
				general.BaseInfo.Name,
			)
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerRound := params.CurrentRound
				triggerGeneral := params.CurrentGeneral

				if triggerRound == consts.Battle_Round_Fourth {
					if util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, g.Id()) {
						hlog.CtxInfof(ctx, "[%s]造成「%v」效果已消失",
							general.BaseInfo.Name,
							consts.DebuffEffectType_LaunchStrategyDamageDeduce,
						)
					}
				}
				return triggerResp
			})
		}
	}

	//第三回合起，使我军全体造成的兵刃伤害提升20%（受速度影响），持续到战斗结束
	pairGenerals := util.GetPairGeneralArr(g.tacticsParams)
	for _, general := range pairGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound
			triggerGeneral := params.CurrentGeneral

			if triggerRound == consts.Battle_Round_Third {
				//受速度影响
				rate := 0.2 + general.BaseInfo.AbilityAttr.SpeedBase/100/100
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate: rate,
					FromTactic: g.Id(),
				}).IsSuccess {
					hlog.CtxInfof(ctx, "[%s]造成兵刃伤害提高20%%",
						general.BaseInfo.Name,
					)
				}
			}

			return triggerResp
		})
	}
}

func (g GallantTactic) Id() consts.TacticId {
	return consts.Gallant
}

func (g GallantTactic) Name() string {
	return "横戈跃马"
}

func (g GallantTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (g GallantTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GallantTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GallantTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (g GallantTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GallantTactic) Execute() {
	return
}
