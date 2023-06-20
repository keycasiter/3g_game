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

// 智计
// 使敌军群体（2人）的武力、智力降低38（受智力影响），持续2回合，最多叠加2次
type IntelligentStrategyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i IntelligentStrategyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.55
	return i
}

func (i IntelligentStrategyTactic) Prepare() {
}

func (i IntelligentStrategyTactic) Id() consts.TacticId {
	return consts.IntelligentStrategy
}

func (i IntelligentStrategyTactic) Name() string {
	return "智计"
}

func (i IntelligentStrategyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (i IntelligentStrategyTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IntelligentStrategyTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IntelligentStrategyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i IntelligentStrategyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IntelligentStrategyTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)
	//使敌军群体（2人）的武力、智力降低38（受智力影响），持续2回合，最多叠加2次
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, i.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		effectVal := 38 + cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100)
		//武力
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
			EffectValue:    effectVal,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromTactic:     i.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_DecrForce,
					TacticId:   i.Id(),
				})

				return revokeResp
			})
		}
		//智力
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
			EffectValue:    effectVal,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromTactic:     i.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_DecrIntelligence,
					TacticId:   i.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (i IntelligentStrategyTactic) IsTriggerPrepare() bool {
	return false
}
