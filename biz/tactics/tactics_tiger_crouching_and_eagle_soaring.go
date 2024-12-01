package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 虎踞鹰扬
// 战斗中使自身免疫缴械（无法进行普通攻击）状态，普通攻击之后，使自己造成兵刃伤害提高7%，最多叠加4次，
// 叠加4次后，使自身获得群攻状态（伤害率30%），持续1回合
type TigerCrouchingAndEagleSoaringTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerCrouchingAndEagleSoaringTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerCrouchingAndEagleSoaringTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//战斗中使自身免疫缴械（无法进行普通攻击）状态
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ImmunityCancelWeapon, &vo.EffectHolderParams{
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	})
	//普通攻击之后，使自己造成兵刃伤害提高7%，最多叠加4次，叠加4次后，使自身获得群攻状态（伤害率30%），持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.07,
			EffectRound:    1,
			EffectTimes:    1,
			MaxEffectTimes: 4,
			FromTactic:     t.Id(),
			ProduceGeneral: triggerGeneral,
		})

		//判定叠加次数
		effectTimes := int64(0)
		if effectParams, ok := util.BuffEffectOfTacticGet(triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, t.Id()); ok {
			for _, param := range effectParams {
				effectTimes += param.EffectTimes
			}
		}
		if effectTimes == 4 {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
				EffectRate:     0.3,
				EffectRound:    1,
				FromTactic:     t.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_GroupAttack,
						TacticId:   t.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (t TigerCrouchingAndEagleSoaringTactic) Id() consts.TacticId {
	return consts.TigerCrouchingAndEagleSoaring
}

func (t TigerCrouchingAndEagleSoaringTactic) Name() string {
	return "虎踞鹰扬"
}

func (t TigerCrouchingAndEagleSoaringTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TigerCrouchingAndEagleSoaringTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerCrouchingAndEagleSoaringTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerCrouchingAndEagleSoaringTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TigerCrouchingAndEagleSoaringTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TigerCrouchingAndEagleSoaringTactic) Execute() {
}

func (t TigerCrouchingAndEagleSoaringTactic) IsTriggerPrepare() bool {
	return false
}

func (a TigerCrouchingAndEagleSoaringTactic) SetTriggerPrepare(triggerPrepare bool) {
}
