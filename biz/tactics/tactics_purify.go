package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 净化
// 提升我军群体（2人）24点武力、智力、速度，持续2回合，并移除负面效果
type PurifyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PurifyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.45
	return p
}

func (p PurifyTactic) Prepare() {

}

func (p PurifyTactic) Id() consts.TacticId {
	return consts.Purify
}

func (p PurifyTactic) Name() string {
	return "净化"
}

func (p PurifyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (p PurifyTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PurifyTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PurifyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PurifyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PurifyTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)

	//提升我军群体（2人）24点武力、智力、速度，持续2回合，并移除负面效果
	pairGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, p.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//武力
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectRound:    2,
			EffectValue:    24,
			FromTactic:     p.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrForce,
					TacticId:   p.Id(),
				})

				return revokeResp
			})
		}
		//智力
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
			EffectRound:    2,
			EffectValue:    24,
			FromTactic:     p.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrIntelligence,
					TacticId:   p.Id(),
				})

				return revokeResp
			})
		}
		//速度
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
			EffectRound:    2,
			EffectValue:    24,
			FromTactic:     p.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrSpeed,
					TacticId:   p.Id(),
				})

				return revokeResp
			})
		}
	}
	//移除负面效果
	util.DebuffEffectClean(ctx, currentGeneral)
}

func (p PurifyTactic) IsTriggerPrepare() bool {
	return false
}
