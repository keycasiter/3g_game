package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 移花接木
// 使敌我全体受到治疗提升18%（受自身最高属性影响），并将敌军全体受到治疗的26%（受自身最高属性影响）转移到自身，持续1回合
type YiHuaJieMuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a YiHuaJieMuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.5
	return a
}

func (a YiHuaJieMuTactic) Prepare() {

}

func (a YiHuaJieMuTactic) Id() consts.TacticId {
	return consts.YiHuaJieMu
}

func (a YiHuaJieMuTactic) Name() string {
	return "移花接木"
}

func (a YiHuaJieMuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a YiHuaJieMuTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a YiHuaJieMuTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a YiHuaJieMuTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a YiHuaJieMuTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a YiHuaJieMuTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 使敌我全体受到治疗提升18%（受自身最高属性影响）
	pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
	_, val := util.GetGeneralHighestAttr(currentGeneral)

	for _, pairGeneral := range pairGenerals {
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_SufferResumeImprove, &vo.EffectHolderParams{
			EffectRate:     0.18 + val/100/100,
			EffectRound:    1,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    triggerGeneral,
					EffectType: consts.BuffEffectType_SufferResumeImprove,
					TacticId:   a.Id(),
				})

				return triggerResp
			})
		}
	}

	//并将敌军全体受到治疗的26%（受自身最高属性影响）转移到自身，持续1回合
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, a.tacticsParams)
	for _, general := range enemyGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferResume, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferResumeDeduce, &vo.EffectHolderParams{
				EffectRate:     0.26 + val/100/100,
				EffectRound:    1,
				FromTactic:     a.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//转移到自己身上
				util.ResumeSoldierNum(&util.ResumeParams{
					Ctx:            ctx,
					TacticsParams:  a.tacticsParams,
					ProduceGeneral: currentGeneral,
					SufferGeneral:  currentGeneral,
					ResumeNum:      a.tacticsParams.CurrentResumeNum,
					TacticId:       a.Id(),
				})

				//注销消失效果
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_SufferResumeDeduce,
						TacticId:   a.Id(),
					})

					return revokeResp
				})
			}

			return triggerResp
		})
	}
}

func (a YiHuaJieMuTactic) IsTriggerPrepare() bool {
	return false
}
