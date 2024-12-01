package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 独行赴斗
// 嘲讽（强迫目标普通攻击自己）敌军全体，同时提高40%统率，持续2回合
type TravelingAloneToFightTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TravelingAloneToFightTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.5
	return t
}

func (t TravelingAloneToFightTactic) Prepare() {
}

func (t TravelingAloneToFightTactic) Id() consts.TacticId {
	return consts.TravelingAloneToFight
}

func (t TravelingAloneToFightTactic) Name() string {
	return "独行赴斗"
}

func (t TravelingAloneToFightTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TravelingAloneToFightTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TravelingAloneToFightTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TravelingAloneToFightTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TravelingAloneToFightTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TravelingAloneToFightTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//嘲讽（强迫目标普通攻击自己）敌军全体，同时提高40%统率，持续2回合
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, t.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//嘲讽
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     t.Id(),
			TauntByTarget:  currentGeneral,
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Taunt,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
	//提高统率
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectRound:    2,
		EffectValue:    40,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrCommand,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TravelingAloneToFightTactic) IsTriggerPrepare() bool {
	return false
}

func (a TravelingAloneToFightTactic) SetTriggerPrepare(triggerPrepare bool) {
}
