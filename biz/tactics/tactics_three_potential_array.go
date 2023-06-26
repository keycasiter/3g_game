package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 三势阵
// 我军三名武将阵营均不相同，且我方主将自带战法为主动战法或突击战法时，战斗前5回合，主将提高16%自带主动、
// 突击战法发动几率，每回合行动前，使损失兵力较多的副将受到伤害降低30%，另一面副将造成伤害提高25%，持续1回合
type ThreePotentialArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThreePotentialArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ThreePotentialArrayTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//我军三名武将阵营均不相同
	pairGenerals := util.GetPairGeneralArr(t.tacticsParams)
	groupMap := make(map[consts.Group]bool, 0)
	for _, pairGeneral := range pairGenerals {
		groupMap[pairGeneral.BaseInfo.Group] = true
	}
	if len(groupMap) != 3 {
		hlog.CtxInfof(ctx, "武将阵营不符合，无法发动")
		return
	}
	//且我方主将自带战法为主动战法或突击战法时
	pairMasterGeneral := util.GetPairMasterGeneral(t.tacticsParams)
	if !consts.ActiveTacticsMap[pairMasterGeneral.EquipTactics[0].Id] &&
		!consts.AssaultTacticsMap[pairMasterGeneral.EquipTactics[0].Id] {
		hlog.CtxInfof(ctx, "主将自带战法不为主动战法或突击战法，无法发动")
		return
	}
	//战斗前5回合，主将提高16%自带主动、 突击战法发动几率
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove, &vo.EffectHolderParams{
		EffectRound:    5,
		TriggerRate:    0.16,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	//突击战法发动几率
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_TacticsAssaultTriggerImprove, &vo.EffectHolderParams{
		EffectRound:    5,
		TriggerRate:    0.16,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TacticsAssaultTriggerImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	//每回合行动前，使损失兵力较多的副将受到伤害降低30%，另一面副将造成伤害提高25%，持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//找到损失兵力较多的副将
		viceGenerals := util.GetPairViceGenerals(t.tacticsParams)
		lowesGeneral := util.GetLowestSoliderNumGeneral(viceGenerals)

		//兵刃伤害降低
		if util.BuffEffectWrapSet(ctx, lowesGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.3,
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    lowesGeneral,
				EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
				TacticId:   t.Id(),
			})
		}
		//谋略伤害降低
		if util.BuffEffectWrapSet(ctx, lowesGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.3,
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    lowesGeneral,
				EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
				TacticId:   t.Id(),
			})
		}

		return triggerResp
	})
}

func (t ThreePotentialArrayTactic) Id() consts.TacticId {
	return consts.ThreePotentialArray
}

func (t ThreePotentialArrayTactic) Name() string {
	return "三势阵"
}

func (t ThreePotentialArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ThreePotentialArrayTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThreePotentialArrayTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThreePotentialArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (t ThreePotentialArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThreePotentialArrayTactic) Execute() {
}

func (t ThreePotentialArrayTactic) IsTriggerPrepare() bool {
	return false
}
