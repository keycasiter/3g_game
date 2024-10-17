package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 暗潮涌动
// 准备1回合，对敌军主将造成一次兵刃攻击（伤害率272%）
// 主动，35%
type UndercurrentSurgeTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (u UndercurrentSurgeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	u.tacticsParams = tacticsParams
	u.triggerRate = 0.35
	return u
}

func (u UndercurrentSurgeTactic) Prepare() {

}

func (u UndercurrentSurgeTactic) Id() consts.TacticId {
	return consts.UndercurrentSurge
}

func (u UndercurrentSurgeTactic) Name() string {
	return "暗潮涌动"
}

func (u UndercurrentSurgeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (u UndercurrentSurgeTactic) GetTriggerRate() float64 {
	return u.triggerRate
}

func (u UndercurrentSurgeTactic) SetTriggerRate(rate float64) {
	u.triggerRate = rate
}

func (u UndercurrentSurgeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (u UndercurrentSurgeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (u UndercurrentSurgeTactic) Execute() {
	ctx := u.tacticsParams.Ctx
	currentGeneral := u.tacticsParams.CurrentGeneral
	currentRound := u.tacticsParams.CurrentRound

	u.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		u.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		//准备回合释放
		if currentRound+2 == triggerRound {
			u.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if u.isTriggered {
				return triggerResp
			} else {
				u.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				u.Name(),
			)

			//准备1回合，对敌军主将造成一次兵刃攻击（伤害率272%）
			enemyMasterGeneral := util.GetEnemyMasterGeneral(u.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.72)
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams: u.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: enemyMasterGeneral,
				DamageType:    consts.DamageType_Weapon,
				Damage:        dmg,
				TacticId:      u.Id(),
				TacticName:    u.Name(),
			})
		}

		return triggerResp
	})
}

func (u UndercurrentSurgeTactic) IsTriggerPrepare() bool {
	return false
}
