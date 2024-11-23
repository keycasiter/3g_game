package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 臻于至善
// 战斗中，自身武力、智力、统率、速度、魅力提升36点，且每回合有36%概率（每种属性对判定，受对应属性影响）使属性提升效果翻倍，持续1回合
type ZhenYuZhiShanTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a ZhenYuZhiShanTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a ZhenYuZhiShanTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//战斗中，自身武力、智力、统率、速度、魅力提升36点，且每回合有36%概率（每种属性对判定，受对应属性影响）使属性提升效果翻倍，持续1回合
	for _, buffEffectType := range []consts.BuffEffectType{
		consts.BuffEffectType_IncrForce,
		consts.BuffEffectType_IncrIntelligence,
		consts.BuffEffectType_IncrCommand,
		consts.BuffEffectType_IncrSpeed,
		consts.BuffEffectType_IncrCharm,
	} {
		util.BuffEffectWrapSet(ctx, currentGeneral, buffEffectType, &vo.EffectHolderParams{
			EffectValue:    36,
			FromTactic:     a.Id(),
			ProduceGeneral: currentGeneral,
		})
	}

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		for idx, buffEffectType := range []consts.BuffEffectType{
			consts.BuffEffectType_IncrForce,
			consts.BuffEffectType_IncrIntelligence,
			consts.BuffEffectType_IncrCommand,
			consts.BuffEffectType_IncrSpeed,
			consts.BuffEffectType_IncrCharm,
		} {
			attArr := []float64{
				triggerGeneral.BaseInfo.AbilityAttr.ForceBase,
				triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase,
				triggerGeneral.BaseInfo.AbilityAttr.CommandBase,
				triggerGeneral.BaseInfo.AbilityAttr.SpeedBase,
				triggerGeneral.BaseInfo.AbilityAttr.CharmBase,
			}

			triggerRate := 0.36 + attArr[idx]/100/100

			if util.GenerateRate(triggerRate) {
				if util.BuffEffectWrapSet(ctx, triggerGeneral, buffEffectType, &vo.EffectHolderParams{
					EffectValue:    36,
					FromTactic:     a.Id(),
					EffectRound:    1,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerGeneral,
						EffectType: buffEffectType,
						TacticId:   a.Id(),
					})
				}
			}
		}

		return triggerResp
	})
}

func (a ZhenYuZhiShanTactic) Id() consts.TacticId {
	return consts.ZhenYuZhiShan
}

func (a ZhenYuZhiShanTactic) Name() string {
	return "臻于至善"
}

func (a ZhenYuZhiShanTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a ZhenYuZhiShanTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a ZhenYuZhiShanTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a ZhenYuZhiShanTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a ZhenYuZhiShanTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a ZhenYuZhiShanTactic) Execute() {
}

func (a ZhenYuZhiShanTactic) IsTriggerPrepare() bool {
	return false
}
