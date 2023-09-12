package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//梦中弑臣
//战斗前2回合，如果自己为主将，则使随机副将为自己分担20%伤害。
//战斗第3回合起，自己行动时如果有负面状态，则获得25%概率反击状态（伤害率150%）
//直到战斗结束
type KillingMinisterInDreamTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (k KillingMinisterInDreamTactic) IsTriggerPrepare() bool {
	return false
}

func (k KillingMinisterInDreamTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	k.tacticsParams = tacticsParams
	k.triggerRate = 1.0
	return k
}

func (k KillingMinisterInDreamTactic) Prepare() {
	ctx := k.tacticsParams.Ctx
	currentGeneral := k.tacticsParams.CurrentGeneral

	//战斗前2回合，如果自己为主将，则使随机副将为自己分担20%伤害。
	if currentGeneral.IsMaster {
		hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
			currentGeneral.BaseInfo.Name,
			k.Name(),
		)
		//找到随机副将
		viceGeneral := util.GetPairViceGeneral(k.tacticsParams)

		//施加分担效果
		if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ShareResponsibilityFor, &vo.EffectHolderParams{
			EffectRate:                      0.2,
			FromTactic:                      k.Id(),
			ProduceGeneral:                  currentGeneral,
			ShareResponsibilityForByGeneral: viceGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral
				triggerRound := params.CurrentRound

				if triggerRound == consts.Battle_Round_Third {
					util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_ShareResponsibilityFor, k.Id())
				}

				return triggerResp
			})
		}
	}
	//战斗第3回合起，自己行动时如果有负面状态，则获得25%概率反击状态（伤害率150%），直到战斗结束
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_KillingMinisterInDream_Prepare, &vo.EffectHolderParams{
		EffectRate: 1.0,
		FromTactic: k.Id(),
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound == consts.Battle_Round_Third {
			//是否有负面效果
			if util.DeBuffEffectContainsCheck(triggerGeneral) {
				//25%
				if util.GenerateRate(0.25) {
					util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_StrikeBack, &vo.EffectHolderParams{
						EffectRate: 1.5,
						FromTactic: k.Id(),
					})
				}
			}
		}
		return triggerResp
	})
}

func (k KillingMinisterInDreamTactic) Id() consts.TacticId {
	return consts.KillingMinisterInDream
}

func (k KillingMinisterInDreamTactic) Name() string {
	return "梦中弑臣"
}

func (k KillingMinisterInDreamTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (k KillingMinisterInDreamTactic) GetTriggerRate() float64 {
	return k.triggerRate
}

func (k KillingMinisterInDreamTactic) SetTriggerRate(rate float64) {
	k.triggerRate = rate
}

func (k KillingMinisterInDreamTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (k KillingMinisterInDreamTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (k KillingMinisterInDreamTactic) Execute() {

}
