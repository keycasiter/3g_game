package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 槊血纵横
// 战斗中，使自己获得34点武力及54%群攻效果，自身为主将时，群攻值为60%
// 被动 100%
type BloodyAndUnrestrainedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BloodyAndUnrestrainedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BloodyAndUnrestrainedTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//使自己获得34点武力
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectValue: 34,
		FromTactic:  b.Id(),
	})
	//54%群攻效果,自身为主将时，群攻值为60%
	rate := 0.54
	if currentGeneral.IsMaster {
		rate = 0.6
	}
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
		EffectRate: rate,
		FromTactic: b.Id(),
	})
}

func (b BloodyAndUnrestrainedTactic) Id() consts.TacticId {
	return consts.BloodyAndUnrestrained
}

func (b BloodyAndUnrestrainedTactic) Name() string {
	return "槊血纵横"
}

func (b BloodyAndUnrestrainedTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) Execute() {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
