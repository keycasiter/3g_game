package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 万军夺帅
// 使敌军速度最快的武将速度降低85%，并有45几率使其进入遇袭、破坏、禁疗、计穷、缴械状态，每种状态独立判定，持续2回合，该战法发动后会进入1回合冷却
type WanJunDuoShuaiTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a WanJunDuoShuaiTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.6
	return a
}

func (a WanJunDuoShuaiTactic) Prepare() {
}

func (a WanJunDuoShuaiTactic) Id() consts.TacticId {
	return consts.WanJunDuoShuai
}

func (a WanJunDuoShuaiTactic) Name() string {
	return "万军夺帅"
}

func (a WanJunDuoShuaiTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (a WanJunDuoShuaiTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a WanJunDuoShuaiTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a WanJunDuoShuaiTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a WanJunDuoShuaiTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Spearman,
		consts.ArmType_Archers,
		consts.ArmType_Mauler,
		consts.ArmType_Apparatus,
	}
}

func (a WanJunDuoShuaiTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound

	//判断是否冷却
	if ok := currentGeneral.TacticFrozenMap[a.Id()]; ok {
		hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效，无法发动",
			currentGeneral.BaseInfo.Name,
			a.Name(),
		)
		return
	}

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

	// 使敌军速度最快的武将速度降低85%，并有45几率使其进入遇袭、破坏、禁疗、计穷、缴械状态，每种状态独立判定，持续2回合，该战法发动后会进入1回合冷却
	enemeyHighestSpeedGeneral := util.GetEnemyGeneralWhoIsHighestSpeed(currentGeneral, a.tacticsParams)

	//使敌军速度最快的武将速度降低85%
	if util.DebuffEffectWrapSet(ctx, enemeyHighestSpeedGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
		EffectRate:     0.85,
		EffectRound:    2,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
			Ctx:        ctx,
			General:    enemeyHighestSpeedGeneral,
			EffectType: consts.DebuffEffectType_DecrSpeed,
			TacticId:   a.Id(),
		})
	}

	//并有45几率使其进入遇袭、破坏、禁疗、计穷、缴械状态，每种状态独立判定，持续2回合
	for _, effectType := range []consts.DebuffEffectType{
		consts.DebuffEffectType_BeAttacked,
		consts.DebuffEffectType_ProhibitionTreatment,
		consts.DebuffEffectType_Break,
		consts.DebuffEffectType_CancelWeapon,
		consts.DebuffEffectType_NoStrategy,
	} {
		if util.GenerateRate(0.45) {
			if util.DebuffEffectWrapSet(ctx, enemeyHighestSpeedGeneral, effectType, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     a.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemeyHighestSpeedGeneral,
					EffectType: effectType,
					TacticId:   a.Id(),
				})
			}
		}
	}
}

func (a WanJunDuoShuaiTactic) IsTriggerPrepare() bool {
	return false
}

func (a WanJunDuoShuaiTactic) SetTriggerPrepare(triggerPrepare bool) {
}
