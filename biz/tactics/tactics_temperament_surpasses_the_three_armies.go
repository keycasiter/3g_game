package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 气凌三军
// 受到普通攻击时对攻击者进行一次反击（伤害率52%），自身为副将时，伤害率提升至74%
type TemperamentSurpassesTheThreeArmiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TemperamentSurpassesTheThreeArmiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TemperamentSurpassesTheThreeArmiesTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//受到普通攻击时对攻击者进行一次反击（伤害率52%），自身为副将时，伤害率提升至74%
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		attackGeneral := params.AttackGeneral

		dmgRate := 0.52
		if !triggerGeneral.IsMaster {
			dmgRate = 0.74
		}

		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: t.tacticsParams,
			AttackGeneral: triggerGeneral,
			SufferGeneral: attackGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      t.Id(),
			TacticName:    t.Name(),
			EffectName:    fmt.Sprintf("%v", consts.BuffEffectType_StrikeBack),
		})

		return triggerResp
	})
}

func (t TemperamentSurpassesTheThreeArmiesTactic) Id() consts.TacticId {
	return consts.TemperamentSurpassesTheThreeArmies
}

func (t TemperamentSurpassesTheThreeArmiesTactic) Name() string {
	return "气凌三军"
}

func (t TemperamentSurpassesTheThreeArmiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TemperamentSurpassesTheThreeArmiesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TemperamentSurpassesTheThreeArmiesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TemperamentSurpassesTheThreeArmiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TemperamentSurpassesTheThreeArmiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TemperamentSurpassesTheThreeArmiesTactic) Execute() {

}

func (t TemperamentSurpassesTheThreeArmiesTactic) IsTriggerPrepare() bool {
	return false
}
