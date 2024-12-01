package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 一身是胆
// 战斗中，使自己获得洞察状态，武力、智力、速度、统率提高40点，自身为主将时，提升值为50点
type BebraveAllThroughTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BebraveAllThroughTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BebraveAllThroughTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//洞察效果
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
		EffectRate:     1.0,
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})

	val := int64(40)
	//自身为主将时，提升值为50点
	if currentGeneral.IsMaster {
		val = 50
	}

	//武力
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectValue:    val,
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})
	//智力
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectValue:    val,
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})
	//速度
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
		EffectValue:    val,
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})
	//统率
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    val,
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})
}

func (b BebraveAllThroughTactic) Id() consts.TacticId {
	return consts.BebraveAllThrough
}

func (b BebraveAllThroughTactic) Name() string {
	return "一身是胆"
}

func (b BebraveAllThroughTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BebraveAllThroughTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BebraveAllThroughTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BebraveAllThroughTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BebraveAllThroughTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BebraveAllThroughTactic) Execute() {

}

func (b BebraveAllThroughTactic) IsTriggerPrepare() bool {
	return false
}

func (a BebraveAllThroughTactic) SetTriggerPrepare(triggerPrepare bool) {
}
