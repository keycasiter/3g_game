package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/spf13/cast"
)

// 鬼神霆威
// 普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率204%）
// 自身为主将且当目标兵力低于50%时，额外提高伤害（受目标损失兵力影响，最多提高50%）
type GhostGodThunderForceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GhostGodThunderForceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.35
	return g
}

func (g GhostGodThunderForceTactic) Prepare() {

}

func (g GhostGodThunderForceTactic) Id() consts.TacticId {
	return consts.GhostGodThunderForce
}

func (g GhostGodThunderForceTactic) Name() string {
	return "鬼神霆威"
}

func (g GhostGodThunderForceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (g GhostGodThunderForceTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GhostGodThunderForceTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GhostGodThunderForceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (g GhostGodThunderForceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GhostGodThunderForceTactic) Execute() {
	//普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率204%）
	//自身为主将且当目标兵力低于50%时，额外提高伤害（受目标损失兵力影响，最多提高50%）

	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)

	dmgRate := 2.04
	if currentGeneral.IsMaster {
		//剩余兵力
		remainSoldierNum := g.tacticsParams.CurrentSufferGeneral.SoldierNum - g.tacticsParams.CurrentSufferGeneral.LossSoldierNum
		//兵力低于50%时，额外提高伤害（受目标损失兵力影响，最多提高50%）
		if cast.ToFloat64(remainSoldierNum/g.tacticsParams.CurrentSufferGeneral.SoldierNum) < 0.5 {
			extraRate := cast.ToFloat64(g.tacticsParams.CurrentSufferGeneral.LossSoldierNum / 100 / 100)
			if extraRate > 0.5 {
				extraRate = 0.5
			}
			dmgRate += extraRate
		}
	}

	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: g.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: g.tacticsParams.CurrentSufferGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    g.Name(),
		TacticId:      g.Id(),
	})

}

func (g GhostGodThunderForceTactic) IsTriggerPrepare() bool {
	return false
}
