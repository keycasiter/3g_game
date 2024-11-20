package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 因利制权
// 使我军群体2-3人获得25%规避（受速度影响），持续1回合，该战法发动后会进入1回合冷却
type ProfitBasedSystemOfPowerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a ProfitBasedSystemOfPowerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.45
	return a
}

func (a ProfitBasedSystemOfPowerTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound

	//该战法发动后回进入1回合冷却
	//判断是否冷却
	if ok := currentGeneral.TacticFrozenMap[a.Id()]; ok {
		hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效，无法发动",
			currentGeneral.BaseInfo.Name,
			a.Name(),
		)
		return
	}

	//该战法发动后会进入1回合冷却
	currentGeneral.TacticFrozenMap[a.Id()] = true
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//注册冷却效果消失
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeRound := params.CurrentRound

		//1回合冷却，下下回合冷却结束
		if currentRound+2 == revokeRound {
			currentGeneral.TacticFrozenMap[a.Id()] = false

			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果已消失",
				currentGeneral.BaseInfo.Name,
				a.Name(),
			)
		}
		return revokeResp
	})

	// 使我军群体2-3人获得25%规避（受速度影响），持续1回合
	pairGenerals := util.GetPairGeneralsTwoOrThreeMap(currentGeneral, a.tacticsParams)
	rate := 0.25 + currentGeneral.BaseInfo.AbilityAttr.SpeedBase
	for _, pairGeneral := range pairGenerals {
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Evade, &vo.EffectHolderParams{
			EffectRate:     rate,
			EffectRound:    1,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Evade,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (a ProfitBasedSystemOfPowerTactic) Id() consts.TacticId {
	return consts.ProfitBasedSystemOfPower
}

func (a ProfitBasedSystemOfPowerTactic) Name() string {
	return "因利制权"
}

func (a ProfitBasedSystemOfPowerTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a ProfitBasedSystemOfPowerTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a ProfitBasedSystemOfPowerTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a ProfitBasedSystemOfPowerTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a ProfitBasedSystemOfPowerTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a ProfitBasedSystemOfPowerTactic) Execute() {
}

func (a ProfitBasedSystemOfPowerTactic) IsTriggerPrepare() bool {
	return false
}
